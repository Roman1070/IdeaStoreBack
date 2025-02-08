package ideas

import (
	"context"
	"fmt"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Ideas interface {
	CreateIdea(
		ctx context.Context,
		idea models.Idea) (id int64, err error)

	GetIdea(
		ctx context.Context,
		id int64,
	) (idea models.Idea, err error)
	DeleteIdea(ctx context.Context, id int64) (emptypb.Empty, error)
	GetAllIdeas(ctx context.Context, userId int64, limit, offset int32) ([]*models.Idea, error)
	GetIdeas(ctx context.Context, ids []int64) ([]*models.Idea, error)
	ChangeLikesCount(ctx context.Context, ideaId int64, increase bool) (int64, error)
	GetIdeasFromSearch(ctx context.Context, userId int64, input string) ([]*models.Idea, error)
}
type serverAPI struct {
	ideasv1.UnimplementedIdeasServer
	ideas Ideas
}

func Register(gRPC *grpc.Server, ideas Ideas) {
	ideasv1.RegisterIdeasServer(gRPC, &serverAPI{ideas: ideas})
}
func (s *serverAPI) CreateIdea(ctx context.Context, req *ideasv1.CreateRequest) (*ideasv1.CreateResponse, error) {
	slog.Info("started to save an idea...")

	id, err := s.ideas.CreateIdea(ctx, models.Idea{
		Image:       req.Image,
		Name:        req.Name,
		Description: req.Description,
		Link:        req.Link,
		Tags:        req.Tags,
		UserID:      req.UserId})
	if err != nil {
		slog.Error(err.Error())
		return nil, status.Error(codes.Internal, "Internal error creating idea")
	}

	resp := &ideasv1.CreateResponse{IdeaId: id}
	return resp, nil
}
func (s *serverAPI) ChangeLikesCount(ctx context.Context, req *ideasv1.ChangeLikesCountRequest) (*ideasv1.ChangeLikesCountResponse, error) {
	slog.Info("ideas grpc started to ChangeLikesCount")

	likesCount, err := s.ideas.ChangeLikesCount(ctx, req.IdeaId, req.Increase)
	if err != nil {
		slog.Error("ideas grpc error: " + err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ideasv1.ChangeLikesCountResponse{
		LikesCount: likesCount,
	}, nil
}
func (s *serverAPI) GetIdea(ctx context.Context, req *ideasv1.GetRequest) (*ideasv1.GetResponse, error) {
	slog.Info("started to get idea")

	idea, err := s.ideas.GetIdea(ctx, req.IdeaId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error getting idea")
	}
	resp := &ideasv1.GetResponse{
		Name:        idea.Name,
		Image:       idea.Image,
		Description: idea.Description,
		Link:        idea.Link,
		Tags:        idea.Tags,
		Likes:       idea.Likes,
		UserId:      idea.UserID}
	return resp, nil
}
func (s *serverAPI) DeleteIdea(ctx context.Context, req *ideasv1.DeleteRequest) (*emptypb.Empty, error) {
	slog.Info("started to delete an idea")

	_, err := s.ideas.DeleteIdea(ctx, req.IdeaId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error deleting idea")
	}

	return nil, nil
}
func (s *serverAPI) GetAllIdeas(ctx context.Context, req *ideasv1.GetAllRequest) (*ideasv1.GetAllResponse, error) {
	slog.Info("started to get all ideas")

	ideas, err := s.ideas.GetAllIdeas(ctx, req.UserId, req.Limit, req.Offset)
	if err != nil {
		slog.Error("client GetAllIdeas error: " + err.Error())
		return nil, fmt.Errorf("client GetAllIdeas error: " + err.Error())
	}

	var result []*ideasv1.IdeaData
	for _, idea := range ideas {
		result = append(result, &ideasv1.IdeaData{
			Id:          idea.ID,
			Name:        idea.Name,
			Description: idea.Description,
			Link:        idea.Link,
			Tags:        idea.Tags,
			Image:       idea.Image,
			Saved:       idea.Saved,
		})
	}
	return &ideasv1.GetAllResponse{Ideas: result}, nil
}

func (s *serverAPI) GetIdeas(ctx context.Context, req *ideasv1.GetIdeasRequest) (*ideasv1.GetIdeasResponse, error) {
	slog.Info("started to get ideas")

	ideas, err := s.ideas.GetIdeas(ctx, req.Ids)
	if err != nil {
		slog.Error("grpc error GetIdeas: " + err.Error())
		return nil, fmt.Errorf("grpc error GetIdeas: %v", err.Error())
	}

	var result []*ideasv1.IdeaData
	for _, idea := range ideas {
		result = append(result, &ideasv1.IdeaData{
			Id:          idea.ID,
			Name:        idea.Name,
			Description: idea.Description,
			Link:        idea.Link,
			Tags:        idea.Tags,
			Image:       idea.Image,
			Saved:       idea.Saved,
		})
	}

	return &ideasv1.GetIdeasResponse{Ideas: result}, nil
}

func (s *serverAPI) GetIdeasFromSearch(ctx context.Context, req *ideasv1.GetIdeasFromSearchRequest) (*ideasv1.GetIdeasFromSearchResponse, error) {
	slog.Info("grpc started to GetIdeasFromSearch")

	resp, err := s.ideas.GetIdeasFromSearch(ctx, req.UserId, req.Input)
	if err != nil {
		slog.Error("grpc error GetIdeasFromSearch: " + err.Error())
		return nil, fmt.Errorf("grpc error GetIdeasFromSearch: %v", err.Error())
	}
	var result []*ideasv1.IdeaData
	for _, idea := range resp {
		result = append(result, &ideasv1.IdeaData{
			Id:          idea.ID,
			Name:        idea.Name,
			Description: idea.Description,
			Link:        idea.Link,
			Tags:        idea.Tags,
			Image:       idea.Image,
			Saved:       idea.Saved,
		})
	}
	return &ideasv1.GetIdeasFromSearchResponse{
		Ideas: result,
	}, nil
}
