package ideas

import (
	"context"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IdeasAPI interface {
	Create(
		ctx context.Context,
		idea models.Idea) (id int64, err error)

	Get(
		ctx context.Context,
		id int64,
	) (idea models.Idea, err error)
	Delete(ctx context.Context, id int64) (emptypb.Empty, error)
}
type serverAPI struct {
	ideasv1.UnimplementedIdeasServer
	ideas IdeasAPI
}
func Register(gRPC *grpc.Server, ideas IdeasAPI) {
	ideasv1.RegisterIdeasServer(gRPC, &serverAPI{ideas: ideas})
}
func (s *serverAPI) Create(ctx context.Context,req *ideasv1.CreateRequest) (*ideasv1.CreateResponse, error){
	slog.Info("started to save an idea...")

	id,err:=s.ideas.Create(ctx,models.Idea{Image: req.Image, Name: req.Name,Description: req.Description, Link: req.Link, Tags: req.Tags})
	if err!=nil{
		slog.Error(err.Error())
		return nil, status.Error(codes.Internal, "Internal error")
	}
	resp := &ideasv1.CreateResponse{IdeaId: id}
	return resp,nil
}
func (s *serverAPI) Get(ctx context.Context,req *ideasv1.GetRequest) ( *ideasv1.GetResponse, error){
	slog.Info("started to get idea")
	
	idea, err := s.ideas.Get(ctx, req.IdeaId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	resp := &ideasv1.GetResponse{Name: idea.Name, Image: idea.Image, Description: idea.Description, Link: idea.Link, Tags: idea.Tags}
	return resp, nil
}
func (s *serverAPI) Delete(ctx context.Context,req *ideasv1.DeleteRequest) (*emptypb.Empty, error){
	return nil,nil
}