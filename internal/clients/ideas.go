package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/utils"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	FilesPath = "/app/files/"
)

type IdeasClient struct {
	api ideasv1.IdeasClient
}
type getIdeaRequest struct {
	Id int64 `json:"id"`
}

func (c *IdeasClient) GetIdea(w http.ResponseWriter, r *http.Request) {
	slog.Info("Client started to get idea")

	var req getIdeaRequest
	var err error
	req.Id, err = strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	resp, err := c.api.GetIdea(r.Context(), &ideasv1.GetRequest{IdeaId: req.Id})

	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}
	m := protojson.MarshalOptions{EmitDefaultValues: true}

	result, err := m.Marshal(resp)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func GetImages(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile(r.RequestURI)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(file)
}

func (c *IdeasClient) Create(w http.ResponseWriter, r *http.Request) {
	slog.Info("Client started to create")

	userId, err := GetUserIdByRequestWithCookie(r)
	if err != nil {
		slog.Error("Error parsing JWT: " + err.Error())
		utils.WriteError(w, "Error parsing JWT: "+err.Error())
		return
	}

	r.ParseMultipartForm(12 << 20)
	defer r.Body.Close()
	file, h, err := r.FormFile("image")
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}
	defer file.Close()
	ext := filepath.Ext(h.Filename)
	hash := md5.Sum([]byte(h.Filename))
	path := FilesPath + hex.EncodeToString(hash[:]) + ext
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

	request := &ideasv1.CreateRequest{
		Image:       hex.EncodeToString(hash[:]) + ext,
		Name:        r.Form.Get("name"),
		Description: r.Form.Get("description"),
		Link:        r.Form.Get("link"),
		Tags:        r.Form.Get("tags"),
		UserId:      userId,
	}

	createResponse, err := c.api.CreateIdea(r.Context(), request)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	result, err := json.Marshal(createResponse)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func (c *IdeasClient) GetAllIdeas(w http.ResponseWriter, r *http.Request) {
	slog.Info("Client started to get ideas")

	userId, _ := GetUserIdByRequestWithCookie(r)
	resp, err := c.api.GetAllIdeas(r.Context(), &ideasv1.GetAllRequest{
		UserId: userId,
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
func (c *IdeasClient) GetIdeas(w http.ResponseWriter, r *http.Request) {
	slog.Info("Client started to get ideas")
	type ideasIds struct {
		Ids []string `json:"ids"`
	}
	idsStringSlice := ideasIds{}
	err := json.NewDecoder(r.Body).Decode(&idsStringSlice)

	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	ids := []int64{}
	for _, s := range idsStringSlice.Ids {
		id, err := strconv.ParseInt(s, 10, 64)

		if err != nil {
			utils.WriteError(w, err.Error())
			return
		}
		ids = append(ids, id)
	}
	resp, err := c.api.GetIdeas(r.Context(), &ideasv1.GetIdeasRequest{
		Ids: ids,
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
func NewIdeasClient(addr string, timeout time.Duration, retriesCount int) (*IdeasClient, error) {
	const op = "client.ideas.New"

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
	return &IdeasClient{
		api: ideasv1.NewIdeasClient(cc),
	}, nil
}
