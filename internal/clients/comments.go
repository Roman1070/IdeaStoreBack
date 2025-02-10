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
	retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
		grpcretry.UnaryClientInterceptor(retryOptions...),
	))
	if err != nil {
		slog.Error("Client NewCommentsClient error : " + err.Error())
		return nil, fmt.Errorf("Client NewCommentsClient error : " + err.Error())
	}

	return &CommentsClient{
		api: commentsv1.NewCommentsClient(cc),
	}, nil
}

func (c *CommentsClient) CreateComment(w http.ResponseWriter, r *http.Request) {
	slog.Info("Client started CreateComment")
	type request struct {
		UserId       int64
		IdeaId       string `json:"idea_id"`
		Text         string `json:"text"`
		CreationDate string
	}

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteError(w, "Internal error")
		slog.Error("client CreateComment error : " + err.Error())
		return
	}

	ideaId, err := strconv.ParseInt(req.IdeaId, 10, 64)
	if err != nil {
		utils.WriteError(w, "Internal error")
		slog.Error("client CreateComment error : " + err.Error())
		return
	}

	id, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		utils.WriteError(w, "Internal error")
		slog.Error("client CreateComment error : " + err.Error())
		return
	}

	time := time.Now()

	req.CreationDate = fmt.Sprintf("%02d.%02d.%04d", time.Day(), time.Month(), time.Year())
	fmt.Println(req)
	resp, err := c.api.CreateComment(r.Context(), &commentsv1.CreateCommentRequest{
		IdeaId:       ideaId,
		UserId:       id,
		Text:         req.Text,
		CreationDate: req.CreationDate,
	})

	if err != nil {
		utils.WriteError(w, "Internal error")
		slog.Error("client CreateComment error : " + err.Error())
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
		utils.WriteError(w, "Internal error")
		slog.Error("client GetComments error : " + err.Error())
		return
	}

	resp, err := c.api.GetComments(r.Context(), &commentsv1.GetCommentsRequest{IdeaId: id})
	if err != nil {
		utils.WriteError(w, "Internal error")
		slog.Error("client GetComments error : " + err.Error())
		return
	}

	m := protojson.MarshalOptions{EmitDefaultValues: true}
	result, err := m.Marshal(resp)
	if err != nil {
		utils.WriteError(w, "Internal error")
		slog.Error("client GetComments error : " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
