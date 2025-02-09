package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/utils"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	_, err = c.api.CreateProfile(r.Context(), &profilesv1.CreateProfileRequest{
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

func (c *ProfilesClient) ToggleLikeIdea(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)

	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error ToggleLikeIdea: "+err.Error())
		return
	}
	type request struct {
		IdeaIdStr string `json:"ideaId"`
	}
	var req request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error ToggleLikeIdea: "+err.Error())
		return
	}
	ideaId, err := strconv.ParseInt(req.IdeaIdStr, 10, 64)

	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error ToggleLikeIdea: "+err.Error())
		return
	}
	resp, err := c.api.ToggleLikeIdea(r.Context(), &profilesv1.ToggleLikeIdeaRequest{
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
func (c *ProfilesClient) IsIdeaLiked(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error IsIdeaLiked: "+err.Error())
		return
	}

	ideaIdStr := r.URL.Query().Get("id")
	ideaId, err := strconv.ParseInt(ideaIdStr, 10, 64)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error IsIdeaLiked: "+err.Error())
		return
	}

	resp, err := c.api.IsIdeaLiked(r.Context(), &profilesv1.IsIdeaLikedRequest{
		UserId: userId,
		IdeaId: ideaId,
	})
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error IsIdeaLiked: "+err.Error())
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
func (c *ProfilesClient) GetCurrentProfile(w http.ResponseWriter, r *http.Request) {
	authorized := true

	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		if !strings.Contains(err.Error(), NoCookieError) {
			slog.Error("Error getting profile: " + err.Error())
			utils.WriteError(w, "Error getting profile: "+err.Error())
			return
		} else {
			authorized = false
		}
	}

	var resp *profilesv1.GetProfileResponse
	if authorized {
		resp, err = c.api.GetProfile(r.Context(), &profilesv1.GetProfileRequest{
			Id: userId,
		})

		if err != nil {
			slog.Error(err.Error())
			utils.WriteError(w, "Error getting profile: "+err.Error())
			return
		}
	} else {
		resp = &profilesv1.GetProfileResponse{
			Data: &profilesv1.ProfileData{
				Id: -1,
			},
		}
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
func (c *ProfilesClient) GetProfile(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, "Error getting profile: "+err.Error())
		return
	}
	resp, err := c.api.GetProfile(r.Context(), &profilesv1.GetProfileRequest{
		Id: id,
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
func (c *ProfilesClient) GetProfilesFromSearch(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	resp, err := c.api.GetProfilesFromSearch(r.Context(), &profilesv1.GetProfilesFromSearchRequest{
		Input: input,
	})

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	result, err := json.Marshal(&resp.Profiles)

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func (c *ProfilesClient) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserIdByRequestWithCookie(r)

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	avatar := ""
	r.ParseMultipartForm(12 << 20)
	defer r.Body.Close()
	file, h, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		ext := filepath.Ext(h.Filename)
		hash := md5.Sum([]byte(h.Filename))
		path := "/app/files/" + hex.EncodeToString(hash[:]) + ext
		tmpfile, err := os.Create(path)
		if err != nil {
			slog.Error(err.Error())
			utils.WriteError(w, err.Error())
			return
		}
		defer tmpfile.Close()
		_, err = io.Copy(tmpfile, file)
		if err != nil {
			slog.Error(err.Error())
			utils.WriteError(w, err.Error())
			return
		}
		avatar = hex.EncodeToString(hash[:]) + ext
		fmt.Printf("starting to save avatar: %v\n", avatar)
	} else {
		fmt.Printf("error parsing avatar: %v\n", err.Error())
	}
	request := &profilesv1.UpdateProfileRequest{
		Avatar:      avatar,
		Name:        r.Form.Get("name"),
		Description: r.Form.Get("description"),
		Link:        r.Form.Get("link"),
		UserId:      userId,
	}

	resp, err := c.api.UpdateProfile(r.Context(), request)

	if err != nil {
		slog.Error(err.Error())
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
	slog.Info("client started GetSavedIdeas")
	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		slog.Error("cliet GetSavedIdeas error: " + err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		slog.Error("client GetSavedIdeas error: " + err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	offsetStr := r.URL.Query().Get("offset")
	offset, err := strconv.ParseInt(offsetStr, 10, 32)
	if err != nil {
		slog.Error("client GetSavedIdeas error: " + err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	resp, err := c.api.GetSavedIdeas(r.Context(), &profilesv1.GetSavedIdeasRequest{
		UserId: userId,
		Limit:  int32(limit),
		Offset: int32(offset),
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
