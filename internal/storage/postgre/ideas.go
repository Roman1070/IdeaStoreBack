package postgre

import (
	"context"
	"fmt"
	profilesv1 "idea-store-auth/gen/go/profiles"
	"idea-store-auth/internal/domain/models"
	"log/slog"
	"slices"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateIdea(ctx context.Context, idea models.Idea) (int64, error) {
	slog.Info("storage started CreateIdea")

	const query = `
		INSERT INTO ideas(image,name,description,link,tags,user_id,likes_count) 
		VALUES($1,$2,$3,$4,$5,$6,$7)
		RETURNING id;
	`

	var lastInsertId int64
	err := s.db.QueryRow(ctx, query, idea.Image, idea.Name, idea.Description, idea.Link, idea.Tags, idea.UserID, 0).Scan(&lastInsertId)
	if err != nil {
		slog.Error("storage CreateIdea error: " + err.Error())
		return emptyValue, fmt.Errorf("storage CreateIdea error: %v", err.Error())
	}

	return lastInsertId, nil
}
func (s *Storage) GetIdea(ctx context.Context, id int64) (models.Idea, error) {
	slog.Info("storage started GetIdea")

	const query = `
		SELECT id,image,name,description,link,tags,user_id,likes_count 
		FROM ideas 
		WHERE id = $1;
	`

	var idea models.Idea
	err := s.db.QueryRow(ctx, query, id).Scan(&idea.ID, &idea.Image, &idea.Name, &idea.Description, &idea.Link, &idea.Tags, &idea.UserID, &idea.Likes)
	if err != nil {
		slog.Error("storage GetIdea error: " + err.Error())
		return models.Idea{}, fmt.Errorf("storage GetIdea error: %v", err.Error())
	}
	return idea, nil
}
func (s *Storage) DeleteIdea(ctx context.Context, id int64) (emptypb.Empty, error) {
	slog.Info("storage started to DeleteIdea")

	const query = `
		DELETE FROM ideas 
		WHERE id = $1;
	`

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		slog.Error("storage DeleteIdea error: " + err.Error())
		return emptypb.Empty{}, fmt.Errorf("storage DeleteIdea error: %v", err.Error())
	}

	return emptypb.Empty{}, nil
}
func (s *Storage) ChangeLikesCount(ctx context.Context, ideaId int64, increase bool) (int64, error) {
	slog.Info("storage started to ChangeLikesCount")

	const selectQuery = `
		SELECT likes_count 
		FROM ideas 
		WHERE id = $1;
	`
	const updateQuery = `
		UPDATE ideas 
		SET likes_count = $1 
		WHERE id = $2;
	`

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		slog.Error("storage ChangeLikesCount error: " + err.Error())
		return emptyValue, fmt.Errorf("storage ChangeLikesCount error: " + err.Error())
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var likesCount int64
	err = tx.QueryRow(ctx, selectQuery, ideaId).Scan(&likesCount)
	if err != nil {
		slog.Error("storage ChangeLikesCount error: " + err.Error())
		return emptyValue, fmt.Errorf("storage ChangeLikesCount error: " + err.Error())
	}

	var newLikesCount int64
	if increase {
		newLikesCount = likesCount + 1
	} else {
		newLikesCount = likesCount - 1
	}

	_, err = tx.Exec(ctx, updateQuery, newLikesCount, ideaId)
	if err != nil {
		slog.Error("storage ChangeLikesCount error: " + err.Error())
		return emptyValue, fmt.Errorf("storage ChangeLikesCount error: " + err.Error())
	}

	return newLikesCount, nil
}
func (s *Storage) GetAllIdeas(ctx context.Context, userId int64, limit, offset int32) ([]*models.Idea, error) {
	slog.Info("storage started GetAllIdeas")

	const query = `
		SELECT id,image,name
		FROM ideas
		LIMIT $1 
		OFFSET $2 
		ORDER BY likes_count DESC, id ASC;
	`

	var savedIdsResponse *profilesv1.GetSavedIdeasIdsResponse
	var err error
	if userId != -1 {
		savedIdsResponse, err = s.profilesClient.GetSavedIdeasIds(ctx, &profilesv1.GetSavedIdeasIdsRequest{
			UserId: userId,
		})
	}
	if err != nil {
		slog.Error("storage GetAllIdeas error: " + err.Error())
		return nil, fmt.Errorf("storage GetAllIdeas error: " + err.Error())
	}

	rows, err := s.db.Query(ctx, query, limit, offset)
	if err != nil {
		slog.Error("storage GetAllIdeas error: " + err.Error())
		return nil, fmt.Errorf("storage GetAllIdeas error: " + err.Error())
	}

	defer rows.Close()

	var ideas []*models.Idea
	for rows.Next() {
		idea := new(models.Idea)
		err = rows.Scan(&idea.ID, &idea.Image, &idea.Name)
		if userId != -1 {
			idea.Saved = slices.Contains(savedIdsResponse.IdeasIds, idea.ID)
		}
		if err != nil {
			slog.Error("storage GetAllIdeas error: " + err.Error())
			return nil, fmt.Errorf("storage GetAllIdeas error: " + err.Error())
		}
		ideas = append(ideas, idea)
	}

	return ideas, nil
}
func (s *Storage) GetIdeasFromSearch(ctx context.Context, userId int64, input string) ([]*models.Idea, error) {
	slog.Info("storage started GetIdeasFromSearch")

	const query = `
		SELECT id,image,name
		FROM ideas
		WHERE name LIKE '%' || $1 || '%' OR description LIKE '%' || $1 || '%' OR tags LIKE '%' || $1 || '%'
		ORDER BY id DESC;
	`
	var savedIdsResponse *profilesv1.GetSavedIdeasIdsResponse
	var err error
	if userId != -1 {
		savedIdsResponse, err = s.profilesClient.GetSavedIdeasIds(ctx, &profilesv1.GetSavedIdeasIdsRequest{
			UserId: userId,
		})
	}
	if err != nil {
		slog.Error("storage GetIdeasFromSearch error: " + err.Error())
		return nil, fmt.Errorf("storage GetIdeasFromSearch error: " + err.Error())
	}

	rows, err := s.db.Query(ctx, query, input)
	if err != nil {
		slog.Error("storage GetIdeasFromSearch error: " + err.Error())
		return nil, fmt.Errorf("storage GetIdeasFromSearch error: " + err.Error())
	}

	defer rows.Close()

	var ideas []*models.Idea
	for rows.Next() {
		idea := new(models.Idea)
		err = rows.Scan(&idea.ID, &idea.Image, &idea.Name)
		if userId != -1 {
			idea.Saved = slices.Contains(savedIdsResponse.IdeasIds, idea.ID)
		}
		if err != nil {
			slog.Error("storage GetIdeasFromSearch error: " + err.Error())
			return nil, fmt.Errorf("storage GetIdeasFromSearch error: " + err.Error())
		}
		ideas = append(ideas, idea)
	}

	return ideas, nil
}

func (s *Storage) GetIdeas(ctx context.Context, ids []int64) ([]*models.Idea, error) {
	slog.Info("storage started to GetIdeas")

	if len(ids) == 0 {
		return []*models.Idea{}, nil
	}

	anySlice := make([]any, len(ids))
	for i, v := range ids {
		anySlice[i] = v
	}

	idsRequestString := "("
	i := 1
	for ; i < len(ids); i++ {
		idsRequestString += fmt.Sprintf("$%v,", i)
	}

	idsRequestString += fmt.Sprintf("$%v)", i)
	query := fmt.Sprintf(`
		SELECT id,image,name 
		FROM ideas 
		WHERE id IN %v;`, idsRequestString)

	rows, err := s.db.Query(ctx, query, anySlice...)
	if err != nil {
		slog.Error("storage GetIdeas error: " + err.Error())
		return nil, fmt.Errorf("storage GetIdeas error: " + err.Error())
	}

	var ideas []*models.Idea
	for rows.Next() {
		idea := new(models.Idea)
		err = rows.Scan(&idea.ID, &idea.Image, &idea.Name)
		if err != nil {
			slog.Error("storage GetIdeas error: " + err.Error())
			return nil, fmt.Errorf("storage GetIdeas error: " + err.Error())
		}

		ideas = append(ideas, idea)
	}

	return ideas, nil
}
