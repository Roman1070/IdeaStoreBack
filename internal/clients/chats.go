package main

import (
	"encoding/json"
	"fmt"
	chatsv1 "idea-store-auth/gen/go/chats"
	"idea-store-auth/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	mykafka "idea-store-auth/internal/kafka"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/websocket"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatsClient struct {
	api      chatsv1.ChatsClient
	producer *kafka.Producer
	clients  map[*websocket.Conn]bool
}

func NewChatsClient(addr string, timeout time.Duration, retriesCount int) (*ChatsClient, error) {
	const op = "client.boards.New"
	clients := make(map[*websocket.Conn]bool) // Track active clients
	retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
		grpcretry.UnaryClientInterceptor(retryOptions...),
	))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	producer, err := mykafka.StartProducer()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &ChatsClient{
		api:      chatsv1.NewChatsClient(cc),
		producer: producer,
		clients:  clients,
	}, nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all connections
}

func (c *ChatsClient) HandleChatWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	slog.Info("web socket recieved ")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		utils.WriteError(w, err.Error())
		return
	}
	defer func() {
		delete(c.clients, ws)
		ws.Close()
	}()

	c.clients[ws] = true
	for {
		// Read message from browser
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			utils.WriteError(w, err.Error())
			break
		}
		fmt.Printf("Received: %s\n", msg)

		fmt.Printf("ws clients: %v\n", c.clients)
		for client := range c.clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("broadcast error:", err)
				client.Close()
				delete(c.clients, client)
			}
		}
	}
}

func (c *ChatsClient) SendMessage(w http.ResponseWriter, r *http.Request) {

	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	checkChatExistance := r.URL.Query().Get("checkChatExistance") == "true"
	type request struct {
		RecieverId int64  `json:"recieverId"`
		Text       string `json:"text"`
		FileName   string `json:"fileName"`
	}
	var req request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("client SendMessage error: %v\n", err.Error())
		utils.WriteError(w, err.Error())
		return
	}
	creationDate := time.Now()
	creationDateStr := fmt.Sprintf("%02d.%02d.%04d %02d:%02d:%02d", creationDate.Day(), creationDate.Month(), creationDate.Year(), creationDate.Hour(), creationDate.Minute(), creationDate.Second())

	resp, err := c.api.SendMessage(r.Context(), &chatsv1.SendMessageRequest{
		Data: &chatsv1.MessageData{
			SenderId:           userId,
			RecieverId:         req.RecieverId,
			Text:               req.Text,
			FileName:           req.FileName,
			SendingDate:        creationDateStr,
			CheckChatExistance: checkChatExistance,
		},
	})
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	mykafka.OnMessageSent(c.producer, resp.Id, req.RecieverId)
	result, err := json.Marshal(resp)

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (c *ChatsClient) GetMessages(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	secondIdStr := r.URL.Query().Get("id")
	secondId, err := strconv.ParseInt(secondIdStr, 10, 64)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	resp, err := c.api.GetMessages(r.Context(), &chatsv1.GetMessagesRequest{
		FirstId:  userId,
		SecondId: secondId,
	})

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	result, err := json.Marshal(resp)

	if err != nil {
		utils.WriteError(w, err.Error())
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
		utils.WriteError(w, err.Error())
		return
	}
	resp, err := c.api.GetUsersChats(r.Context(), &chatsv1.GetUsersChatsRequest{
		UserId: userId,
	})
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	result, err := json.Marshal(resp)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
