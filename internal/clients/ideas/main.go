package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/config"
	"idea-store-auth/internal/utils"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
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
	router.HandleFunc("/get-ideas", ideasClient.GetAllIdeas).Methods("GET","OPTIONS")
	fmt.Println("Server is listening...")
	
	corsHandler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(clientAddr, corsHandler))
}

func (c *IdeasClient) Create(w http.ResponseWriter, r *http.Request) {
	slog.Info("Client started to create")

	r.ParseMultipartForm(12 << 20)
	file, h, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ext:= filepath.Ext(h.Filename)
	hash :=md5.Sum([]byte(h.Filename))
	path:= "F:/Roman/WEB/IdeaStoreFront/Images/"+hex.EncodeToString(hash[:])+ext
	tmpfile, err := os.Create(path)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tmpfile.Close()
	
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	var req createRequest
	json.NewDecoder(r.Body).Decode(&req)
	request := &ideasv1.CreateRequest{
		Image:       path,
		Name:        req.Name,
		Description: req.Description,
		Link:        req.Link,
		Tags:        req.Tags,
	}
	slog.Info(req.Image)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(result)
}
func (c *IdeasClient) GetAllIdeas(w http.ResponseWriter, r *http.Request){
	slog.Info("Client started to get ideas")
	resp,err:= c.api.GetAllIdeas(r.Context(),&emptypb.Empty{})
	if err!=nil{
		utils.WriteError(w,err.Error())
		return
	}
	result, err := json.Marshal(resp.Ideas)
	if err!=nil{
		utils.WriteError(w,err.Error())
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