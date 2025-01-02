package sqlite

import (
	"context"
	"fmt"
	boardsv1 "idea-store-auth/gen/go/boards"
	ideasv1 "idea-store-auth/gen/go/idea"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"log/slog"
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
	pairsSlice,err :=ParseIdPairs(ideasStr)
	if err != nil {
		slog.Error("storage GetProfile error parsing ideas-board: " + err.Error())
		return models.Profile{}, err
	}
	profile.SavedIdeas= toGRPCFormat(pairsSlice)
	profile.Boards, err = ParseIdsSqlite(boardsStr)

	if err != nil {
		slog.Error("storage GetProfile error parsing boards ids: " + err.Error())
		return models.Profile{}, err
	}

	return profile, nil
}
//TODO: если убрать на доске сохраненную идею, из доски она не уберется
func (s *Storage) ToggleSaveIdea(ctx context.Context, userId, ideaId, boardId int64) (bool, error) {
	slog.Info("storage start ToggleSaveIdea, boardId = "+fmt.Sprint(boardId))

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
	fmt.Println("ids pairs: "+fmt.Sprint(idsPairs))
	pair:=getPairByIdea(idsPairs, ideaId)
	if  pair!=nil {
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
		if boardId!=-1{
			s.boardsClient.SetIdeaSaved(ctx, &boardsv1.SetIdeaSavedRequest{
				IdeaId:  ideaId,
				BoardId: boardId,
				Saved:   false,
			})
		}
		return false, nil
	} else {
		var newIdeasStr string
			pair:= ideaBoardPair{ideaId: ideaId, boardId: boardId}
		if len(ideasIdsStr) == 0 {
			newIdeasStr = pair.toString()
		} else {
			newIdeasStr = ideasIdsStr + " " +  pair.toString()
		}

		stmt, err = tx.Prepare("UPDATE profiles SET savedIdeas = ? WHERE id=?")
		if err != nil {
			slog.Error("storage ToggleSaveIdea UPDATE(add) error: " + err.Error())
			return false, err
		}
		stmt.ExecContext(ctx, newIdeasStr, userId)
		tx.Commit()
		if boardId!=-1{
			s.boardsClient.SetIdeaSaved(ctx, &boardsv1.SetIdeaSavedRequest{
				IdeaId:  ideaId,
				BoardId: boardId,
				Saved:   true,
			})
		}
		return true, nil
	}
}

func (s *Storage) IsIdeaSaved(ctx context.Context, userId, ideaId int64) (bool, int64,error) {
	slog.Info("storage start IsIdeaSaved")

	stmt, err := s.db.Prepare("SELECT savedIdeas FROM profiles WHERE id = ?")
	if err != nil {
		slog.Error("storage IsIdeaSaved Prepare error: " + err.Error())
		return false, -1,err
	}
	row := stmt.QueryRowContext(ctx, userId)
	var ideasStr string
	err = row.Scan(&ideasStr)
	if err != nil {
		slog.Error("storage IsIdeaSaved error: " + err.Error())
		return false,-1, err
	}
	pairsSlice, err := ParseIdPairs(ideasStr)
	if err != nil {
		slog.Error("storage IsIdeaSaved error: " + err.Error())
		return false, -1, err
	}
	pair:= getPairByIdea(pairsSlice,ideaId)
	if pair==nil{
		return false,-1,nil
	}else{
		return true, pair.boardId, nil
	}
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
			IdeaId:      pair.ideaId,
			Name:        resp.Name,
			Description: resp.Description,
			Link:        resp.Link,
			Image:       resp.Image,
			Tags:        resp.Tags,
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
		return nil, fmt.Errorf("sotrage GetSavedIdeasIds error: " + err.Error())
	}

	row := stmt.QueryRow(userId)
	var idsStr string
	err = row.Scan(&idsStr)
	if err != nil {
		slog.Error("storage GetSavedIdeasIds error: " + err.Error())
		return nil, fmt.Errorf("sotrage GetSavedIdeasIds error: " + err.Error())
	}
	pairs, err := ParseIdPairs(idsStr)
	if err != nil {
		slog.Error("storage GetSavedIdeasIds error: " + err.Error())
		return nil, fmt.Errorf("sotrage GetSavedIdeasIds error: " + err.Error())
	}
	return getIdeasIds(pairs), nil
}



type ideaBoardPair struct{
	ideaId int64
	boardId int64
}

func (p *ideaBoardPair) toString() string{
	return fmt.Sprint(p.ideaId)+":"+fmt.Sprint(p.boardId)
}
func getIdeasIds(slice []ideaBoardPair) []int64{
	var ids []int64
	for _, pair := range slice{
		ids = append(ids, pair.ideaId)
	}
	return ids
}
func getPairByIdea(slice []ideaBoardPair, id int64) (*ideaBoardPair){
	for _,pair := range slice{
		if pair.ideaId == id{
			return &pair
		}
	}
	return nil
}
func parseIdeaBoardPair(s string) (ideaBoardPair, error){
	slice:= strings.Split(s,":")
	ideaId, err :=strconv.ParseInt(slice[0],10,64)
	if err!=nil{
		return ideaBoardPair{},fmt.Errorf("error parsing idea-board :" + err.Error())
	}
	boardId, err:= strconv.ParseInt(slice[1],10,64)
	if err!=nil{
		return ideaBoardPair{},fmt.Errorf("error parsing idea-board :" + err.Error())
	}
	return ideaBoardPair{
		ideaId: ideaId,
		boardId: boardId,
	},nil
}
func toGRPCFormat(slice []ideaBoardPair) ([]*profilesv1.IdeaBoardPair){
	var result []*profilesv1.IdeaBoardPair
	for _, pair := range slice{
		result = append(result, &profilesv1.IdeaBoardPair{
			IdeaId: pair.ideaId,
			BoardId: pair.boardId,
		})
	}
	return result
}