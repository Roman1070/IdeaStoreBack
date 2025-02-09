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
	Api ideas.Ideas
}

// New returns a new instance of the Ideas service.
func New(log *slog.Logger, ideasApi ideas.Ideas) *Ideas {
	return &Ideas{
		log: log,
		Api: ideasApi,
	}
}
func (i *Ideas) CreateIdea(ctx context.Context, idea models.Idea) (int64, error) {
	const op = "service.ideas.Create"

	log := i.log.With(
		slog.String("op", op),
		slog.String("idea name", idea.Name),
	)
	log.Info("Creating an idea...")

	id, err := i.Api.CreateIdea(ctx, idea)
	if err != nil {
		return -1, fmt.Errorf("%s: %v", op, "Internal error")
	}
	return id, nil
}
func (i *Ideas) GetIdea(ctx context.Context, id int64) (models.Idea, error) {
	const op = "service.ideas.Create"

	log := i.log.With(
		slog.String("op", op),
		slog.Int64("idea id", id),
	)
	log.Info("Getting an idea...")

	idea, err := i.Api.GetIdea(ctx, id)

	if err != nil {
		return models.Idea{}, fmt.Errorf("%s: %v", op, "Internal error")
	}
	return idea, nil
}
func (i *Ideas) ChangeLikesCount(ctx context.Context, ideaId int64, increase bool) (int64, error) {
	slog.Info("ideas service started to ChangeLikesCount")

	likesCount, err := i.Api.ChangeLikesCount(ctx, ideaId, increase)
	if err != nil {
		slog.Error("ideas service ChangeLikesCount error:" + err.Error())
		return -1, fmt.Errorf("ideas service ChangeLikesCount error: %v", err.Error())
	}
	return likesCount, nil
}
func (i *Ideas) DeleteIdea(ctx context.Context, id int64) (emptypb.Empty, error) {
	const op = "service.ideas.Delete"

	log := i.log.With(
		slog.String("op", op),
		slog.Int64("idea id", id),
	)
	log.Info("Deleting an idea...")

	_, err := i.Api.DeleteIdea(ctx, id)
	if err != nil {
		return emptypb.Empty{}, fmt.Errorf("%s: %v", op, "Internal error")
	}
	return emptypb.Empty{}, nil
}

func (i *Ideas) GetAllIdeas(ctx context.Context, userId int64, limit, offset int32) ([]*models.Idea, error) {

	const op = "service.ideas.GetAllIdeas"

	log := i.log.With(
		slog.String("op", op),
	)
	log.Info("Getting all ideas...")

	ideas, err := i.Api.GetAllIdeas(ctx, userId, limit, offset)

	if err != nil {
		log.Error(err.Error())
		return nil, fmt.Errorf("%s: %v", op, "Internal error")
	}

	return ideas, nil
}

func (i *Ideas) GetIdeas(ctx context.Context, ids []int64, limit, offset int32) ([]*models.Idea, error) {
	slog.Info("service started to GetIdeas")
	ideas, err := i.Api.GetIdeas(ctx, ids, limit, offset)
	if err != nil {
		slog.Error("service GetIdeas error: " + err.Error())
		return nil, fmt.Errorf("service GetIdeas error: " + err.Error())
	}

	return ideas, nil
}

func (i *Ideas) GetIdeasFromSearch(ctx context.Context, userId int64, input string) ([]*models.Idea, error) {
	slog.Info("service started GetIdeasFromSearch")

	resp, err := i.Api.GetIdeasFromSearch(ctx, userId, input)
	if err != nil {
		slog.Error("error GetIdeasFromSearch: " + err.Error())
		return nil, fmt.Errorf("error GetIdeasFromSearch: " + err.Error())
	}

	return resp, nil
}
