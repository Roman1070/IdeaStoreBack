package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/storage"

	"github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Storage struct {
	db *sql.DB
}

const emptyValue = -1

func New(storagePath string) (*Storage, error) {
	const op = "sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
func (s *Storage) CreateIdea(ctx context.Context, idea models.Idea) (int64, error){
	const op = "storage.sqlite.SaveIdea"
	
	stmt, err := s.db.Prepare("INSERT INTO ideas(image,name,description,link,tags) VALUES(?,?,?,?,?)")
	if err != nil {
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	
	res, err := stmt.ExecContext(ctx, idea.Image, idea.Name, idea.Description, idea.Link, idea.Tags)
	
	if err != nil {
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
func (s *Storage) GetIdea(ctx context.Context, id int64) (models.Idea, error){
	const op = "storage.sqlite.GetIdea"
	
	stmt, err := s.db.Prepare("SELECT id,image,name,description,link,tags FROM ideas WHERE id = ?")
	if err != nil {
		return models.Idea{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, id)
	var idea models.Idea
	err = row.Scan(&idea.ID, &idea.Image, &idea.Name, &idea.Description, &idea.Link,&idea.Tags)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Idea{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}
		return models.Idea{}, fmt.Errorf("%s: %w", op, err)
	}
	return idea, nil
}
func (s *Storage) DeleteIdea(ctx context.Context, id int64)  (emptypb.Empty, error){	
	const op = "storage.sqlite.GetIdea"
	
	stmt, err := s.db.Prepare("DELETE FROM ideas WHERE id = ?")
	stmt.ExecContext(ctx, id)
	if err != nil {
		return emptypb.Empty{}, fmt.Errorf("%s: %w", op, err)
	}
	return emptypb.Empty{},nil
}
func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = "storage.sqlite.SaveUser"
	stmt, err := s.db.Prepare("INSERT INTO users(email, pass_hash) VALUES(?,?)")
	if err != nil {
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.ExecContext(ctx, email, passHash)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return emptyValue, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}
		return emptyValue, fmt.Errorf("%s: %w", op, err)

	}
	id, err := res.LastInsertId()
	if err != nil {
		return emptyValue, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const op = "storage.sqlite.User"
	stmt, err := s.db.Prepare("SELECT id,email,pass_hash FROM users WHERE email = ?")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, email)
	var user models.User
	err = row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}

func (s *Storage) IsAdmin(ctx context.Context, userId int64) (bool, error) {

	const op = "storage.sqlite.User"
	stmt, err := s.db.Prepare("SELECT is_admin FROM users WHERE id = ?")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userId)
	var isAdmin bool
	err = row.Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return isAdmin, nil
}
func (s *Storage) App(ctx context.Context, id int) (models.App, error) {
	const op = "storage.sqlite.App"

	stmt, err := s.db.Prepare("SELECT id, name, secret FROM apps WHERE id = ?")
	if err != nil {
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, id)

	var app models.App
	err = row.Scan(&app.ID, &app.Name, &app.Secret)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}

		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}
