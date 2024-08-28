package auth

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Tibirlayn/R2Hunter/pkg/jwt"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/storage"
	"github.com/gofiber/fiber/v2"
	// "golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

// Регистрация
type UserSaver interface {
	SaveUser(ctx *fiber.Ctx, user account.Member, appID int) (uid int64, err error) 
}

// Авторизация
type UserProvider interface {
	User(ctx *fiber.Ctx, email string) (account.Member, account.User, error) 
}

type AppProvider interface {
	App(ctx *fiber.Ctx, appID int) (models.App, error) // App
}

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration) *Auth {
	return &Auth{
		log:         log,
		usrSaver:    userSaver,
		usrProvider: userProvider,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

func (a *Auth) Login(ctx *fiber.Ctx, user account.Member, appID int) (string, error) {
	const op = "service.account.auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", user.Email),
	)

	log.Info("attempting to login user")

	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		a.log.Error(fmt.Sprintf("%s: %v", op, err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	resMember, resUser, err := a.usrProvider.User(ctx, user.Email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", "error", err)
			return "", fmt.Errorf("%s: %w", op, storage.ErrInvalidCredentials)
		}
		a.log.Error("failed to get user", "error", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if resMember.MUserPswd != user.MUserPswd {
		a.log.Info("invalid credentials", "error", err)
		return "", fmt.Errorf("%s: %w", op, storage.ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	token, err := jwt.NewToken(resUser, app, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", "error", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) RegisterNewUser(ctx *fiber.Ctx, user account.Member, appID int) (int64, error) {

	const op = "Auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", user.Email),
	)

	log.Info("registering user")

	app, err := a.appProvider.App(ctx, appID)
	if err != nil || app.ID != appID {
		a.log.Error(fmt.Sprintf("%s: %v", op, err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.usrSaver.SaveUser(ctx, user, appID)
	if err != nil {
		log.Error("failed to save user", "error", err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (a *Auth) ValidJWT(ctx *fiber.Ctx, op string) (int64, error) {
	// получаем секрет
	// TODO: должен брать из файла config 
	app, err := a.appProvider.App(ctx, 2491)
	if err != nil {
		a.log.Error(fmt.Sprintf("%s: %v", op, err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	// валидация токена
	userID, err := jwt.ValidateToken(ctx, app.Secret)
	if err != nil {
		a.log.Error(fmt.Sprintf("Token validation error: %v", err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return userID, nil
}
