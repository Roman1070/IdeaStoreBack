package profiles

import (
	"context"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/grpc/profiles"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Profiles struct {
	Log *slog.Logger
	Api profiles.Profiles
}

func New(log *slog.Logger, api profiles.Profiles) *Profiles {
	return &Profiles{
		Log: log,
		Api: api,
	}
}

func (p *Profiles) CreateProfile(ctx context.Context, id int64, name, email string) (*emptypb.Empty, error) {
	slog.Info("service start CreateProfile")

	_, err := p.Api.CreateProfile(ctx, id, name, email)

	if err != nil {
		slog.Error("service CreateProfile error: " + err.Error())
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (p *Profiles) GetProfile(ctx context.Context, id int64) (models.Profile, error) {
	slog.Info("service start GetProfile")

	resp, err := p.Api.GetProfile(ctx, id)

	if err != nil {
		slog.Error("service GetProfile error: " + err.Error())
		return models.Profile{}, err
	}

	return models.Profile{
		ID:          resp.ID,
		Name:        resp.Name,
		AvatarImage: resp.AvatarImage,
		Description: resp.Description,
		Link:        resp.Link,
		Email:       resp.Email,
		Boards:      resp.Boards,
		SavedIdeas:  resp.SavedIdeas,
	}, nil
}
func (p *Profiles) GetProfileLight(ctx context.Context, id int64) (models.ProfileLight, error) {
	slog.Info("service started GetProfileLight")

	resp, err := p.Api.GetProfileLight(ctx, id)

	if err != nil {
		slog.Error("service GetProfile error: " + err.Error())
		return models.ProfileLight{}, err
	}
	return models.ProfileLight{
		ID:          resp.ID,
		Name:        resp.Name,
		AvatarImage: resp.AvatarImage,
	}, nil
}
func (p *Profiles) UpdateProfile(ctx context.Context, userId int64, name, avatarImage, description, link string) (*emptypb.Empty, error) {
	slog.Info("service start UpdateProfile")
	_, err := p.Api.UpdateProfile(ctx, userId, name, avatarImage, description, link)
	if err != nil {
		slog.Error("service UpdateProfile error: " + err.Error())
		return nil, err
	}
	return nil, nil
}

func (p *Profiles) ToggleSaveIdea(ctx context.Context, userId, ideaId, boardId int64) (bool, error) {
	slog.Info("service start ToggleSaveIdea")

	resp, err := p.Api.ToggleSaveIdea(ctx, userId, ideaId, boardId)

	if err != nil {
		slog.Error("service ToggleSaveIdea error: " + err.Error())
		return false, err
	}
	return resp, nil
}

func (p *Profiles) IsIdeaSaved(ctx context.Context, userId, ideaId int64) (bool, int64, error) {
	slog.Info("service start IsIdeaSaved")

	saved, boardId, err := p.Api.IsIdeaSaved(ctx, userId, ideaId)

	if err != nil {
		slog.Error("service IsIdeaSaved error: " + err.Error())
		return false, -1, err
	}
	return saved, boardId, nil
}

func (p *Profiles) GetSavedIdeas(ctx context.Context, userId int64) ([]*profilesv1.IdeaData, error) {
	slog.Info("service start GetSavedIdeas")

	resp, err := p.Api.GetSavedIdeas(ctx, userId)

	if err != nil {
		slog.Error("service GetSavedIdeas error: " + err.Error())
		return nil, err
	}
	return resp, nil
}

func (p *Profiles) GetSavedIdeasIds(ctx context.Context, userId int64) ([]int64, error) {
	slog.Info("service start GetSavedIdeasIds")

	resp, err := p.Api.GetSavedIdeasIds(ctx, userId)

	if err != nil {
		slog.Error("service GetSavedIdeas error: " + err.Error())
		return nil, err
	}
	return resp, nil
}

func (p *Profiles) MoveIdeasToBoard(ctx context.Context, userId, oldBoardId, newBoardId int64) (*emptypb.Empty, error) {
	slog.Info("service start MoveIdeasToBoard")

	resp, err := p.Api.MoveIdeasToBoard(ctx, userId, oldBoardId, newBoardId)

	if err != nil {
		slog.Error("service MoveIdeasToBoard error: " + err.Error())
		return nil, err
	}
	return resp, nil
}

func (p *Profiles) AddBoardToProfile(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {
	slog.Info("service start AddBoardToProfile")

	resp, err := p.Api.AddBoardToProfile(ctx, userId, boardId)

	if err != nil {
		slog.Error("service AddBoardToProfile error: " + err.Error())
		return nil, err
	}
	return resp, nil
}

func (p *Profiles) RemoveBoardFromProfile(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {
	slog.Info("service start RemoveBoardFromProfile")

	resp, err := p.Api.RemoveBoardFromProfile(ctx, userId, boardId)

	if err != nil {
		slog.Error("service RemoveBoardFromProfile error: " + err.Error())
		return nil, err
	}
	return resp, nil
}
