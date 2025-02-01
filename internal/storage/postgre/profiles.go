package postgre

import (
	"context"
	"fmt"
	boardsv1 "idea-store-auth/gen/go/boards"
	ideasv1 "idea-store-auth/gen/go/idea"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateProfile(ctx context.Context, id int64, name, email string) (*emptypb.Empty, error) {
	slog.Info("storage start CreateProfile")

	const query = `
		INSERT INTO profiles(id,email,avatarImage,name,description,link,boards,savedIdeas)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8);
	`
	_, err := s.db.Exec(ctx, query)
	if err != nil {
		slog.Error("storage CreateProfile error: " + err.Error())
		return nil, fmt.Errorf("storage CreateProfile error: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Storage) GetProfile(ctx context.Context, id int64) (models.Profile, error) {
	slog.Info("storage start GetProfile")

	const query = `
		SELECT id,email,avatarImage,name,description,link,boards,savedIdeas 
		FROM profiles 
		WHERE id=$1;	
	`

	var profile models.Profile
	var boardsStr string
	var ideasStr string
	err := s.db.QueryRow(ctx, query, id).
		Scan(&profile.ID, &profile.Email, &profile.AvatarImage, &profile.Name, &profile.Description, &profile.Link, &boardsStr, &ideasStr)

	if err != nil {
		slog.Error("storage GetProfile error scaninng row: " + err.Error())
		return models.Profile{}, fmt.Errorf("storage GetProfile error scaninng row: %v", err.Error())
	}
	pairsSlice, err := ParseIdPairs(ideasStr)
	if err != nil {
		slog.Error("storage GetProfile error parsing ideas-board: " + err.Error())
		return models.Profile{}, fmt.Errorf("storage GetProfile error parsing ideas-board: %v", err.Error())
	}
	profile.SavedIdeas = toGRPCFormat(pairsSlice)
	profile.Boards, err = ParseIdsSqlite(boardsStr)

	if err != nil {
		slog.Error("storage GetProfile error parsing boards ids: " + err.Error())
		return models.Profile{}, fmt.Errorf("storage GetProfile error parsing boards ids: %v", err.Error())
	}

	return profile, nil
}
func (s *Storage) GetProfileLight(ctx context.Context, id int64) (models.ProfileLight, error) {
	slog.Info("storage start GetProfileLight")

	const query = `
		SELECT id,name,avatarImage 
		FROM profiles 
		WHERE id=$1;
	`
	var profile models.ProfileLight
	err := s.db.QueryRow(ctx, query, id).Scan(&profile.ID, &profile.Name, &profile.AvatarImage)

	if err != nil {
		slog.Error("storage GetProfileLight error: " + err.Error())
		return models.ProfileLight{}, fmt.Errorf("storage GetProfileLight error: %v", err.Error())
	}
	return profile, nil
}
func (s *Storage) GetProfilesFromSearch(ctx context.Context, input string) ([]*models.ProfileLight, error) {

	const query = `
		SELECT id,name,avatarImage 
		FROM profiles 
		WHERE name LIKE $1 OR email LIKE $1;
	`
	rows, err := s.db.Query(ctx, query, input)

	if err != nil {
		slog.Error("storage GetProfilesFromSearch error: " + err.Error())
		return nil, err
	}
	var result []*models.ProfileLight
	for rows.Next() {
		var profile models.ProfileLight
		err = rows.Scan(&profile.ID, &profile.Name, &profile.AvatarImage)
		if err != nil {
			slog.Error("storage GetProfilesFromSearch error: " + err.Error())
			return nil, err
		}
		result = append(result, &profile)
	}
	return result, nil
}
func (s *Storage) ToggleLikeIdea(ctx context.Context, userId, ideaId int64) (bool, int64, error) {
	slog.Info("storage started ToggleLikeIdea")

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	const query = `
		SELECT liked_ideas 
		FROM profiles 
		WHERE id = $1;
	`
	const queryUpdate = `
		UPDATE profiles 
		SET liked_ideas = $1 
		WHERE id = $2;
	`

	var likedIdeas string
	err = tx.QueryRow(ctx, query, userId).Scan(&likedIdeas)
	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}

	likedIdeasSlice, err := ParseIdsSqlite(likedIdeas)
	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}

	alreadyLiked := slices.Contains(likedIdeasSlice, ideaId)
	likesCountReponse, err := s.ideasClient.ChangeLikesCount(ctx, &ideasv1.ChangeLikesCountRequest{
		IdeaId:   ideaId,
		Increase: !alreadyLiked,
	})
	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}
	var newLikedIdeas string
	if alreadyLiked {
		newLikedIdeas = IdsSliceToString(likedIdeasSlice, ideaId)
	} else {
		likedIdeasSlice = append(likedIdeasSlice, ideaId)
		newLikedIdeas = IdsSliceToString(likedIdeasSlice, -1)
	}

	_, err = tx.Exec(ctx, queryUpdate, newLikedIdeas, userId)
	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}

	return !alreadyLiked, likesCountReponse.LikesCount, nil
}
func (s *Storage) UpdateProfile(ctx context.Context, userId int64, name, avatarImage, description, link string) (*emptypb.Empty, error) {
	slog.Info("storage started UpdateProfile")

	var query string
	if avatarImage != "" {
		query = "UPDATE profiles SET name = $1, avatarImage = $2, description = $3, link=$4 WHERE id = $5"
	} else {
		query = "UPDATE profiles SET name = $1, description = $3, link=$4 WHERE id = $5"
	}

	_, err := s.db.Exec(ctx, query, name, avatarImage, description, link, userId)
	if err != nil {
		slog.Error("storage UpdateProfile error: " + err.Error())
		return nil, fmt.Errorf("storage UpdateProfile error: %v", err.Error())
	}

	return nil, nil
}
func (s *Storage) ToggleSaveIdea(ctx context.Context, userId, ideaId, boardId int64) (bool, error) {
	slog.Info("storage start ToggleSaveIdea, boardId = " + fmt.Sprint(boardId))

	const selectQuery = `
		SELECT savedIdeas 
		FROM profiles 
		WHERE id=$1;
	`
	const updateQuery = `
		UPDATE profiles 
		SET savedIdeas = $1
		WHERE id=$2;
	`
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		slog.Error("storage ToggleSaveIdea Begin error: " + err.Error())
		return false, fmt.Errorf("storage ToggleSaveIdea Begin error: %v", err.Error())
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var ideasIdsStr string
	err = tx.QueryRow(ctx, selectQuery, userId).Scan(&ideasIdsStr)
	if err != nil {
		slog.Error("storage ToggleSaveIdea error: " + err.Error())
		return false, fmt.Errorf("storage ToggleSaveIdea error: %v", err.Error())
	}

	idsPairs, err := ParseIdPairs(ideasIdsStr)
	if err != nil {
		slog.Error("storage ToggleSaveIdea error: " + err.Error())
		return false, fmt.Errorf("storage ToggleSaveIdea error: %v", err.Error())
	}

	pair := getPairByIdea(idsPairs, ideaId)
	if pair != nil {
		var newIdeasStr string
		if len(idsPairs) == 1 {
			newIdeasStr = ""
		} else {
			newIdeasStr = strings.Trim(strings.Replace(ideasIdsStr, pair.toString(), "", 1), " ")
			newIdeasStr = strings.ReplaceAll(newIdeasStr, "  ", " ")
		}

		_, err = tx.Exec(ctx, updateQuery, newIdeasStr, userId)
		if err != nil {
			slog.Error("storage ToggleSaveIdea error: " + err.Error())
			return false, fmt.Errorf("storage ToggleSaveIdea error: %v", err.Error())
		}

		if boardId != -1 {
			_, err := s.boardsClient.SetIdeaSaved(ctx, &boardsv1.SetIdeaSavedRequest{
				IdeaId:  ideaId,
				BoardId: boardId,
				Saved:   false,
			})
			if err != nil {
				return false, err
			}
		}

		return false, nil
	} else {
		var newIdeasStr string
		pair := ideaBoardPair{ideaId: ideaId, boardId: boardId}
		if len(ideasIdsStr) == 0 {
			newIdeasStr = pair.toString()
		} else {
			newIdeasStr = ideasIdsStr + " " + pair.toString()
		}

		_, err = tx.Exec(ctx, updateQuery, newIdeasStr, userId)
		if err != nil {
			slog.Error("storage ToggleSaveIdea error: " + err.Error())
			return false, fmt.Errorf("storage ToggleSaveIdea error: %v", err.Error())
		}

		if boardId != -1 {
			_, err := s.boardsClient.SetIdeaSaved(ctx, &boardsv1.SetIdeaSavedRequest{
				IdeaId:  ideaId,
				BoardId: boardId,
				Saved:   true,
			})
			if err != nil {
				return false, err
			}
		}

		return true, nil
	}
}

func (s *Storage) IsIdeaSaved(ctx context.Context, userId, ideaId int64) (bool, int64, error) {
	slog.Info("storage start IsIdeaSaved")

	const query = `
		SELECT savedIdeas 
		FROM profiles 
		WHERE id = $1;
	`
	var ideasStr string
	err := s.db.QueryRow(ctx, query, userId).Scan(&ideasStr)
	if err != nil {
		slog.Error("storage IsIdeaSaved error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage IsIdeaSaved error: %v", err.Error())
	}

	pairsSlice, err := ParseIdPairs(ideasStr)
	if err != nil {
		slog.Error("storage IsIdeaSaved error: " + err.Error())
		return false, -1, err
	}

	pair := getPairByIdea(pairsSlice, ideaId)
	if pair == nil {
		return false, -1, nil
	} else {
		return true, pair.boardId, nil
	}
}

func (s *Storage) IsIdeaLiked(ctx context.Context, userId, ideaId int64) (bool, error) {
	slog.Info("storage started IsIdeaLiked")

	const query = `
		SELECT liked_ideas 
		FROM profiles 
		WHERE id = $1;
	`
	var likedIdeas string
	err := s.db.QueryRow(ctx, query, userId).Scan(&likedIdeas)
	if err != nil {
		slog.Error("storage IsIdeaLiked error: " + err.Error())
		return false, fmt.Errorf("storage IsIdeaLiked error: %v", err.Error())
	}

	likedIdeasSlice, err := ParseIdsSqlite(likedIdeas)
	if err != nil {
		slog.Error("storage IsIdeaLiked error: " + err.Error())
		return false, fmt.Errorf("storage IsIdeaLiked error: %v", err.Error())
	}
	return slices.Contains(likedIdeasSlice, ideaId), nil
}

func (s *Storage) GetSavedIdeas(ctx context.Context, userId int64) ([]*profilesv1.IdeaData, error) {
	slog.Info("storage start GetSavedIdeas")

	const query = `
		SELECT savedIdeas 
		FROM profiles 
		WHERE id = $1;
	`
	var ideasStr string
	err := s.db.QueryRow(ctx, query, userId).Scan(&ideasStr)
	if err != nil {
		slog.Error("storage GetSavedIdeas error: " + err.Error())
		return nil, fmt.Errorf("storage GetSavedIdeas error: %v", err.Error())
	}

	pairsSlice, err := ParseIdPairs(ideasStr)
	if err != nil {
		slog.Error("storage GetSavedIdeas error: " + err.Error())
		return nil, fmt.Errorf("storage GetSavedIdeas error: %v", err.Error())
	}

	ideas := make([]*profilesv1.IdeaData, 0, 20)
	for _, pair := range pairsSlice {

		resp, err := s.ideasClient.GetIdea(ctx, &ideasv1.GetRequest{
			IdeaId: pair.ideaId,
		})
		if err != nil {
			slog.Error("storage IsIdeaSaved error: " + err.Error())
			return nil, err
		}
		ideas = append(ideas, &profilesv1.IdeaData{
			Id:      pair.ideaId,
			Name:    resp.Name,
			Image:   resp.Image,
			BoardId: pair.boardId,
		})
	}
	return ideas, nil
}

func (s *Storage) GetSavedIdeasIds(ctx context.Context, userId int64) ([]int64, error) {
	slog.Info("storage started GetSavedIdeasIds")

	const query = `
		SELECT savedIdeas 
		FROM profiles 
		WHERE id = $1;
	`

	var idsStr string
	err := s.db.QueryRow(ctx, query, userId).Scan(idsStr)
	if err != nil {
		slog.Error("storage GetSavedIdeasIds error: " + err.Error())
		return nil, fmt.Errorf("storage GetSavedIdeasIds error: " + err.Error())
	}

	pairs, err := ParseIdPairs(idsStr)
	if err != nil {
		slog.Error("storage GetSavedIdeasIds error: " + err.Error())
		return nil, fmt.Errorf("storage GetSavedIdeasIds error: " + err.Error())
	}

	return getIdeasIds(pairs), nil
}
func (s *Storage) MoveIdeasToBoard(ctx context.Context, userId, oldBoardId, newBoardId int64) (*emptypb.Empty, error) {
	slog.Info("storage started MoveIdeasToBoard")

	const selectQuery = `
		SELECT savedIdeas 
		FROM profiles 
		WHERE id = $1;
	`
	const updateQuery = `
		UPDATE profiles 
		SET savedIdeas = $1
		WHERE id = $2;
	`
	var idsStr string
	err := s.db.QueryRow(ctx, selectQuery, userId).Scan(&idsStr)
	if err != nil {
		slog.Error("storage MoveIdeasToBoard error: " + err.Error())
		return nil, fmt.Errorf("storage MoveIdeasToBoard error: " + err.Error())
	}

	pairs, err := ParseIdPairs(idsStr)
	if err != nil {
		slog.Error("storage MoveIdeasToBoard error: " + err.Error())
		return nil, fmt.Errorf("storage MoveIdeasToBoard error: " + err.Error())
	}

	var resultStr string
	for _, pair := range pairs {
		if pair.boardId == oldBoardId {
			pair.boardId = newBoardId
		}
		resultStr += pair.toString() + " "
	}
	resultStr = strings.TrimSpace(resultStr)

	_, err = s.db.Exec(ctx, updateQuery, resultStr, userId)
	if err != nil {
		slog.Error("storage MoveIdeasToBoard error: " + err.Error())
		return nil, fmt.Errorf("storage MoveIdeasToBoard error: " + err.Error())
	}

	return nil, nil
}

func (s *Storage) AddBoardToProfile(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {
	slog.Info("storage started AddBoardToProfile")
	const selectQuery = `
		SELECT boards 
		FROM profiles 
		WHERE id = $1;
	`
	const updateQuery = `
		UPDATE profiles 
		SET boards = $1
		WHERE id = $2;
	`

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var boardsStr string
	err = tx.QueryRow(ctx, selectQuery, userId).Scan(&boardsStr)
	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}

	boardsStr += " " + fmt.Sprint(boardId)
	boardsStr = strings.TrimSpace(boardsStr)
	_, err = tx.Exec(ctx, updateQuery, boardsStr, userId)
	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}

	return nil, nil
}
func (s *Storage) RemoveBoardFromProfile(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {
	slog.Info("storage started RemoveBoardFromProfile")

	const selectQuery = `
		SELECT boards 
		FROM profiles 
		WHERE id = $1;
	`
	const updateQuery = `
		UPDATE profiles 
		SET boards = $1 
		WHERE id = $2;
	`
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		slog.Error("storage RemoveBoardFromProfile error: " + err.Error())
		return nil, fmt.Errorf("storage RemoveBoardFromProfile error: " + err.Error())
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var boardsStr string
	err = tx.QueryRow(ctx, selectQuery, userId).Scan(&boardsStr)
	if err != nil {
		slog.Error("storage RemoveBoardFromProfile error: " + err.Error())
		return nil, fmt.Errorf("storage RemoveBoardFromProfile error: " + err.Error())
	}

	boardsStr = strings.Replace(boardsStr, fmt.Sprint(boardId), "", 1)
	boardsStr = strings.ReplaceAll(boardsStr, "  ", " ")
	boardsStr = strings.TrimSpace(boardsStr)
	_, err = tx.Exec(ctx, updateQuery, boardsStr, userId)
	if err != nil {
		slog.Error("storage RemoveBoardFromProfile error: " + err.Error())
		return nil, fmt.Errorf("storage RemoveBoardFromProfile error: " + err.Error())
	}

	return nil, nil
}

type ideaBoardPair struct {
	ideaId  int64
	boardId int64
}

func (p *ideaBoardPair) toString() string {
	return fmt.Sprint(p.ideaId) + ":" + fmt.Sprint(p.boardId)
}
func getIdeasIds(slice []ideaBoardPair) []int64 {
	var ids []int64
	for _, pair := range slice {
		ids = append(ids, pair.ideaId)
	}
	return ids
}
func getPairByIdea(slice []ideaBoardPair, id int64) *ideaBoardPair {
	for _, pair := range slice {
		if pair.ideaId == id {
			return &pair
		}
	}
	return nil
}
func parseIdeaBoardPair(s string) (ideaBoardPair, error) {
	slice := strings.Split(s, ":")
	ideaId, err := strconv.ParseInt(slice[0], 10, 64)
	if err != nil {
		return ideaBoardPair{}, fmt.Errorf("error parsing idea-board :" + err.Error())
	}
	boardId, err := strconv.ParseInt(slice[1], 10, 64)
	if err != nil {
		return ideaBoardPair{}, fmt.Errorf("error parsing idea-board :" + err.Error())
	}
	return ideaBoardPair{
		ideaId:  ideaId,
		boardId: boardId,
	}, nil
}
func toGRPCFormat(slice []ideaBoardPair) []*profilesv1.IdeaBoardPair {
	var result []*profilesv1.IdeaBoardPair
	for _, pair := range slice {
		result = append(result, &profilesv1.IdeaBoardPair{
			IdeaId:  pair.ideaId,
			BoardId: pair.boardId,
		})
	}
	return result
}
