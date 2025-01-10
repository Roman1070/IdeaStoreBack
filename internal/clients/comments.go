package main

import (
	"encoding/json"
	"fmt"
	commentsv1 "idea-store-auth/gen/go/comments"
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

type CommentsClient struct {
	api commentsv1.CommentsClient
}

func NewCommentsClient(addr string, timeout time.Duration, retriesCount int) (*CommentsClient, error) {
	const op = "client.comments.New"

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
	return &CommentsClient{
		api: commentsv1.NewCommentsClient(cc),
	}, nil
}

func (c *CommentsClient) CreateComment(w http.ResponseWriter, r *http.Request) {
	slog.Info("Client started CreateComment")
	var req commentsv1.CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	id, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	req.UserId = id
	resp, err := c.api.CreateComment(r.Context(), &req)

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

func (c *CommentsClient) GetComments(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("idea")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	resp, err := c.api.GetComments(r.Context(), &commentsv1.GetCommentsRequest{IdeaId: id})

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
