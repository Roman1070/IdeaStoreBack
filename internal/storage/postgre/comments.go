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

	const query = `
		INSERT INTO comments(idea_id,user_id,content,creation_date) 
		VALUES($1,$2,$3,$4);
	`

	_, err := s.db.Exec(ctx, query, ideaId, userId, text, creationDate)
	if err != nil {
		slog.Error("storage CreateComment error: " + err.Error())
		return nil, fmt.Errorf("storage CreateComment error: %v", err.Error())
	}

	return nil, nil
}
func (s *Storage) GetComments(ctx context.Context, ideaId int64) ([]*models.Comment, error) {
	slog.Info("storage started GetComments")

	const query = `
		SELECT id,user_id,content,creation_date 
		FROM comments 
		WHERE idea_id = $1;
	`
	rows, err := s.db.Query(ctx, query, ideaId)
	if err != nil {
		slog.Error("storage GetComments error: " + err.Error())
		return nil, fmt.Errorf("storage GetComments error: %v", err.Error())
	}

	defer rows.Close()

	var result []*models.Comment
	for rows.Next() {
		comment := new(models.Comment)
		err = rows.Scan(&comment.ID, &comment.UserId, &comment.Text, &comment.CreationDate)
		if err != nil {
			slog.Error("storage GetComments error: " + err.Error())
			return nil, fmt.Errorf("storage GetComments error: %v", err.Error())
		}

		profile, err := s.profilesClient.GetProfile(ctx, &profilesv1.GetProfileRequest{
			Id: comment.UserId,
		})
		if err != nil {
			slog.Error("storage GetComments error: " + err.Error())
			return nil, fmt.Errorf("storage GetComments error: %v", err.Error())
		}

		comment.Avatar = profile.Data.AvatarImage
		comment.Username = profile.Data.Name
		result = append(result, comment)
	}

	return result, nil
}
