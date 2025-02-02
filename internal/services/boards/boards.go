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
	const op = "service.boards.Create"

	log := b.log.With(
		slog.String("op", op),
		slog.String("board name", name),
	)
	log.Info("Creating a board...")

	id, err := b.Api.CreateBoard(ctx, name, userId)
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, "Internal error creating board")
	}
	return id, nil
}

func (b *Boards) GetBoard(ctx context.Context, id int64) (models.Board, error) {
	const op = "service.boards.Get"

	log := b.log.With(
		slog.String("op", op),
		slog.Int64("board id", id),
	)
	log.Info("Getting a board...")

	board, err := b.Api.GetBoard(ctx, id)
	if err != nil {
		return models.Board{}, fmt.Errorf("%s: %v", op, "Internal error getting board")
	}
	return board, nil
}

func (b *Boards) GetCurrentUsersBoards(ctx context.Context, userId int64) ([]*boardsv1.BoardData, error) {

	const op = "service.boards.GetAll"

	slog.Info("Getting all boards...")

	board, err := b.Api.GetCurrentUsersBoards(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, "Internal error getting board")
	}
	return board, nil
}

func (b *Boards) SetIdeaSaved(ctx context.Context, boardId, ideaId int64, saved bool) (*emptypb.Empty, error) {

	const op = "service.boards.GetAll"

	slog.Info("started SetIdeaSaved...")

	_, err := b.Api.SetIdeaSaved(ctx, boardId, ideaId, saved)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, "Internal error SetIdeaSaved")
	}
	return &emptypb.Empty{}, nil
}

func (b *Boards) GetIdeasInBoard(ctx context.Context, boardId int64) ([]*boardsv1.IdeaData, error) {

	const op = "service.boards.GetIdeasInBoard"

	slog.Info("started GetIdeasInBoard...")

	ideas, err := b.Api.GetIdeasInBoard(ctx, boardId)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, "Internal error SetIdeaSaved")
	}
	return ideas, nil
}
func (b *Boards) DeleteBoard(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {
	const op = "service.boards.DeleteBoard"

	slog.Info("started DeleteBoard...")

	_, err := b.Api.DeleteBoard(ctx, userId, boardId)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, "Internal error DeleteBoard")
	}
	return nil, nil
}
