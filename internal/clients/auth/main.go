package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	authv1 "idea-store-auth/gen/go/auth"
	"idea-store-auth/internal/config"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const grpcHost = "localhost"

type Client struct {
	api authv1.AuthClient
}
type ErrorWrapper struct{
	Err string `json:"err"`
}
const appID = int32(2)
const clientAddr = "localhost:8181"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/register", Regsiter).Methods("POST","OPTIONS")
	router.HandleFunc("/login", Login).Methods("GET","OPTIONS")
	fmt.Println("Server is listening...")
	
	corsHandler := cors.Default().Handler(router)

	// Запускаем сервер с обработчиком Cors
	log.Fatal(http.ListenAndServe(clientAddr, corsHandler))
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	cfg := config.MustLoad()
	authClient, err := New(grpcAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)

	if err != nil {
		err := ErrorWrapper{Err: err.Error()}
		w.WriteHeader(http.StatusOK)		
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json,_:=json.Marshal(err)
		w.Write(json)
		return
	}
	request := &authv1.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    appID,
	}
	loginResponse, err := (*authClient).api.Login(context.Background(), request)
	if err != nil {
		if(errors.Is(err, status.Error(codes.InvalidArgument,"Invalid credentials"))){
			err := ErrorWrapper{Err: "Invalid credentials"}
			w.WriteHeader(http.StatusOK)		
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json,_:=json.Marshal(err)
			w.Write(json)
			return
		}
		errWrapper := ErrorWrapper{Err: err.Error()}
		w.WriteHeader(http.StatusOK)		
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json,_:=json.Marshal(errWrapper)
		w.Write(json)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	responseJson, _ := json.Marshal(loginResponse)
	w.Write(responseJson)
}

func Regsiter(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	cfg := config.MustLoad()

	authClient, err := New(grpcAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)

	if err != nil {
		err := ErrorWrapper{Err: err.Error()}
		w.WriteHeader(http.StatusOK)		
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json,_:=json.Marshal(err)
		w.Write(json)
		return
	}
	request := &authv1.RegisterRequest{
		Email:    email,
		Password: password,
	}
	registerResponse, err := authClient.api.Register(r.Context(), request)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "User already exists")) {
			err := ErrorWrapper{Err: "User already exists"}
			w.WriteHeader(http.StatusOK)		
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json,_:=json.Marshal(err)
			w.Write(json)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %v", "Error  during register", err.Error())))
		return
	}

	result, err := json.Marshal(registerResponse)
	if err != nil {
		err := ErrorWrapper{Err: "Error marshaling response"}
		w.WriteHeader(http.StatusOK)		
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json,_:=json.Marshal(err)
		w.Write(json)
		return
	}
	w.WriteHeader(http.StatusOK) 
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	w.Write(result)
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}

func New(addr string, timeout time.Duration, retriesCount int) (*Client, error) {
	const op = "client.auth.New"

	/*retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}*/

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Client{
		api: authv1.NewAuthClient(cc),
	}, nil
}
