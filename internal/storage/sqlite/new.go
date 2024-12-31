package sqlite

import (
	"database/sql"
	"fmt"
	common "idea-store-auth/cmd"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/config"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Storage struct {
	db          *sql.DB
	ideasClient ideasv1.IdeasClient
}

const emptyValue = -1

func New(storagePath string) (*Storage, error) {
	const op = "sqlite.New"
	cfg := config.MustLoad()

	retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(5)),
		grpcretry.WithPerRetryTimeout(5 * time.Second),
	}
	db, err := sql.Open("sqlite3", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	ideasClient, err := grpc.NewClient(common.GrpcIdeasAddress(cfg), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
		grpcretry.UnaryClientInterceptor(retryOptions...),
	))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db, ideasClient: ideasv1.NewIdeasClient(ideasClient)}, nil
}
