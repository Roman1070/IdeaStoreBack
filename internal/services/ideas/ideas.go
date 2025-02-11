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
	slog.Info("service started to CreateIdea")

	id, err := i.Api.CreateIdea(ctx, idea)

	if err != nil {
		slog.Error("service CreateIdea error: " + err.Error())
		return -1, fmt.Errorf("service CreateIdea error: " + err.Error())
	}

	return id, nil
}
func (i *Ideas) GetIdea(ctx context.Context, id int64) (models.Idea, error) {
	slog.Info("service started to GetIdea")

	idea, err := i.Api.GetIdea(ctx, id)

	if err != nil {
		slog.Error("service GetIdea error: " + err.Error())
		return models.Idea{}, fmt.Errorf("service GetIdea error: " + err.Error())
	}

	return idea, nil
}
func (i *Ideas) ChangeLikesCount(ctx context.Context, ideaId int64, increase bool) (int64, error) {
	slog.Info("service started to ChangeLikesCount")

	likesCount, err := i.Api.ChangeLikesCount(ctx, ideaId, increase)
	if err != nil {
		slog.Error("service GetIdeas error: " + err.Error())
		return -1, fmt.Errorf("service GetIdeas error: " + err.Error())
	}

	return likesCount, nil
}
func (i *Ideas) DeleteIdea(ctx context.Context, id int64) (*emptypb.Empty, error) {
	slog.Info("service started to DeleteIdea")
	_, err := i.Api.DeleteIdea(ctx, id)
	if err != nil {
		slog.Error("service GetIdeas error: " + err.Error())
		return nil, fmt.Errorf("service GetIdeas error: " + err.Error())
	}

	return nil, nil
}

func (i *Ideas) GetAllIdeas(ctx context.Context, userId int64, limit, offset int32) ([]*models.Idea, error) {
	slog.Info("service started to GetAllIdeas")

	ideas, err := i.Api.GetAllIdeas(ctx, userId, limit, offset)

	if err != nil {
		slog.Error("service GetIdeas error: " + err.Error())
		return nil, fmt.Errorf("service GetIdeas error: " + err.Error())
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
