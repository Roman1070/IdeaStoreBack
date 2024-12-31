package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/storage"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateIdea(ctx context.Context, idea models.Idea) (int64, error) {
	const op = "storage.sqlite.SaveIdea"

	stmt, err := s.db.Prepare("INSERT INTO ideas(image,name,description,link,tags,user_id) VALUES(?,?,?,?,?,?)")
	if err != nil {
		slog.Error("CreateIdea storage Prepare error: " + err.Error())
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, idea.Image, idea.Name, idea.Description, idea.Link, idea.Tags, idea.UserID)

	if err != nil {
		slog.Error("CreateIdea storage ExecContext error: " + err.Error())
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		slog.Error("CreateIdea storage LastInsertId error: " + err.Error())
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
func (s *Storage) GetIdea(ctx context.Context, id int64) (models.Idea, error) {
	const op = "storage.sqlite.GetIdea"

	stmt, err := s.db.Prepare("SELECT id,image,name,description,link,tags,user_id,likes_count FROM ideas WHERE id = ?")
	if err != nil {
		return models.Idea{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, id)
	var idea models.Idea
	err = row.Scan(&idea.ID, &idea.Image, &idea.Name, &idea.Description, &idea.Link, &idea.Tags, &idea.UserID, &idea.Likes)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Idea{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}
		return models.Idea{}, fmt.Errorf("%s: %w", op, err)
	}
	return idea, nil
}
func (s *Storage) DeleteIdea(ctx context.Context, id int64) (emptypb.Empty, error) {
	const op = "storage.sqlite.GetIdea"

	stmt, err := s.db.Prepare("DELETE FROM ideas WHERE id = ?")
	stmt.ExecContext(ctx, id)
	if err != nil {
		return emptypb.Empty{}, fmt.Errorf("%s: %w", op, err)
	}
	return emptypb.Empty{}, nil
}
func (s *Storage) GetAllIdeas(ctx context.Context, _ *emptypb.Empty) ([]*ideasv1.IdeaData, error) {
	const op = "storage.sqlite.GetIdea"

	stmt, err := s.db.Prepare("SELECT id,image,name,description,link,tags FROM ideas")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ideas := []*ideasv1.IdeaData{}
	for rows.Next() {
		idea := new(ideasv1.IdeaData)
		err = rows.Scan(&idea.IdeaId, &idea.Image, &idea.Name, &idea.Description, &idea.Link, &idea.Tags)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		ideas = append(ideas, idea)
	}

	return ideas, nil
}
