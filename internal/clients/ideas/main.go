package main

import (
	"encoding/json"
	"fmt"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/config"
	"idea-store-auth/internal/utils"
	"log"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type IdeasClient struct {
	api ideasv1.IdeasClient
}
const grpcHost = "localhost"

const clientAddr = "localhost:8182"
type createRequest struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Tags        string `json:"tags"`
}

func main() {
	cfg := config.MustLoad()
	ideasClient, _ := NewIdeasClient(grpcAddress(cfg), cfg.Clients.Ideas.Timeout, cfg.Clients.Ideas.RetriesCount)
	
	router := mux.NewRouter()
	router.HandleFunc("/create-pin", ideasClient.Create).Methods("POST","OPTIONS")
	fmt.Println("Server is listening...")
	
	corsHandler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(clientAddr, corsHandler))
}

func (c *IdeasClient) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	slog.Info("Client started to create")
	json.NewDecoder(r.Body).Decode(&req)
	request := &ideasv1.CreateRequest{
		Image:       req.Image,
		Name:        req.Name,
		Description: req.Description,
		Link:        req.Link,
		Tags:        req.Tags,
	}
	slog.Info(req.Name)
	createResponse, err := c.api.Create(r.Context(), request)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(result)
}

func NewIdeasClient(addr string, timeout time.Duration, retriesCount int) (*IdeasClient, error) {
	const op = "client.auth.New"

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
func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.IdeasMS.Port))
}