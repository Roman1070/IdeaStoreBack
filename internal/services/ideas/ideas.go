package ideas

import (
	"context"
	"fmt"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/grpc/ideas"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Ideas struct {
	log *slog.Logger
	Api ideas.IdeasAPI
}


// New returns a new instance of the Auth service.
func New(log *slog.Logger, ideasApi ideas.IdeasAPI) *Ideas {
	return &Ideas{
		log:         log,
		Api:ideasApi,
	}
}
func (i *Ideas) CreateIdea(ctx context.Context, idea models.Idea) (int64, error){
	const op = "service.ideas.Create"

	log := i.log.With(
		slog.String("op", op),
		slog.String("idea name", idea.Name),
	)
	log.Info("Creating an idea...")
	
	id, err := i.Api.CreateIdea(ctx, idea)
	if err!=nil{
		return -1, fmt.Errorf("%s: %v", op, "Internal error")
	}
	return id, nil
}
func (i *Ideas) GetIdea(ctx context.Context, id int64) (models.Idea, error){
	const op = "service.ideas.Create"

	log := i.log.With(
		slog.String("op", op),
		slog.Int64("idea id", id),
	)
	log.Info("Getting an idea...")

	idea, err := i.Api.GetIdea(ctx,id)
	
	if err!=nil{
		return models.Idea{}, fmt.Errorf("%s: %v", op, "Internal error")
	}
	return idea, nil
}

func (i *Ideas) DeleteIdea(ctx context.Context, id int64) (emptypb.Empty, error){
	const op = "service.ideas.Delete"

	log := i.log.With(
		slog.String("op", op),
		slog.Int64("idea id", id),
	)
	log.Info("Deleting an idea...")

	_,err:= i.Api.DeleteIdea(ctx,id)
	if err!=nil{
		return emptypb.Empty{}, fmt.Errorf("%s: %v", op, "Internal error")
	}
	return emptypb.Empty{}, nil
}
