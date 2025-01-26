package sqlite

import (
	"database/sql"
	"fmt"
	common "idea-store-auth/cmd"
	boardsv1 "idea-store-auth/gen/go/boards"
	ideasv1 "idea-store-auth/gen/go/idea"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/config"
	"strconv"
	"strings"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Storage struct {
	db             *sql.DB
	ideasClient    ideasv1.IdeasClient
	boardsClient   boardsv1.BoardsClient
	profilesClient profilesv1.ProfilesClient
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

	boardsClient, err := grpc.NewClient(common.GrpcBoardsAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(retryOptions...),
		))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	profilesClient, err := grpc.NewClient(common.GrpcProfilesAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(retryOptions...),
		))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db, ideasClient: ideasv1.NewIdeasClient(ideasClient),
		boardsClient:   boardsv1.NewBoardsClient(boardsClient),
		profilesClient: profilesv1.NewProfilesClient(profilesClient)}, nil
}

func ParseIdsSqlite(str string) ([]int64, error) {
	if len(str) == 0 {
		return []int64{}, nil
	}
	slice := strings.Split(str, " ")
	var ids []int64
	for _, i := range slice {
		val, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing id %v", i)
		}
		ids = append(ids, val)
	}
	return ids, nil
}

func ParseIdPairs(str string) ([]ideaBoardPair, error) {
	if len(str) == 0 {
		return []ideaBoardPair{}, nil
	}
	slice := strings.Split(str, " ")
	var pairs []ideaBoardPair
	for _, i := range slice {
		val, err := parseIdeaBoardPair(i)
		if err != nil {
			return nil, fmt.Errorf("error parsing savedIdeas str %v", i)
		}
		pairs = append(pairs, val)
	}
	return pairs, nil
}
