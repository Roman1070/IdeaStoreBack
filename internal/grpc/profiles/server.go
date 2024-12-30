package profiles

import (
	"context"
	"fmt"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Profiles interface {
	CreateProfile(ctx context.Context, id int64, name,email string) (*emptypb.Empty,error)
	GetProfile(ctx context.Context, id int64) (models.Profile, error)
	ToggleSaveIdea(ctx context.Context, userId,ideaId,boardId int64) (bool,error)
	IsIdeaSaved(ctx context.Context, userId,ideaId int64) (bool,error)
}

type serverAPI struct{
	profilesv1.UnimplementedProfilesServer
	profiles Profiles
}
func Register(gRPC *grpc.Server, profiles Profiles) {
	profilesv1.RegisterProfilesServer(gRPC, &serverAPI{profiles: profiles})
}

func (s *serverAPI) CreateProfile(ctx context.Context, req *profilesv1.CreateProfileRequest) (*emptypb.Empty,error){
	slog.Info("grpc start CreateProfile")

	_,err:= s.profiles.CreateProfile(ctx,req.Id,req.Name,req.Email)
	if err!=nil{
		slog.Error(err.Error())
		return &emptypb.Empty{}, fmt.Errorf("grpc create profile error: "+err.Error())
	}
	return &emptypb.Empty{},nil
}

func(s *serverAPI) GetProfile(ctx context.Context, req *profilesv1.GetProfileRequest) (*profilesv1.GetProfileResponse, error){
	slog.Info("grpc start GetProfile")

	resp, err:= s.profiles.GetProfile(ctx,req.Id)
	if err!=nil{
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc get profile error: "+err.Error())
	}
	return &profilesv1.GetProfileResponse{
		Data: &profilesv1.ProfileData{
			Id: resp.ID,
			Email: resp.Email,
			AvatarImage: resp.AvatarImage,
			Name: resp.Name,
			Description: resp.Description,
			Link: resp.Link,
			Boards: resp.Boards,
			SavedIdeas: resp.SavedIdeas,
		},
	}, nil
}

func(s *serverAPI) ToggleSaveIdea(ctx context.Context, req *profilesv1.ToggleSaveRequest) (*profilesv1.ToggleSaveResponse,error){
	slog.Info("grpc start ToggleSaveIdea")

	resp, err:= s.profiles.ToggleSaveIdea(ctx,req.UserId,req.IdeaId,req.BoardId)
	if err!=nil{
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc error toggle save idea: "+err.Error())
	}
	return &profilesv1.ToggleSaveResponse{
		NowSaved: resp,
	},nil
}

func(s *serverAPI) IsIdeaSaved(ctx context.Context, req *profilesv1.IsIdeaSavedRequest) (*profilesv1.IsIdeaSavedResponse, error){
	slog.Info("grpc start IsIdeaSaved")

	resp, err:= s.profiles.IsIdeaSaved(ctx,req.UserId,req.IdeaId)
	if err!=nil{
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc error IsIdeaSaved: "+err.Error())
	}
	return &profilesv1.IsIdeaSavedResponse{
		Saved: resp,
	},nil
}