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

const (
	emptyValue = -1
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appId int,
	) (token string, err error)

	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userId int64, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
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
	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
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
func (s *serverAPI) IsAdmin(ctx context.Context, req *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {
	if err := validateIsAdmin(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}
	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &authv1.IsAdminResponse{IsAdmin: isAdmin}, nil
}
func validateLogin(req *authv1.LoginRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email must not be empty")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password must not be empty")
	}

	if req.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "app id is required")
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

func validateIsAdmin(req *authv1.IsAdminRequest) error {
	if req.GetUserId() == emptyValue {
		return status.Error(codes.InvalidArgument, "used id is required")
	}
	return nil
}
