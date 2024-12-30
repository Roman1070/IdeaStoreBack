package profiles

import (
	"context"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/grpc/profiles"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Profiles struct {
	Log *slog.Logger
	Api profiles.Profiles
}

func New(log *slog.Logger, api profiles.Profiles) *Profiles{
	return &Profiles{
		Log:log,
		Api:api,
	}
}

func (p *Profiles) CreateProfile(ctx context.Context, id int64, name,email string) (*emptypb.Empty,error){
	slog.Info("service start CreateProfile")

	_,err:= p.Api.CreateProfile(ctx,id,name,email)

	if err!=nil{
		slog.Error("service CreateProfile error: "+err.Error())
		return nil,err
	}
	return &emptypb.Empty{},nil
}
func(p *Profiles) GetProfile(ctx context.Context, id int64) (models.Profile, error){
	slog.Info("service start GetProfile")

	resp, err:= p.Api.GetProfile(ctx,id)

	if err!=nil{
		slog.Error("service GetProfile error: "+err.Error())
		return models.Profile{},err
	}

	return models.Profile{
		ID: resp.ID,
		Name: resp.Name,
		AvatarImage: resp.AvatarImage,
		Description: resp.Description,
		Link: resp.Link,
		Email: resp.Email,
		Boards: resp.Boards,
		SavedIdeas: resp.SavedIdeas,
	},nil
}

func (p *Profiles) ToggleSaveIdea(ctx context.Context, userId,ideaId,boardId int64) (bool, error){
	slog.Info("service start ToggleSaveIdea")

	resp, err:= p.Api.ToggleSaveIdea(ctx,userId,ideaId,boardId)

	if err!=nil{
		slog.Error("service ToggleSaveIdea error: "+err.Error())
		return false,err
	}
	return resp,nil
}