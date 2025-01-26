package main

import (
	"encoding/json"
	"fmt"
	chatsv1 "idea-store-auth/gen/go/chats"
	"idea-store-auth/internal/utils"
	"net/http"
	"strconv"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatsClient struct {
	api chatsv1.ChatsClient
}

func NewChatsClient(addr string, timeout time.Duration, retriesCount int) (*ChatsClient, error) {
	const op = "client.boards.New"

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
	return &ChatsClient{
		api: chatsv1.NewChatsClient(cc),
	}, nil
}

func (c *ChatsClient) SendMessage(w http.ResponseWriter, r *http.Request) {

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
	fmt.Println(resp)
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
