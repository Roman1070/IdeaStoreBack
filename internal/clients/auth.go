package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	authv1 "idea-store-auth/gen/go/auth"
	"idea-store-auth/internal/utils"
	"net/http"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)
type AuthClient struct {
	authAPi authv1.AuthClient
}
const appID = 2

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
		utils.WriteError(w,"Invalid credentials")
					return
		}
		utils.WriteError(w,err.Error())
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

	err:=json.NewDecoder(r.Body).Decode(&req)
	
	if err!=nil{
		utils.WriteError(w,err.Error())
		return
	}
	request := &authv1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	registerResponse, err := c.authAPi.Register(r.Context(), request)
	if err != nil {
		if errors.Is(err, status.Error(codes.AlreadyExists, "User already exists")) {
			utils.WriteError(w,"User already exists")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %v", "Error  during register", err.Error())))
		return
	}

	result, err := json.Marshal(registerResponse)
	if err != nil {
		utils.WriteError(w,"Error masrhalling response")
		return
	}
	w.WriteHeader(http.StatusOK) 
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	w.Write(result)
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