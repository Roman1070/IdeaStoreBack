package main

import (
	"encoding/json"
	"fmt"
	boardsv1 "idea-store-auth/gen/go/boards"
	"idea-store-auth/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

type BoardsClient struct {
	api boardsv1.BoardsClient
}

func NewBoardsClient(addr string, timeout time.Duration, retriesCount int) (*BoardsClient, error) {
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
	return &BoardsClient{
		api: boardsv1.NewBoardsClient(cc),
	}, nil
}

func (c *BoardsClient) CreateBoard(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name string `json:"name"`
	}
	const op = "client.boards.New"
	userId, err := GetUserIdByRequestWithCookie(r)

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	log := slog.With(slog.String("op", op))

	var req request
	err = json.NewDecoder(r.Body).Decode(&req)
	
	if err!=nil{
		utils.WriteError(w,err.Error())
		return
	}
	log.Info(fmt.Sprintf("started to create a board, id = %v", req.Name))

	resp, err := c.api.CreateBoard(r.Context(), &boardsv1.CreateBoardRequest{
		Name:   req.Name,
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
func (c *BoardsClient) GetBoard(w http.ResponseWriter, r *http.Request) {
	const op = "client.boards.Get"

	log := slog.With(slog.String("op", op))

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	log.Info(fmt.Sprintf("started to get a board, id = %v", id))

	resp, err := c.api.GetBoard(r.Context(), &boardsv1.GetBoardRequest{
		Id: id,
	})
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	m := protojson.MarshalOptions{EmitDefaultValues: true}

	result, err := m.Marshal(resp)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
func (c *BoardsClient) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	const op = "client.boards.GetAll"
	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	log := slog.With(slog.String("op", op))

	log.Info("started to get aall boards")

	resp, err := c.api.GetAllBoards(r.Context(), &boardsv1.GetAllBoardsRequest{
		UserId: userId,
	})
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	m := protojson.MarshalOptions{EmitDefaultValues: true}

	result, err := m.Marshal(resp)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (c *BoardsClient) GetIdeasInBoard(w http.ResponseWriter, r *http.Request) {
	boardId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	resp, err := c.api.GetIdeasInBoard(r.Context(), &boardsv1.GetIdeasInBoardRequest{
		BoardId: boardId,
	})
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	result, err := json.Marshal(resp.Ideas)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func(c *BoardsClient) DeleteBoard(w http.ResponseWriter, r *http.Request){
	
	type request struct{
		Id string `json:"id"`
	}
	userId,err := GetUserIdByRequestWithCookie(r)
	if err!=nil{
		utils.WriteError(w,err.Error())
		return
	}
	var req request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err!=nil{
		utils.WriteError(w,err.Error())
		slog.Error(err.Error())
		return
	}
	id, err := strconv.ParseInt(req.Id,10,64)
	if err!=nil{
		utils.WriteError(w,err.Error())
		slog.Error(err.Error())
		return
	}
	_, err = c.api.DeleteBoard(r.Context(),&boardsv1.DeleteBoardRequest{
		BoardId: id,
		UserId: userId,
	})
	if err!=nil{
		utils.WriteError(w,err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}