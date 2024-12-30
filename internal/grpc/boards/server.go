package boards

import (
	"context"
	boardsv1 "idea-store-auth/gen/go/boards"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Boards interface {
	CreateBoard(ctx context.Context, name string) (int64,error)
	GetBoard(ctx context.Context, id int64) (models.Board, error)
	GetAllBoards(ctx context.Context,e *emptypb.Empty) ([]*boardsv1.BoardData, error)
}

type serverAPI struct {
	boardsv1.UnimplementedBoardsServer
	boards Boards
}

func Register(gRPC *grpc.Server, boards Boards) {
	boardsv1.RegisterBoardsServer(gRPC, &serverAPI{boards: boards})
}

func (s *serverAPI) CreateBoard(ctx context.Context,req *boardsv1.CreateBoardRequest) (*boardsv1.CreateBoardResponse, error){
	slog.Info("started to save an idea...")

	id,err:=s.boards.CreateBoard(ctx,req.Name)
	if err!=nil{
		slog.Error(err.Error())
		return nil, status.Error(codes.Internal, "Internal error creating board")
	}
	resp := &boardsv1.CreateBoardResponse{Id: id}
	return resp,nil
}
func (s *serverAPI) GetBoard(ctx context.Context,req *boardsv1.GetBoardRequest) ( *boardsv1.GetBoardResponse, error){
	slog.Info("started to get a board")
	
	board, err := s.boards.GetBoard(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error getting board")
	}
	resp := &boardsv1.GetBoardResponse{Id: board.ID, Name: board.Name, IdeasIds: board.IdeasIds}
	return resp, nil
}

func (s *serverAPI) GetAllBoards(ctx context.Context, e *emptypb.Empty) (*boardsv1.GetAllBoardsResponse, error){
	
	slog.Info("started to get all ideas")
	boards, err := s.boards.GetAllBoards(ctx,e)
	
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error getting all ideas")
	}
	return &boardsv1.GetAllBoardsResponse{Boards: boards},nil
}