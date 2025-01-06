package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	boardsv1 "idea-store-auth/gen/go/boards"
	ideasv1 "idea-store-auth/gen/go/idea"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/storage"
	"log/slog"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateBoard(ctx context.Context, name string, userId int64) (int64, error) {
	const op = "storage.sqlite.CreateBoard"

	stmt, err := s.db.Prepare("INSERT INTO boards(name,ideas_ids,user_id) VALUES(?,?,?)")
	if err != nil {
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, name, "",userId)

	if err != nil {
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *Storage) GetBoard(ctx context.Context, id int64) (models.Board, error) {

	const op = "storage.sqlite.CreateBoard"

	stmt, err := s.db.Prepare("SELECT id,name,ideas_ids FROM boards WHERE id=?")
	if err != nil {
		slog.Error(err.Error())
		return models.Board{}, fmt.Errorf("%s: %w", op, err)
	}
	var board models.Board
	var idsString string
	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&board.ID, &board.Name, &idsString)
	if err != nil {
		slog.Error(err.Error())
		return models.Board{}, fmt.Errorf("internal error %v", err.Error())
	}
	var ids []int64
	if len(idsString) > 0 {
		ids, err = ParseIdsSqlite(idsString)
		if err != nil {
			slog.Error(err.Error())
			return models.Board{}, fmt.Errorf("internal error %v", err.Error())
		}
	}
	board.IdeasIds = ids
	return board, nil
}
func (s *Storage) SetIdeaSaved(ctx context.Context, boardId, ideaId int64, saved bool) (*emptypb.Empty, error) {
	slog.Info("storage started SetIdeaSaved")
	tx, err := s.db.Begin()
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("storage error SetIdeaSaved : %v", err.Error())
	}
	stmt, err := tx.Prepare("SELECT ideas_ids FROM boards WHERE id = ?")
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("storage error SetIdeaSaved : %v", err.Error())
	}
	var idsString string
	var newIdsString string
	row := stmt.QueryRowContext(ctx, boardId)
	err = row.Scan(&idsString)

	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("internal error %v", err.Error())
	}
	if saved {
		newIdsString = idsString + " " + fmt.Sprint(ideaId)
	} else {
		newIdsString = strings.Replace(idsString, fmt.Sprint(ideaId), "", 1)
	}
	newIdsString = strings.Trim(strings.ReplaceAll(newIdsString, "  ", " "), " ")
	stmt, err = tx.Prepare("UPDATE boards SET ideas_ids = ? WHERE id = ?")
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("storage error SetIdeaSaved : %v", err.Error())
	}
	_, err = stmt.ExecContext(ctx, newIdsString, boardId)

	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("storage error SetIdeaSaved : %v", err.Error())
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}
func (s *Storage) GetAllBoards(ctx context.Context, userId int64) ([]*boardsv1.BoardData, error) {
	const op = "storage.sqlite.GetAllBoards"

	stmt, err := s.db.Prepare("SELECT id,name,ideas_ids FROM boards WHERE user_id = ?")
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	rows, err := stmt.QueryContext(ctx,userId)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	boards := []*boardsv1.BoardData{}
	for rows.Next() {
		board := new(boardsv1.BoardData)
		var ideasStr string
		err = rows.Scan(&board.Id, &board.Name, &ideasStr)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				slog.Error(err.Error())
				return nil, fmt.Errorf("%s: %w", op, storage.ErrBoardNotFound)
			}
			slog.Error(err.Error())
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		var ids []int64
		if len(ideasStr) > 0 {
			ids, err = ParseIdsSqlite(ideasStr)
			if err != nil {
				slog.Error(err.Error())
				slog.Error(err.Error())
				return nil, fmt.Errorf("internal error %v", err.Error())
			}
		}
		board.IdeasIds = ids
		boards = append(boards, board)
	}
	return boards, nil
}

func (s *Storage) GetIdeasInBoard(ctx context.Context, boardId int64)([]*boardsv1.IdeaData, error){
	const op = "storage.sqlite.GetAllBoards"

	stmt, err := s.db.Prepare("SELECT ideas_ids FROM boards WHERE id = ?")
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	row:= stmt.QueryRowContext(ctx,boardId)
	var idsStr string
	err = row.Scan(&idsStr)
	if err!=nil{
		slog.Error("Error scanning ideas_ids")
		return nil, fmt.Errorf("internal storage error: %v",err.Error())
	}
	idsSlice,err:= ParseIdsSqlite(idsStr)
	if err!=nil{
		slog.Error("Error parsing ideas_ids")
		return nil, fmt.Errorf("internal storage error: %v",err.Error())
	}
	//TODO: сделать за один запрос
	var result []*boardsv1.IdeaData
	for _,id := range idsSlice{
		idea,err := s.ideasClient.GetIdea(ctx,&ideasv1.GetRequest{
			IdeaId: id,
		})
		if err!=nil{
			slog.Error("Error parsing ideas_ids")
			return nil, fmt.Errorf("internal storage error: %v",err.Error())
		}
		result = append(result, &boardsv1.IdeaData{
			IdeaId: id,
			Image: idea.Image,
			Name: idea.Name,
		})
	}

	return result,nil
}

func(s *Storage) DeleteBoard(ctx context.Context, userId, boardId int64)(*emptypb.Empty, error){
	const op = "storage.sqlite.DeleteBoard"

	stmt, err := s.db.Prepare("DELETE FROM boards WHERE id = ?")
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	_, err=stmt.ExecContext(ctx,boardId)
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	_, err = s.profilesClient.MoveIdeasToBoard(ctx,&profilesv1.MoveIdeaToBoardRequest{
		UserId: userId,
		OldBoardId: boardId,
		NewBoardId: -1,
	})
	if err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return nil,nil
}