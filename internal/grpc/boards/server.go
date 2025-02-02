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
	CreateBoard(ctx context.Context, name string, userId int64) (int64, error)
	GetBoard(ctx context.Context, id int64) (models.Board, error)
	GetCurrentUsersBoards(ctx context.Context, userId int64) ([]*boardsv1.BoardData, error)
	SetIdeaSaved(ctx context.Context, boardId, ideaId int64, saved bool) (*emptypb.Empty, error)
	GetIdeasInBoard(ctx context.Context, boardId int64) ([]*boardsv1.IdeaData, error)
	DeleteBoard(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error)
}

type serverAPI struct {
	boardsv1.UnimplementedBoardsServer
	boards Boards
}

func Register(gRPC *grpc.Server, boards Boards) {
	boardsv1.RegisterBoardsServer(gRPC, &serverAPI{boards: boards})
}

func (s *serverAPI) CreateBoard(ctx context.Context, req *boardsv1.CreateBoardRequest) (*boardsv1.CreateBoardResponse, error) {
	slog.Info("started to save an idea...")

	id, err := s.boards.CreateBoard(ctx, req.Name, req.UserId)
	if err != nil {
		slog.Error(err.Error())
		return nil, status.Error(codes.Internal, "Internal error creating board")
	}
	resp := &boardsv1.CreateBoardResponse{Id: id}
	return resp, nil
}
func (s *serverAPI) GetBoard(ctx context.Context, req *boardsv1.GetBoardRequest) (*boardsv1.GetBoardResponse, error) {
	slog.Info("started to get a board")

	board, err := s.boards.GetBoard(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error getting board")
	}
	resp := &boardsv1.GetBoardResponse{Id: board.ID, Name: board.Name, IdeasIds: board.IdeasIds}
	return resp, nil
}

func (s *serverAPI) GetAllBoards(ctx context.Context, req *boardsv1.GetAllBoardsRequest) (*boardsv1.GetAllBoardsResponse, error) {

	slog.Info("started to get all ideas")
	boards, err := s.boards.GetCurrentUsersBoards(ctx, req.UserId)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error getting all ideas")
	}
	return &boardsv1.GetAllBoardsResponse{Boards: boards}, nil
}

func (s *serverAPI) SetIdeaSaved(ctx context.Context, req *boardsv1.SetIdeaSavedRequest) (*emptypb.Empty, error) {

	slog.Info("started to SetIdeaSaved grpc")
	_, err := s.boards.SetIdeaSaved(ctx, req.BoardId, req.IdeaId, req.Saved)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error SetIdeaSaved")
	}
	return &emptypb.Empty{}, nil
}

func (s *serverAPI) GetIdeasInBoard(ctx context.Context, req *boardsv1.GetIdeasInBoardRequest) (*boardsv1.GetIdeasInBoardResponse, error) {

	slog.Info("started to GetIdeasInBoard grpc")
	ideas, err := s.boards.GetIdeasInBoard(ctx, req.BoardId)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error SetIdeaSaved")
	}
	return &boardsv1.GetIdeasInBoardResponse{
		Ideas: ideas,
	}, nil
}
func (s *serverAPI) DeleteBoard(ctx context.Context, req *boardsv1.DeleteBoardRequest) (*emptypb.Empty, error) {
	slog.Info("started to DeleteBoard grpc")
	_, err := s.boards.DeleteBoard(ctx, req.UserId, req.BoardId)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error DeleteBoard")
	}
	return &emptypb.Empty{}, nil
}
