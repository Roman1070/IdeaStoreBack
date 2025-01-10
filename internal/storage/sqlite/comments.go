package sqlite

import (
	"context"
	"fmt"
	"idea-store-auth/internal/domain/models"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateComment(ctx context.Context, ideaId, userId int64, text, creationDate string) (*emptypb.Empty, error) {
	slog.Info("storage started CreateComment")

	stmt, err := s.db.Prepare("INSERT INTO comments (idea_id,user_id,content,creation_date) VALUES(?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("storage CreateComment error: %v", err.Error())
	}
	_, err = stmt.ExecContext(ctx, ideaId, userId, text, creationDate)
	if err != nil {
		return nil, fmt.Errorf("storage CreateComment error: %v", err.Error())
	}
	return nil, nil
}
func (s *Storage) GetComments(ctx context.Context, ideId int64) ([]*models.Comment, error) {
	return nil, nil
}
