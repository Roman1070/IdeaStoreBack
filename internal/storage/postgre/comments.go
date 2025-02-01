package postgre

import (
	"context"
	"fmt"
	profilesv1 "idea-store-auth/gen/go/profiles"
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
func (s *Storage) GetComments(ctx context.Context, ideaId int64) ([]*models.Comment, error) {
	slog.Info("storage started GetComments")

	stmt, err := s.db.Prepare("SELECT id,user_id,content,creation_date from comments WHERE idea_id = ?")
	if err != nil {
		return nil, fmt.Errorf("storage GetComments error: %v", err.Error())
	}
	rows, err := stmt.QueryContext(ctx, ideaId)

	if err != nil {
		return nil, fmt.Errorf("storage GetComments error: %v", err.Error())
	}
	var result []*models.Comment
	for rows.Next() {
		comment := new(models.Comment)
		err = rows.Scan(&comment.ID, &comment.UserId, &comment.Text, &comment.CreationDate)
		if err != nil {
			return nil, fmt.Errorf("storage GetComments error: %v", err.Error())
		}
		profile, err := s.profilesClient.GetProfile(ctx, &profilesv1.GetProfileRequest{
			Id: comment.UserId,
		})
		if err != nil {
			return nil, fmt.Errorf("storage GetComments error: %v", err.Error())
		}
		comment.Avatar = profile.Data.AvatarImage
		comment.Username = profile.Data.Name
		result = append(result, comment)
	}

	return result, nil
}
