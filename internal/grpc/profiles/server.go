package profiles

import (
	"context"
	"fmt"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Profiles interface {
	CreateProfile(ctx context.Context, id int64, name, email string) (*emptypb.Empty, error)
	GetProfile(ctx context.Context, id int64) (models.Profile, error)
	GetProfileLight(ctx context.Context, id int64) (models.ProfileLight, error)
	UpdateProfile(ctx context.Context, userId int64, name, avatarImage, description, link string) (*emptypb.Empty, error)
	ToggleSaveIdea(ctx context.Context, userId, ideaId, boardId int64) (bool, error)
	ToggleLikeIdea(ctx context.Context, userId, ideaId int64) (bool, int64, error)
	IsIdeaSaved(ctx context.Context, userId, ideaId int64) (bool, int64, error)
	IsIdeaLiked(ctx context.Context, userId, ideaId int64) (bool, error)
	GetSavedIdeas(ctx context.Context, userId int64) ([]*profilesv1.IdeaData, error)
	GetSavedIdeasIds(ctx context.Context, userId int64) ([]int64, error)
	MoveIdeasToBoard(ctx context.Context, userId, oldBoardId, newBoardId int64) (*emptypb.Empty, error)
	AddBoardToProfile(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error)
	RemoveBoardFromProfile(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error)
	GetProfilesFromSearch(ctx context.Context, input string) ([]*models.ProfileLight, error)
}

type serverAPI struct {
	profilesv1.UnimplementedProfilesServer
	profiles Profiles
}

func Register(gRPC *grpc.Server, profiles Profiles) {
	profilesv1.RegisterProfilesServer(gRPC, &serverAPI{profiles: profiles})
}

func (s *serverAPI) CreateProfile(ctx context.Context, req *profilesv1.CreateProfileRequest) (*emptypb.Empty, error) {
	slog.Info("grpc start CreateProfile")

	_, err := s.profiles.CreateProfile(ctx, req.Id, req.Name, req.Email)
	if err != nil {
		slog.Error(err.Error())
		return &emptypb.Empty{}, fmt.Errorf("grpc create profile error: " + err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (s *serverAPI) IsIdeaLiked(ctx context.Context, req *profilesv1.IsIdeaLikedRequest) (*profilesv1.IsIdeaLikedResponse, error) {
	slog.Info("grpc started IsIdeaLiked")

	resp, err := s.profiles.IsIdeaLiked(ctx, req.UserId, req.IdeaId)
	if err != nil {
		slog.Error("grpc IsIdeaLiked error: " + err.Error())
		return nil, fmt.Errorf("grpc IsIdeaLiked error: %v", err.Error())
	}

	return &profilesv1.IsIdeaLikedResponse{
		Liked: resp,
	}, nil
}
func (s *serverAPI) GetProfile(ctx context.Context, req *profilesv1.GetProfileRequest) (*profilesv1.GetProfileResponse, error) {
	slog.Info("grpc start GetProfile")

	resp, err := s.profiles.GetProfile(ctx, req.Id)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc get profile error: " + err.Error())
	}
	return &profilesv1.GetProfileResponse{
		Data: &profilesv1.ProfileData{
			Id:          resp.ID,
			Email:       resp.Email,
			AvatarImage: resp.AvatarImage,
			Name:        resp.Name,
			Description: resp.Description,
			Link:        resp.Link,
			Boards:      resp.Boards,
			SavedIdeas:  resp.SavedIdeas,
		},
	}, nil
}
func (s *serverAPI) ToggleLikeIdea(ctx context.Context, req *profilesv1.ToggleLikeIdeaRequest) (*profilesv1.ToggleLikeIdeaResponse, error) {
	slog.Info("grpc started ToggleLikeIdea")

	nowLiked, likesCount, err := s.profiles.ToggleLikeIdea(ctx, req.UserId, req.IdeaId)

	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc ToggleLikeIdea error: " + err.Error())
	}
	return &profilesv1.ToggleLikeIdeaResponse{
		NowLiked:   nowLiked,
		LikesCount: likesCount,
	}, nil
}
func (s *serverAPI) GetProfilesFromSearch(ctx context.Context, req *profilesv1.GetProfilesFromSearchRequest) (*profilesv1.GetProfilesFromSearchResponse, error) {
	slog.Info("grpc start GetProfilesFromSearch")

	resp, err := s.profiles.GetProfilesFromSearch(ctx, req.Input)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc get profile error: " + err.Error())
	}
	var result []*profilesv1.ProfileDataLight
	for _, prof := range resp {
		result = append(result, &profilesv1.ProfileDataLight{
			Id:     prof.ID,
			Name:   prof.Name,
			Avatar: prof.AvatarImage,
		})
	}
	return &profilesv1.GetProfilesFromSearchResponse{
		Profiles: result,
	}, nil
}
func (s *serverAPI) GetProfileLight(ctx context.Context, req *profilesv1.GetProfileLightRequest) (*profilesv1.GetProfileLightResponse, error) {
	slog.Info("grpc start GetProfileLight")

	resp, err := s.profiles.GetProfileLight(ctx, req.UserId)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc GetProfileLight error: " + err.Error())
	}
	return &profilesv1.GetProfileLightResponse{
		Name:   resp.Name,
		Avatar: resp.AvatarImage,
	}, nil
}
func (s *serverAPI) UpdateProfile(ctx context.Context, req *profilesv1.UpdateProfileRequest) (*emptypb.Empty, error) {
	slog.Info("grpc start UpdateProfile")

	_, err := s.profiles.UpdateProfile(ctx, req.UserId, req.Name, req.Avatar, req.Description, req.Link)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc update profile error: " + err.Error())
	}
	return nil, nil
}
func (s *serverAPI) ToggleSaveIdea(ctx context.Context, req *profilesv1.ToggleSaveRequest) (*profilesv1.ToggleSaveResponse, error) {
	slog.Info("grpc start ToggleSaveIdea")

	resp, err := s.profiles.ToggleSaveIdea(ctx, req.UserId, req.IdeaId, req.BoardId)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc error toggle save idea: " + err.Error())
	}
	return &profilesv1.ToggleSaveResponse{
		NowSaved: resp,
	}, nil
}

func (s *serverAPI) IsIdeaSaved(ctx context.Context, req *profilesv1.IsIdeaSavedRequest) (*profilesv1.IsIdeaSavedResponse, error) {
	slog.Info("grpc start IsIdeaSaved")

	saved, boardId, err := s.profiles.IsIdeaSaved(ctx, req.UserId, req.IdeaId)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc error IsIdeaSaved: " + err.Error())
	}
	return &profilesv1.IsIdeaSavedResponse{
		Saved:   saved,
		BoardId: boardId,
	}, nil
}

func (s *serverAPI) GetSavedIdeas(ctx context.Context, req *profilesv1.GetSavedIdeasRequest) (*profilesv1.GetSavedIdeasResponse, error) {
	slog.Info("grpc start GetSavedIdeas")

	resp, err := s.profiles.GetSavedIdeas(ctx, req.UserId)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc error GetSavedIdeas: " + err.Error())
	}
	return &profilesv1.GetSavedIdeasResponse{
		Ideas: resp,
	}, nil
}

func (s *serverAPI) GetSavedIdeasIds(ctx context.Context, req *profilesv1.GetSavedIdeasIdsRequest) (*profilesv1.GetSavedIdeasIdsResponse, error) {
	slog.Info("grpc start GetSavedIdeasIds")

	resp, err := s.profiles.GetSavedIdeasIds(ctx, req.UserId)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("grpc error GetSavedIdeasIds: " + err.Error())
	}
	return &profilesv1.GetSavedIdeasIdsResponse{
		IdeasIds: resp,
	}, nil
}

func (s *serverAPI) MoveIdeasToBoard(ctx context.Context, req *profilesv1.MoveIdeaToBoardRequest) (*emptypb.Empty, error) {

	slog.Info("started to MoveIdeasToBoard grpc")
	_, err := s.profiles.MoveIdeasToBoard(ctx, req.UserId, req.OldBoardId, req.NewBoardId)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error MoveIdeasToBoard")
	}
	return nil, nil
}
func (s *serverAPI) AddBoardToProfile(ctx context.Context, req *profilesv1.AddBoardToProfileRequest) (*emptypb.Empty, error) {

	slog.Info("started to AddBoardToProfile grpc")
	_, err := s.profiles.AddBoardToProfile(ctx, req.UserId, req.BoardId)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error AddBoardToProfile")
	}
	return nil, nil
}

func (s *serverAPI) RemoveBoardFromProfile(ctx context.Context, req *profilesv1.RemoveBoardFromProfileRequest) (*emptypb.Empty, error) {

	slog.Info("started to RemoveBoardFromProfile grpc")
	_, err := s.profiles.RemoveBoardFromProfile(ctx, req.UserId, req.BoardId)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error RemoveBoardFromProfile")
	}
	return nil, nil
}
