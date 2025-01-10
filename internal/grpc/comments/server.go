package comments

import (
	"context"
	"fmt"
	commentsv1 "idea-store-auth/gen/go/comments"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Comments interface {
	CreateComment(ctx context.Context, ideaId, userId int64, text, creationDate string) (*emptypb.Empty, error)
	GetComments(ctx context.Context, ideaId int64) ([]*models.Comment, error)
}

type serverAPI struct {
	commentsv1.UnimplementedCommentsServer
	comments Comments
}

func Register(gRPC *grpc.Server, comments Comments) {
	commentsv1.RegisterCommentsServer(gRPC, &serverAPI{comments: comments})
}
func (s *serverAPI) CreateComment(ctx context.Context, req *commentsv1.CreateCommentRequest) (*emptypb.Empty, error) {
	slog.Info("grpc started CreateComment")

	_, err := s.comments.CreateComment(ctx, req.IdeaId, req.UserId, req.Text, req.CreationDate)
	if err != nil {
		return nil, fmt.Errorf("error grpc CreateComment: %v", err.Error())
	}
	return nil, nil
}
func (s *serverAPI) GetComments(ctx context.Context, req *commentsv1.GetCommentsRequest) (*commentsv1.GetCommentsResponse, error) {

	slog.Info("grpc started GetComments")

	resp, err := s.comments.GetComments(ctx, req.IdeaId)
	if err != nil {
		return nil, fmt.Errorf("error grpc GetComments: %v", err.Error())
	}

	var result []*commentsv1.CommentData
	for _, c := range resp {
		result = append(result, &commentsv1.CommentData{
			Id:           c.ID,
			UserId:       c.UserId,
			Text:         c.Text,
			CreationDate: c.CreationDate,
			Username:     c.Username,
			Avatar:       c.Avatar,
		})
	}
	return &commentsv1.GetCommentsResponse{
		Comments: result,
	}, nil
}
