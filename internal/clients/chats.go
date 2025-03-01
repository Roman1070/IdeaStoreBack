package main

import (
	"encoding/json"
	"fmt"
	chatsv1 "idea-store-auth/gen/go/chats"
	"idea-store-auth/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatsClient struct {
	api chatsv1.ChatsClient
	//producer *kafka.Producer
	clients sync.Map
}

func NewChatsClient(addr string, timeout time.Duration, retriesCount int) (*ChatsClient, error) {
	retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
		grpcretry.UnaryClientInterceptor(retryOptions...),
	))

	if err != nil {
		slog.Error("client NewChatsClient error: " + err.Error())
		return nil, fmt.Errorf("client NewChatsClient error: " + err.Error())
	}
	/*producer, err := mykafka.StartProducer()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}*/
	return &ChatsClient{
		api: chatsv1.NewChatsClient(cc),
		//	producer: producer,
		clients: sync.Map{},
	}, nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all connections
}

func (c *ChatsClient) HandleChatWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("client HandleChatWebSocket error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	defer func() {
		c.clients.Delete(ws)
		ws.Close()
	}()
	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		slog.Error("client HandleChatWebSocket error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}
	c.clients.Store(ws, userId)
	for {
		// Read message from browser
		_, msg, err := ws.ReadMessage()

		if err != nil {
			slog.Error("HandleChatWebSocket error: " + err.Error())
			utils.WriteError(w, "Internal error")
			return
		}

		fmt.Println(string(msg))
		type recieverIdWrapper struct {
			ReceiverId int64 `json:"reciever_id"`
		}
		var reciever recieverIdWrapper
		err = json.Unmarshal(msg, &reciever)
		if err != nil {
			slog.Error("HandleChatWebSocket error: " + err.Error())
			utils.WriteError(w, "Internal error")
			return
		}

		c.clients.Range(func(client, id interface{}) bool {
			if id == reciever.ReceiverId {
				if err := client.(*websocket.Conn).WriteMessage(websocket.TextMessage, msg); err != nil {
					fmt.Println("HandleChatWebSocket WriteMessage error:", err)
					client.(*websocket.Conn).Close()
					c.clients.Delete(client)
					return false
				}
			}
			return true
		})
	}
}

func (c *ChatsClient) SendMessage(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		slog.Error("client SendMessage error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	type request struct {
		RecieverId int64  `json:"recieverId"`
		Text       string `json:"text,omitempty"`
		FileName   string `json:"fileName,omitempty"`
		IdeaId     string `json:"ideaId,omitempty"`
	}

	var req request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("client SendMessage error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	creationDate := time.Now()
	creationDateStr := fmt.Sprintf("%02d.%02d.%04d %02d:%02d:%02d", creationDate.Day(), creationDate.Month(), creationDate.Year(), creationDate.Hour(), creationDate.Minute(), creationDate.Second())
	ideaId := int64(-1)
	if req.IdeaId != "" {
		ideaId, err = strconv.ParseInt(req.IdeaId, 10, 64)
		if err != nil {
			slog.Error("client SendMessage error: " + err.Error())
			utils.WriteError(w, "Internal error")
			return
		}
	}

	resp, err := c.api.SendMessage(r.Context(), &chatsv1.SendMessageRequest{
		Data: &chatsv1.MessageData{
			SenderId:    userId,
			RecieverId:  req.RecieverId,
			Text:        req.Text,
			FileName:    req.FileName,
			SendingDate: creationDateStr,
			IdeaId:      ideaId,
		},
	})
	if err != nil {
		slog.Error("client SendMessage error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}
	//mykafka.OnMessageSent(c.producer, resp.Id, req.RecieverId)
	result, err := json.Marshal(resp)
	if err != nil {
		slog.Error("client SendMessage error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (c *ChatsClient) GetMessages(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		slog.Error("client GetMessages error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	secondIdStr := r.URL.Query().Get("id")
	secondId, err := strconv.ParseInt(secondIdStr, 10, 64)
	if err != nil {
		slog.Error("client GetMessages error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	resp, err := c.api.GetMessages(r.Context(), &chatsv1.GetMessagesRequest{
		FirstId:  userId,
		SecondId: secondId,
	})
	if err != nil {
		slog.Error("client GetMessages error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		slog.Error("client GetMessages error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func (c *ChatsClient) CreateChat(w http.ResponseWriter, r *http.Request) {

}

func (c *ChatsClient) GetChats(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		slog.Error("client GetChats error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	resp, err := c.api.GetUsersChats(r.Context(), &chatsv1.GetUsersChatsRequest{
		UserId: userId,
	})
	if err != nil {
		slog.Error("client GetChats error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		slog.Error("client GetChats error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
