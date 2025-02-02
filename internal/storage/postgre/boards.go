package postgre

import (
	"context"
	"fmt"
	boardsv1 "idea-store-auth/gen/go/boards"
	ideasv1 "idea-store-auth/gen/go/idea"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"log/slog"
	"strings"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateBoard(ctx context.Context, name string, userId int64) (int64, error) {
	slog.Info("storage started to CreateBoard")

	const query = `
		INSERT INTO boards(name,ideas_ids,user_id) 
		VALUES($1,$2,$3)
		RETURNING id;
	`

	var lastInsertId int64
	err := s.db.QueryRow(ctx, query, name, "", userId).Scan(&lastInsertId)
	if err != nil {
		slog.Error("storage CreateBoard error: " + err.Error())
		return emptyValue, fmt.Errorf("storage CreateBoard error: %v", err.Error())
	}

	_, err = s.profilesClient.AddBoardToProfile(ctx, &profilesv1.AddBoardToProfileRequest{BoardId: lastInsertId, UserId: userId})
	if err != nil {
		slog.Error("storage CreateBoard error: " + err.Error())
		return emptyValue, fmt.Errorf("storage CreateBoard error: %v", err.Error())
	}

	return lastInsertId, nil
}

func (s *Storage) GetBoard(ctx context.Context, id int64) (models.Board, error) {
	slog.Info("storage started to GetBoard")

	const query = `
		SELECT id,name,ideas_ids 
		FROM boards 
		WHERE id=$1;
	`
	var board models.Board
	var idsString string

	err := s.db.QueryRow(ctx, query, id).Scan(&board.ID, &board.Name, &idsString)
	if err != nil {
		slog.Error("storage GetBoard error: " + err.Error())
		return models.Board{}, fmt.Errorf("storage GetBoard error: %v", err.Error())
	}

	var ids []int64
	if len(idsString) > 0 {
		ids, err = ParseIdsString(idsString)
		if err != nil {
			slog.Error("storage GetBoard error: " + err.Error())
			return models.Board{}, fmt.Errorf("storage GetBoard error: %v", err.Error())
		}
	}

	board.IdeasIds = ids
	return board, nil
}
func (s *Storage) SetIdeaSaved(ctx context.Context, boardId, ideaId int64, saved bool) (*emptypb.Empty, error) {
	slog.Info("storage started SetIdeaSaved")

	const selectQuery = `
		SELECT ideas_ids 
		FROM boards 
		WHERE id = $1;
	`
	const updateQuery = `
		UPDATE boards 
		SET ideas_ids = $1
		WHERE id = $2;
	`
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		slog.Error("storage SetIdeaSaved error: " + err.Error())
		return nil, fmt.Errorf("storage SetIdeaSaved error: %v", err.Error())
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var idsString string
	var newIdsString string
	err = tx.QueryRow(ctx, selectQuery, boardId).Scan(&idsString)
	if err != nil {
		slog.Error("storage SetIdeaSaved error: " + err.Error())
		return nil, fmt.Errorf("storage SetIdeaSaved error: %v", err.Error())
	}

	if saved {
		newIdsString = strings.TrimSpace(idsString + " " + fmt.Sprint(ideaId))
	} else {
		idsSlice, err := ParseIdsString(idsString)
		if err != nil {
			slog.Error("storage SetIdeaSaved error: " + err.Error())
			return nil, fmt.Errorf("storage SetIdeaSaved error: %v", err.Error())
		}

		newIdsString = IdsSliceToString(idsSlice, ideaId)
	}

	_, err = tx.Exec(ctx, updateQuery, newIdsString, boardId)
	if err != nil {
		slog.Error("storage SetIdeaSaved error: " + err.Error())
		return nil, fmt.Errorf("storage SetIdeaSaved error: %v", err.Error())
	}

	return &emptypb.Empty{}, nil
}
func (s *Storage) GetCurrentUsersBoards(ctx context.Context, userId int64) ([]*boardsv1.BoardData, error) {
	slog.Info("storage started to GetCurrentUsersBoards")

	const query = `
		SELECT id,name,ideas_ids 
		FROM boards 
		WHERE user_id = $1;
	`

	rows, err := s.db.Query(ctx, query, userId)
	if err != nil {
		slog.Error("storage GetCurrentUsersBoards error: " + err.Error())
		return nil, fmt.Errorf("storage GetCurrentUsersBoards error: %v", err.Error())
	}

	defer rows.Close()

	var boards []*boardsv1.BoardData
	for rows.Next() {
		board := new(boardsv1.BoardData)
		var ideasStr string
		err = rows.Scan(&board.Id, &board.Name, &ideasStr)
		if err != nil {
			slog.Error("storage GetCurrentUsersBoards error: " + err.Error())
			return nil, fmt.Errorf("storage GetCurrentUsersBoards error: %v", err.Error())
		}

		var ids []int64
		if len(ideasStr) > 0 {
			ids, err = ParseIdsString(ideasStr)
			if err != nil {
				slog.Error("storage GetCurrentUsersBoards error: " + err.Error())
				return nil, fmt.Errorf("storage GetCurrentUsersBoards error: %v", err.Error())
			}
		}

		board.IdeasIds = ids
		boards = append(boards, board)
	}
	return boards, nil
}

func (s *Storage) GetIdeasInBoard(ctx context.Context, boardId int64) ([]*boardsv1.IdeaData, error) {
	slog.Info("storage started to GetIdeasInBoard")

	const query = `
		SELECT ideas_ids 
		FROM boards 
		WHERE id = $1;
	`

	var idsStr string
	err := s.db.QueryRow(ctx, query, boardId).Scan(&idsStr)
	if err != nil {
		slog.Error("storage GetIdeasInBoard error: " + err.Error())
		return nil, fmt.Errorf("storage GetIdeasInBoard error: %v", err.Error())
	}

	idsSlice, err := ParseIdsString(idsStr)
	if err != nil {
		slog.Error("storage GetIdeasInBoard error: " + err.Error())
		return nil, fmt.Errorf("storage GetIdeasInBoard error: %v", err.Error())
	}

	ideas, err := s.ideasClient.GetIdeas(ctx, &ideasv1.GetIdeasRequest{
		Ids: idsSlice,
	})
	if err != nil {
		slog.Error("storage GetIdeasInBoard error: " + err.Error())
		return nil, fmt.Errorf("storage GetIdeasInBoard error: %v", err.Error())
	}

	var result []*boardsv1.IdeaData
	for _, idea := range ideas.Ideas {
		result = append(result, &boardsv1.IdeaData{
			Id:    idea.Id,
			Image: idea.Image,
			Name:  idea.Name,
		})
	}

	return result, nil
}

func (s *Storage) DeleteBoard(ctx context.Context, userId, boardId int64) (*emptypb.Empty, error) {
	slog.Info("storage started to DeleteBoard")

	const query = `
		DELETE FROM boards 
		WHERE id = $1;
	`

	_, err := s.db.Exec(ctx, query, boardId)
	if err != nil {
		slog.Error("storage DeleteBoard error: " + err.Error())
		return nil, fmt.Errorf("storage DeleteBoard error: %v", err.Error())
	}

	_, err = s.profilesClient.MoveIdeasToBoard(ctx, &profilesv1.MoveIdeaToBoardRequest{
		UserId:     userId,
		OldBoardId: boardId,
		NewBoardId: -1,
	})
	if err != nil {
		slog.Error("storage DeleteBoard error: " + err.Error())
		return nil, fmt.Errorf("storage DeleteBoard error: %v", err.Error())
	}

	_, err = s.profilesClient.RemoveBoardFromProfile(ctx, &profilesv1.RemoveBoardFromProfileRequest{BoardId: boardId, UserId: userId})
	if err != nil {
		slog.Error("storage DeleteBoard error: " + err.Error())
		return nil, fmt.Errorf("storage DeleteBoard error: %v", err.Error())
	}
	return nil, nil
}
