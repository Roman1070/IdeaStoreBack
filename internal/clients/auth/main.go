package main

import (
	"context"
	"fmt"
	authv1 "idea-store-auth/gen/go/auth"
	"idea-store-auth/internal/config"
	"net"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcHost = "localhost"

type Client struct {
	api authv1.AuthClient
}

const appID = int32(2)
const clientAddr = "localhost:8181"

func main() {
	http.HandleFunc("/register", Regsiter)
	http.HandleFunc("/login", Login)
	fmt.Println("Server is listening...")
	http.ListenAndServe(clientAddr, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	cfg := config.MustLoad()
	authClient, err := New(grpcAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %v", "Error starting client", err.Error())))
		return
	}
	request := &authv1.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    appID,
	}
	loginResponse, err := (*authClient).api.Login(context.Background(), request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %v", "Error  during login", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(loginResponse.GetToken()))
}

func Regsiter(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	cfg := config.MustLoad()

	authClient, err := New(grpcAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %v", "Error starting client", err.Error())))
		return
	}
	request := &authv1.RegisterRequest{
		Email:    email,
		Password: password,
	}
	registerResponse, err := authClient.api.Register(r.Context(), request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %v", "Error  during register", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(int(registerResponse.UserId))))
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
