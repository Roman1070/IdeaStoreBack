package auth

import (
	"context"
	"errors"
	"fmt"
	"idea-store-auth/internal/domain/models"
	"idea-store-auth/internal/lib/jwt"
	"idea-store-auth/internal/lib/logger/sl"
	"idea-store-auth/internal/storage"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passwordHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppID       = errors.New("invalid app id")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

// New returns a new instance of the Auth service.
func New(log *slog.Logger, userProvider UserProvider, userSaver UserSaver, tokenTTL time.Duration) *Auth {
	return &Auth{
		userSaver:    userSaver,
		userProvider: userProvider,
		log:          log,
		tokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appID int,
) (string, error) {
	const op = "auth.Login"
	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)
	log.Info("Trying to login user")

	user, err := a.userProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrAppNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		a.log.Error("failed to get user", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}
	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")
	token, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		a.log.Error("Failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (a *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	password string,
) (int64, error) {
	const op = "auth.RegisterNewUser"
	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)
	log.Info("Registering user")
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {

		log.Error("Failed to generate password hash", sl.Err(err))
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	id, err := a.userSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			log.Warn("User already exists")

			return -1, fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		log.Error("Failed to save user", sl.Err(err))
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
