package auth

import (
	"context"
	"errors"
	authv1 "idea-store-auth/gen/go/auth"
	"idea-store-auth/internal/services/auth"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
	) (token string, err error)

	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userId int64, err error)
}
type serverAPI struct {
	authv1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	authv1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	slog.Info("started to login")
	if err := validateLogin(req); err != nil {
		return nil, err
	}
	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "Invalid credentials")
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}
	resp := &authv1.LoginResponse{Token: token}
	return resp, nil
}

func (s *serverAPI) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	email := req.GetEmail()
	password := req.GetPassword()

	if err := validateRegister(req); err != nil {
		return nil, err
	}

	userID, err := s.auth.RegisterNewUser(ctx, email, password)

	if err != nil {
		if errors.Is(err, auth.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "User already exists")
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &authv1.RegisterResponse{
		UserId: userID,
	}, nil
}
func validateLogin(req *authv1.LoginRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email must not be empty")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password must not be empty")
	}
	return nil
}

func validateRegister(req *authv1.RegisterRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email must not be empty")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password must not be empty")
	}
	return nil
}
