package sqlite

import (
	"context"
	"fmt"
	"idea-store-auth/internal/domain/models"
	"log/slog"
	"slices"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateProfile(ctx context.Context, id int64,name,email string)  (*emptypb.Empty,error){
	slog.Info("storage start CreateProfile")

	stmt,err:= s.db.Prepare("INSERT INTO profiles(id,email,avatarImage,name,description,link,boards,savedIdeas) VALUES (?,?,?,?,?,?,?,?);")

	if err!=nil{
		slog.Error("storage CreateProfile error: "+err.Error())
		return nil,err
	}
	
	_, err = stmt.ExecContext(ctx,id,email,"",name,"","","","")

	if err != nil {
		slog.Error("storage CreateProfile error: "+err.Error())
		return nil, fmt.Errorf("storage CreateProfile error: %w", err)
	}
	return &emptypb.Empty{},nil
}

func (s *Storage) GetProfile(ctx context.Context, id int64) (models.Profile, error){
	slog.Info("storage start GetProfile")

	stmt,err := s.db.Prepare("SELECT id,email,avatarImage,name,description,link,boards,savedIdeas FROM profiles WHERE id=?")

	if err!=nil{
		slog.Error("storage GetProfile db Prepare error: "+err.Error())
		return models.Profile{}, err
	}

	row := stmt.QueryRowContext(ctx,id)
	var profile models.Profile
	var boardsStr string
	var ideasStr string
	err = row.Scan(&profile.ID, &profile.Email, &profile.AvatarImage, &profile.Name, &profile.Description, &profile.Link, &boardsStr, &ideasStr)
	if err!=nil{
		slog.Error("storage GetProfile error scaninng row: "+err.Error())
		return models.Profile{}, err
	}

	profile.SavedIdeas,err = ParseIdsSqlite(ideasStr)
	if err!=nil{
		slog.Error("storage GetProfile error parsing ideas ids: "+err.Error())
		return models.Profile{}, err
	}
	profile.Boards,err = ParseIdsSqlite(boardsStr)

	if err!=nil{
		slog.Error("storage GetProfile error parsing boards ids: "+err.Error())
		return models.Profile{}, err
	}

	return profile,nil
}

func (s* Storage) ToggleSaveIdea(ctx context.Context, userId,ideaId, boardId int64) (bool,error){
	slog.Info("storage start ToggleSaveIdea")

	tx,err:= s.db.Begin()	
	if err!=nil{
		slog.Error("storage ToggleSaveIdea Begin error: "+err.Error())
		return false, err
	}

	stmt,err := tx.Prepare("SELECT savedIdeas FROM profiles WHERE id=?")
	if err!=nil{
		slog.Error("storage ToggleSaveIdea SELECT error: "+err.Error())
		return false, err
	}
	row:= stmt.QueryRowContext(ctx,userId)
	var ideasIdsStr string
	var ideasIds []int64
	err= row.Scan(&ideasIdsStr)
	if err!=nil{
		slog.Error("storage ToggleSaveIdea Scan error: "+err.Error())
		return false, err
	}
	ideasIds,err = ParseIdsSqlite(ideasIdsStr)
	if err!=nil{
		slog.Error("storage ToggleSaveIdea ideas parse error: "+err.Error())
		return false,err
	}
	if slices.Contains(ideasIds,ideaId){
		var newIdeasStr string
		if(string(ideasIdsStr[0])==fmt.Sprint(ideaId)){
			newIdeasStr = strings.Replace(ideasIdsStr,fmt.Sprint(ideaId),"",1)
		}else{
			newIdeasStr = strings.Replace(ideasIdsStr," "+fmt.Sprint(ideaId),"",1)
		}
		stmt,err = tx.Prepare("UPDATE profiles SET savedIdeas = ? WHERE id=?")
		if err!=nil{
			slog.Error("storage ToggleSaveIdea UPDATE(remove) error: "+err.Error())
			return false, err
		}
		stmt.ExecContext(ctx,newIdeasStr,userId)
		tx.Commit()
		//TODO: убрать идею с доски
		return false,nil
	}else{
		var newIdeasStr string
		if(len(ideasIdsStr)==0){
			newIdeasStr = fmt.Sprint(ideaId)
		}else{
			newIdeasStr = ideasIdsStr +" "+fmt.Sprint(ideaId)
		}
		
		stmt,err = tx.Prepare("UPDATE profiles SET savedIdeas = ? WHERE id=?")
		if err!=nil{
			slog.Error("storage ToggleSaveIdea UPDATE(add) error: "+err.Error())
			return false, err
		}
		stmt.ExecContext(ctx,newIdeasStr,userId)
		tx.Commit()
		//TODO: Добавить идею на доску
		return true,nil
	}
}