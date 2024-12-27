package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	authv1 "idea-store-auth/gen/go/auth"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/config"
	"log"
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
	"google.golang.org/grpc/status"
)

const grpcHost = "localhost"

type AuthClient struct {
	authAPi authv1.AuthClient
}
type IdeasClient struct {
	api ideasv1.IdeasClient
}
type ErrorWrapper struct{
	Err string `json:"err"`
}
const appID = int32(2)
const clientAddr = "localhost:8181"

func main() {
	cfg := config.MustLoad()
	authClient, _ := NewAuthClient(grpcAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	ideasClient, _ := NewIdeasClient(grpcAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	
	router := mux.NewRouter()
	router.HandleFunc("/register", authClient.Regsiter).Methods("POST","OPTIONS")
	router.HandleFunc("/login", authClient.Login).Methods("POST","OPTIONS")
	router.HandleFunc("/create-pin", ideasClient.Create).Methods("POST","OPTIONS")
	fmt.Println("Server is listening...")
	
	corsHandler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(clientAddr, corsHandler))
}

func (c *AuthClient) Login(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	json.NewDecoder(r.Body).Decode(&req)
	request := &authv1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
		AppId:    appID,
	}

	loginResponse, err := c.authAPi.Login(context.Background(), request)
	if err != nil {
		if(errors.Is(err, status.Error(codes.InvalidArgument,"Invalid credentials"))){
		writeError(w,"Invalid credentials")
					return
		}
		writeError(w,err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	responseJson, _ := json.Marshal(loginResponse)
	w.Write(responseJson)
}

type registerRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


func(c *AuthClient)  Regsiter(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	json.NewDecoder(r.Body).Decode(&req)
	request := &authv1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	registerResponse, err := c.authAPi.Register(r.Context(), request)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "User already exists")) {
			writeError(w,"User already exists")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %v", "Error  during register", err.Error())))
		return
	}

	result, err := json.Marshal(registerResponse)
	if err != nil {
		writeError(w,"Error masrhalling response")
		return
	}
	w.WriteHeader(http.StatusOK) 
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	w.Write(result)
}
type createRequest struct{
	Image string `json:"image"`
	Name string `json:"name"`
	Description string `json:"description"`
	Link string `json:"link"`
	Tags string `json:"tags"`
}
func (c *IdeasClient) Create(w http.ResponseWriter, r *http.Request){
	var req createRequest

	json.NewDecoder(r.Body).Decode(&req)
	request := &ideasv1.CreateRequest{
		Image: req.Image,
		Name: req.Name,
		Description: req.Description,
		Link: req.Link,
		Tags: req.Tags,
	}
	createResponse,err := c.api.Create(r.Context(),request)
	if err!=nil{
		writeError(w,err.Error())
	}
	
	result, err := json.Marshal(createResponse)
	if err!=nil{
		writeError(w,err.Error())
	}
	w.WriteHeader(http.StatusOK) 
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	w.Write(result)
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}

func NewAuthClient(addr string, timeout time.Duration, retriesCount int) (*AuthClient, error) {
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
	return &AuthClient{
		authAPi: authv1.NewAuthClient(cc),
	}, nil
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

func writeError(w http.ResponseWriter, err string){
	errWrapper := ErrorWrapper{Err: err}
	w.WriteHeader(http.StatusOK)		
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json,_:=json.Marshal(errWrapper)
	w.Write(json)
}