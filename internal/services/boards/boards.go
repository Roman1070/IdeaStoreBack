package boards

import (
	"context"
	"fmt"
	boardsv1 "idea-store-auth/gen/go/boards"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/grpc/boards"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Boards struct {
	log *slog.Logger
	Api boards.Boards
}

// New returns a new instance of the Auth service.
func New(log *slog.Logger, boardsApi boards.Boards) *Boards {
	return &Boards{
		log: log,
		Api: boardsApi,
	}
}
func (b *Boards) CreateBoard(ctx context.Context, name string, userId int64) (int64, error) {
	slog.Info("service started to CreateBoard")

	id, err := b.Api.CreateBoard(ctx, name, userId)
	if err != nil {
		slog.Error("service CreateBoard error: " + err.Error())
		return -1, fmt.Errorf("service CreateBoard error: " + err.Error())
	}

	return id, nil
}

func (b *Boards) GetBoard(ctx context.Context, id int64) (models.Board, error) {
	slog.Info("service started to GetBoard")

	board, err := b.Api.GetBoard(ctx, id)
	if err != nil {
		slog.Error("service GetBoard error: " + err.Error())
		return models.Board{}, fmt.Errorf("service GetBoard error: " + err.Error())
	}

	return board, nil
}

func (b *Boards) GetCurrentUsersBoards(ctx context.Context, userId int64) ([]*boardsv1.BoardData, error) {
	slog.Info("service started to GetCurrentUsersBoards")

	board, err := b.Api.GetCurrentUsersBoards(ctx, userId)
	if err != nil {
		slog.Error("service GetCurrentUsersBoards error: " + err.Error())
		return nil, fmt.Errorf("service GetCurrentUsersBoards error: " + err.Error())
	}

	return board, nil
}

func (b *Boards) SetIdeaSaved(ctx context.Context, boardId, ideaId int64, saved bool) (*emptypb.Empty, error) {
	slog.Info("service started SetIdeaSaved")
	_, err := b.Api.SetIdeaSaved(ctx, boardId, ideaId, saved)
	if err != nil {
		slog.Error("service SetIdeaSaved error: " + err.Error())
		return nil, fmt.Errorf("service SetIdeaSaved error: " + err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (b *Boards) GetIdeasInBoard(ctx context.Context, boardId int64) ([]*boardsv1.IdeaData, error) {
	slog.Info("service started GetIdeasInBoard")

	ideas, err := b.Api.GetIdeasInBoard(ctx, boardId)
	if err != nil {
		slog.Error("service GetIdeasInBoard error: " + err.Error())
		return nil, fmt.Errorf("service GetIdeasInBoard error: " + err.Error())
	}

	return ideas, nil
}
func (b *Boards) DeleteBoard(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {
	slog.Info("service started DeleteBoard")

	_, err := b.Api.DeleteBoard(ctx, userId, boardId)
	if err != nil {
		slog.Error("service DeleteBoard error: " + err.Error())
		return nil, fmt.Errorf("service DeleteBoard error: " + err.Error())
	}
	return nil, nil
}
