package main

import (
	"encoding/json"
	"fmt"
	profilesv1 "idea-store-auth/gen/go/profiles"
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

type ProfilesClient struct {
	api profilesv1.ProfilesClient
}

func NewProfilesClient(addr string, timeout time.Duration, retriesCount int) (*ProfilesClient, error) {
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
	return &ProfilesClient{
		api: profilesv1.NewProfilesClient(cc),
	}, nil
}

func (c *ProfilesClient) CreateProfile(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Id    int64  `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	var req request

	json.NewDecoder(r.Body).Decode(&req)
	_, err := c.api.CreateProfile(r.Context(), &profilesv1.CreateProfileRequest{
		Id:    req.Id,
		Email: req.Email,
		Name:  req.Name,
	})
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error creatingprofile: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ProfilesClient) GetProfile(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)

	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error getting profile: "+err.Error())
		return
	}
	resp, err := c.api.GetProfile(r.Context(), &profilesv1.GetProfileRequest{
		Id: userId,
	})

	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error getting profile: "+err.Error())
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

func (c *ProfilesClient) ToggleSaveIdea(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	boardId, err := strconv.ParseInt(r.URL.Query().Get("board_id"), 10, 64)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	ideaId, err := strconv.ParseInt(r.URL.Query().Get("idea_id"), 10, 64)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	resp, err := c.api.ToggleSaveIdea(r.Context(), &profilesv1.ToggleSaveRequest{
		UserId:  userId,
		BoardId: boardId,
		IdeaId:  ideaId,
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

func (c *ProfilesClient) IsIdeaSaved(w http.ResponseWriter, r *http.Request) {

	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	ideaId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	resp, err := c.api.IsIdeaSaved(r.Context(), &profilesv1.IsIdeaSavedRequest{
		UserId: userId,
		IdeaId: ideaId,
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

func (c *ProfilesClient) GetSavedIdeas(w http.ResponseWriter, r *http.Request) {

	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	resp, err := c.api.GetSavedIdeas(r.Context(), &profilesv1.GetSavedIdeasRequest{
		UserId: userId,
	})
	if err != nil {
		slog.Error("c.api.GetSavedIdeas err: " + err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
