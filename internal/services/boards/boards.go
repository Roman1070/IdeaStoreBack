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
func (b *Boards) CreateBoard(ctx context.Context, name string) (int64, error) {
	const op = "service.boards.Create"

	log := b.log.With(
		slog.String("op", op),
		slog.String("board name", name),
	)
	log.Info("Creating a board...")

	id, err := b.Api.CreateBoard(ctx, name)
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, "Internal error creating board")
	}
	return id, nil
}

func (b *Boards) GetBoard(ctx context.Context, id int64) (models.Board, error){
	const op = "service.boards.Get"

	log := b.log.With(
		slog.String("op", op),
		slog.Int64("board id", id),
	)
	log.Info("Getting a board...")

	board, err:= b.Api.GetBoard(ctx,id)
	if err!=nil{
		return models.Board{}, fmt.Errorf("%s: %v", op, "Internal error getting board")
	}
	return board,nil
}

func (b *Boards) GetAllBoards(ctx context.Context, e *emptypb.Empty) ([]*boardsv1.BoardData, error){

	const op = "service.boards.GetAll"
	
	slog.Info("Getting all boards...")

	board, err:= b.Api.GetAllBoards(ctx,e)
	if err!=nil{
		return nil, fmt.Errorf("%s: %v", op, "Internal error getting board")
	}
	return board,nil
}