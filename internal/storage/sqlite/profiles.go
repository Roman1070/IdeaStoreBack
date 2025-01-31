package sqlite

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

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateProfile(ctx context.Context, id int64, name, email string) (*emptypb.Empty, error) {
	slog.Info("storage start CreateProfile")

	stmt, err := s.db.Prepare("INSERT INTO profiles(id,email,avatarImage,name,description,link,boards,savedIdeas) VALUES (?,?,?,?,?,?,?,?);")

	if err != nil {
		slog.Error("storage CreateProfile error: " + err.Error())
		return nil, err
	}

	_, err = stmt.ExecContext(ctx, id, email, "", name, "", "", "", "")

	if err != nil {
		slog.Error("storage CreateProfile error: " + err.Error())
		return nil, fmt.Errorf("storage CreateProfile error: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *Storage) GetProfile(ctx context.Context, id int64) (models.Profile, error) {
	slog.Info("storage start GetProfile")

	stmt, err := s.db.Prepare("SELECT id,email,avatarImage,name,description,link,boards,savedIdeas FROM profiles WHERE id=?")

	if err != nil {
		slog.Error("storage GetProfile db Prepare error: " + err.Error())
		return models.Profile{}, err
	}

	row := stmt.QueryRowContext(ctx, id)
	var profile models.Profile
	var boardsStr string
	var ideasStr string
	err = row.Scan(&profile.ID, &profile.Email, &profile.AvatarImage, &profile.Name, &profile.Description, &profile.Link, &boardsStr, &ideasStr)
	if err != nil {
		slog.Error("storage GetProfile error scaninng row: " + err.Error())
		return models.Profile{}, err
	}
	pairsSlice, err := ParseIdPairs(ideasStr)
	if err != nil {
		slog.Error("storage GetProfile error parsing ideas-board: " + err.Error())
		return models.Profile{}, err
	}
	profile.SavedIdeas = toGRPCFormat(pairsSlice)
	profile.Boards, err = ParseIdsSqlite(boardsStr)

	if err != nil {
		slog.Error("storage GetProfile error parsing boards ids: " + err.Error())
		return models.Profile{}, err
	}

	return profile, nil
}
func (s *Storage) GetProfileLight(ctx context.Context, id int64) (models.ProfileLight, error) {
	slog.Info("storage start GetProfileLight")

	stmt, err := s.db.Prepare("SELECT id,name,avatarImage FROM profiles WHERE id=?")

	if err != nil {
		slog.Error("storage GetProfileLight db Prepare error: " + err.Error())
		return models.ProfileLight{}, err
	}

	row := stmt.QueryRowContext(ctx, id)
	var profile models.ProfileLight

	err = row.Scan(&profile.ID, &profile.Name, &profile.AvatarImage)
	if err != nil {
		slog.Error("storage GetProfile db Prepare error: " + err.Error())
		return models.ProfileLight{}, err
	}
	return profile, nil
}
func (s *Storage) GetProfilesFromSearch(ctx context.Context, input string) ([]*models.ProfileLight, error) {
	slog.Info("storage start GetProfilesFromSearch")

	stmt, err := s.db.Prepare("SELECT id,name,avatarImage FROM profiles WHERE name LIKE ?  OR email LIKE ?")
	if err != nil {
		slog.Error("storage GetProfilesFromSearch db Prepare error: " + err.Error())
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, input+"%", input+"%")
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

	tx, err := s.db.Begin()

	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}
	stmt, err := tx.Prepare("SELECT liked_ideas FROM profiles WHERE id = ?")

	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}
	row := stmt.QueryRowContext(ctx, userId)
	var likedIdeas string
	err = row.Scan(&likedIdeas)

	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}
	likedIdeasSlice, err := ParseIdsSqlite(likedIdeas)

	if err != nil {
		slog.Error("storage ToggleLikeIdea error: " + err.Error())
		return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
	}
	fmt.Printf("liked ideas slice : \"%v\"\n", likedIdeasSlice)
	if slices.Contains(likedIdeasSlice, ideaId) {
		likesCountReponse, err := s.ideasClient.ChangeLikesCount(ctx, &ideasv1.ChangeLikesCountRequest{
			IdeaId:   ideaId,
			Increase: false,
		})
		if err != nil {
			slog.Error("storage ToggleLikeIdea error: " + err.Error())
			return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
		}
		newLikedIdeas := IdsSliceToString(likedIdeasSlice, ideaId)
		fmt.Printf("newLikedIdeas : \"%v\"\n", newLikedIdeas)
		stmt, err = tx.Prepare("UPDATE profiles SET liked_ideas = ? WHERE id = ?")

		if err != nil {
			slog.Error("storage ToggleLikeIdea error: " + err.Error())
			return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
		}
		_, err = stmt.ExecContext(ctx, newLikedIdeas, userId)

		if err != nil {
			slog.Error("storage ToggleLikeIdea error: " + err.Error())
			return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
		}
		err = tx.Commit()

		if err != nil {
			slog.Error("storage ToggleLikeIdea error: " + err.Error())
			return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
		}
		return false, likesCountReponse.LikesCount, nil
	} else {
		likesCountReponse, err := s.ideasClient.ChangeLikesCount(ctx, &ideasv1.ChangeLikesCountRequest{
			IdeaId:   ideaId,
			Increase: true,
		})
		if err != nil {
			slog.Error("storage ToggleLikeIdea error: " + err.Error())
			return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
		}
		likedIdeasSlice = append(likedIdeasSlice, ideaId)
		newLikedIdeas := IdsSliceToString(likedIdeasSlice, -1)
		fmt.Printf("newLikedIdeas : \"%v\"\n", newLikedIdeas)
		stmt, err = tx.Prepare("UPDATE profiles SET liked_ideas = ? WHERE id = ?")

		if err != nil {
			slog.Error("storage ToggleLikeIdea error: " + err.Error())
			return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
		}
		_, err = stmt.ExecContext(ctx, newLikedIdeas, userId)

		if err != nil {
			slog.Error("storage ToggleLikeIdea error: " + err.Error())
			return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
		}
		err = tx.Commit()

		if err != nil {
			slog.Error("storage ToggleLikeIdea error: " + err.Error())
			return false, emptyValue, fmt.Errorf("storage ToggleLikeIdea error: " + err.Error())
		}
		return true, likesCountReponse.LikesCount, nil
	}
}
func (s *Storage) UpdateProfile(ctx context.Context, userId int64, name, avatarImage, description, link string) (*emptypb.Empty, error) {
	slog.Info("storage started UpdateProfile")
	var query string
	if avatarImage != "" {
		query = "UPDATE profiles SET name = ?, avatarImage = ?, description = ?, link=? WHERE id = ?"
	} else {
		query = "UPDATE profiles SET name = ?, description = ?, link=? WHERE id = ?"
	}
	stmt, err := s.db.Prepare(query)
	if err != nil {
		slog.Error("storage UpdateProfile error: " + err.Error())
		return nil, err
	}
	if avatarImage != "" {
		_, err = stmt.ExecContext(ctx, name, avatarImage, description, link, userId)
	} else {
		_, err = stmt.ExecContext(ctx, name, description, link, userId)
	}

	if err != nil {
		slog.Error("storage UpdateProfile error: " + err.Error())
		return nil, err
	}

	return nil, nil
}
func (s *Storage) ToggleSaveIdea(ctx context.Context, userId, ideaId, boardId int64) (bool, error) {
	slog.Info("storage start ToggleSaveIdea, boardId = " + fmt.Sprint(boardId))

	tx, err := s.db.Begin()
	if err != nil {
		slog.Error("storage ToggleSaveIdea Begin error: " + err.Error())
		return false, err
	}

	stmt, err := tx.Prepare("SELECT savedIdeas FROM profiles WHERE id=?")
	if err != nil {
		slog.Error("storage ToggleSaveIdea SELECT error: " + err.Error())
		return false, err
	}
	row := stmt.QueryRowContext(ctx, userId)
	var ideasIdsStr string
	var idsPairs []ideaBoardPair
	err = row.Scan(&ideasIdsStr)
	if err != nil {
		slog.Error("storage ToggleSaveIdea Scan error: " + err.Error())
		return false, err
	}
	idsPairs, err = ParseIdPairs(ideasIdsStr)
	if err != nil {
		slog.Error("storage ToggleSaveIdea ideas parse error: " + err.Error())
		return false, err
	}
	fmt.Println("ids pairs: " + fmt.Sprint(idsPairs))
	pair := getPairByIdea(idsPairs, ideaId)
	if pair != nil {
		var newIdeasStr string
		if len(idsPairs) == 1 {
			newIdeasStr = ""
		} else {
			newIdeasStr = strings.Trim(strings.Replace(ideasIdsStr, pair.toString(), "", 1), " ")
			newIdeasStr = strings.ReplaceAll(newIdeasStr, "  ", " ")
		}
		stmt, err = tx.Prepare("UPDATE profiles SET savedIdeas = ? WHERE id=?")
		if err != nil {
			slog.Error("storage ToggleSaveIdea UPDATE(remove) error: " + err.Error())
			return false, err
		}
		stmt.ExecContext(ctx, newIdeasStr, userId)
		tx.Commit()
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

		stmt, err = tx.Prepare("UPDATE profiles SET savedIdeas = ? WHERE id=?")
		if err != nil {
			slog.Error("storage ToggleSaveIdea UPDATE(add) error: " + err.Error())
			return false, err
		}
		stmt.ExecContext(ctx, newIdeasStr, userId)
		tx.Commit()
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

	stmt, err := s.db.Prepare("SELECT savedIdeas FROM profiles WHERE id = ?")
	if err != nil {
		slog.Error("storage IsIdeaSaved Prepare error: " + err.Error())
		return false, -1, err
	}
	row := stmt.QueryRowContext(ctx, userId)
	var ideasStr string
	err = row.Scan(&ideasStr)
	if err != nil {
		slog.Error("storage IsIdeaSaved error: " + err.Error())
		return false, -1, err
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

	stmt, err := s.db.Prepare("SELECT liked_ideas from profiles WHERE id = ?")
	if err != nil {
		slog.Error("storage IsIdeaLiked error: " + err.Error())
		return false, fmt.Errorf("storage IsIdeaLiked error: %v", err.Error())
	}

	row := stmt.QueryRowContext(ctx, userId)
	var likedIdeas string
	err = row.Scan(&likedIdeas)
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

	stmt, err := s.db.Prepare("SELECT savedIdeas FROM profiles WHERE id = ?")
	if err != nil {
		slog.Error("storage GetSavedIdeas Prepare error: " + err.Error())
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, userId)
	var ideasStr string
	err = row.Scan(&ideasStr)
	if err != nil {
		slog.Error("storage IsIdeaSaved error: " + err.Error())
		return nil, err
	}
	pairsSlice, err := ParseIdPairs(ideasStr)
	if err != nil {
		slog.Error("storage IsIdeaSaved error: " + err.Error())
		return nil, err
	}
	var ideas []*profilesv1.IdeaData
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

	stmt, err := s.db.Prepare("SELECT savedIdeas FROM profiles WHERE id = ?")
	if err != nil {
		slog.Error("storage GetSavedIdeasIds error: " + err.Error())
		return nil, fmt.Errorf("storage GetSavedIdeasIds error: " + err.Error())
	}

	row := stmt.QueryRow(userId)
	var idsStr string
	err = row.Scan(&idsStr)
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

	stmt, err := s.db.Prepare("SELECT savedIdeas FROM profiles WHERE id = ?")
	if err != nil {
		slog.Error("storage MoveIdeasToBoard error: " + err.Error())
		return nil, fmt.Errorf("storage MoveIdeasToBoard error: " + err.Error())
	}

	row := stmt.QueryRow(userId)
	var idsStr string
	err = row.Scan(&idsStr)
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
	stmt, err = s.db.Prepare("UPDATE profiles SET savedIdeas = ?  WHERE id = ?")
	if err != nil {
		slog.Error("storage MoveIdeasToBoard error: " + err.Error())
		return nil, fmt.Errorf("storage MoveIdeasToBoard error: " + err.Error())
	}
	_, err = stmt.ExecContext(ctx, resultStr, userId)

	if err != nil {
		slog.Error("storage MoveIdeasToBoard error: " + err.Error())
		return nil, fmt.Errorf("storage MoveIdeasToBoard error: " + err.Error())
	}
	return nil, nil
}

func (s *Storage) AddBoardToProfile(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {

	slog.Info("storage started AddBoardToProfile")
	tx, err := s.db.Begin()
	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}
	stmt, err := tx.Prepare("SELECT boards FROM profiles WHERE id = ?")
	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}
	row := stmt.QueryRowContext(ctx, userId)
	var boardsStr string
	err = row.Scan(&boardsStr)
	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}
	boardsStr += " " + fmt.Sprint(boardId)
	boardsStr = strings.TrimSpace(boardsStr)
	stmt, err = tx.Prepare("UPDATE profiles SET boards = ? WHERE id = ?")
	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}
	_, err = stmt.ExecContext(ctx, boardsStr, userId)
	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}
	err = tx.Commit()

	if err != nil {
		slog.Error("storage AddBoardToProfile error: " + err.Error())
		return nil, fmt.Errorf("storage AddBoardToProfile error: " + err.Error())
	}
	return nil, nil
}
func (s *Storage) RemoveBoardFromProfile(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {

	slog.Info("storage started RemoveBoardFromProfile")
	tx, err := s.db.Begin()
	if err != nil {
		slog.Error("storage RemoveBoardFromProfile error: " + err.Error())
		return nil, fmt.Errorf("storage RemoveBoardFromProfile error: " + err.Error())
	}
	stmt, err := tx.Prepare("SELECT boards FROM profiles WHERE id = ?")
	if err != nil {
		slog.Error("storage RemoveBoardFromProfile error: " + err.Error())
		return nil, fmt.Errorf("storage RemoveBoardFromProfile error: " + err.Error())
	}
	row := stmt.QueryRowContext(ctx, userId)
	var boardsStr string
	err = row.Scan(&boardsStr)
	if err != nil {
		slog.Error("storage RemoveBoardFromProfile error: " + err.Error())
		return nil, fmt.Errorf("storage RemoveBoardFromProfile error: " + err.Error())
	}
	boardsStr = strings.Replace(boardsStr, fmt.Sprint(boardId), "", 1)
	boardsStr = strings.ReplaceAll(boardsStr, "  ", " ")
	boardsStr = strings.TrimSpace(boardsStr)
	stmt, err = tx.Prepare("UPDATE profiles SET boards = ? WHERE id = ?")
	if err != nil {
		slog.Error("storage RemoveBoardFromProfile error: " + err.Error())
		return nil, fmt.Errorf("storage RemoveBoardFromProfile error: " + err.Error())
	}
	_, err = stmt.ExecContext(ctx, boardsStr, userId)
	if err != nil {
		slog.Error("storage RemoveBoardFromProfile error: " + err.Error())
		return nil, fmt.Errorf("storage RemoveBoardFromProfile error: " + err.Error())
	}
	err = tx.Commit()

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
