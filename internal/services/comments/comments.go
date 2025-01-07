package comments

import (
	"context"
	"fmt"
	commentsv1 "idea-store-auth/gen/go/comments"
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

func(c *Comments) CreateComment(ctx context.Context, ideaId, userId int64, text, creationDate string) (*emptypb.Empty, error){
	slog.Info("server started CreateComment")
	_, err:= c.Api.CreateComment(ctx,ideaId,userId,text,creationDate)
	if err!=nil{
		return nil, fmt.Errorf("server error CreateComment: %v",err.Error())
	}
	return &emptypb.Empty{}, nil
}
func(c *Comments) GetComments(ctx context.Context, ideaId int64)([]*commentsv1.CommentData, error){
	
	slog.Info("server started GetComments")
	resp, err:= c.Api.GetComments(ctx,ideaId)
	if err!=nil{
		return nil, fmt.Errorf("server error CreateComment: %v",err.Error())
	}
	var result []*commentsv1.CommentData //create DTO files with toGRPC method
	for _, c := range resp{
		result = append(result, &commentsv1.CommentData{
			Id: c.ID,
			UserId: c.UserId,
			Text: c.Text,
			CreationDate: c.CreationDate,
		})
	}
	return result, nil
}