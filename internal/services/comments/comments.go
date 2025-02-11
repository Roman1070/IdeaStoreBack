package comments

import (
	"context"
	"fmt"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/grpc/comments"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Comments struct {
	log *slog.Logger
	Api comments.Comments
}

// New returns a new instance of the Comments service.
func New(log *slog.Logger, commentsApi comments.Comments) *Comments {
	return &Comments{
		log: log,
		Api: commentsApi,
	}
}

func (c *Comments) CreateComment(ctx context.Context, ideaId, userId int64, text, creationDate string) (*emptypb.Empty, error) {
	slog.Info("service started CreateComment")

	_, err := c.Api.CreateComment(ctx, ideaId, userId, text, creationDate)
	if err != nil {
		c.log.Error("service error CreateComment: " + err.Error())
		return nil, fmt.Errorf("serivce error CreateComment: %v", err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (c *Comments) GetComments(ctx context.Context, ideaId int64) ([]*models.Comment, error) {
	slog.Info("server started GetComments")

	resp, err := c.Api.GetComments(ctx, ideaId)
	if err != nil {
		c.log.Error("service error GetComments: " + err.Error())
		return nil, fmt.Errorf("serivce error GetComments: %v", err.Error())
	}

	return resp, nil
}
