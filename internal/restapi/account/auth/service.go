package auth

import (
	"fmt"
	"strconv"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	gen "github.com/Tibirlayn/R2Hunter/pkg/lib/genlogin"
	routersAuth "github.com/Tibirlayn/R2Hunter/internal/routers/account/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

type LoginRequestValid struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type RegisterRequestValid struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func init() {
	validate = validator.New()
}

type Auth interface {
	Login(ctx *fiber.Ctx, user account.Member, appID int) (token string, err error)
	RegisterNewUser(ctx *fiber.Ctx, user account.Member, appID int) (userID int64, err error)
}

type ServerAPI struct {
	auth Auth
}

func Register(RestAPI *fiber.App, auth Auth) {
	api := &ServerAPI{auth: auth}

	// Передача api как реализации интерфейса AuthHandler
	routersAuth.NewRoutersAuth(RestAPI, api)
}

func (s *ServerAPI) Login(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.account.auth.service.Login"
		appId = "Unknown AppID"
	)

	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	appID, err := strconv.Atoi(data["app_id"])
	if err != nil {
		return fmt.Errorf("%s: %s", op, appId)
	}

	user := account.Member{
		Email: data["email"],
		MUserPswd: data["password"],
	}

	ValidLogin := LoginRequestValid{
		Email:    user.Email,
		Password: user.MUserPswd,
	}

	if err := validate.Struct(ValidLogin); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if appID == 0 {
		return fmt.Errorf("%s: %s", op, appId)
	}

	token, err := s.auth.Login(ctx, user, appID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return ctx.JSON(token)
}

func (s *ServerAPI) Register(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.account.auth.service.Login"
		appId = "Unknown AppID"
	)

	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}
	
	appID, err := strconv.Atoi(data["app_id"])
	if err != nil {
		return fmt.Errorf("%s: %s", op, appId)
	}

	login := gen.RemoveEmailSymbols(data["email"])

	user := account.Member{
		MUserId: login,
		Email: data["email"],
		MUserPswd: data["password"],
	}

	ValidReg := RegisterRequestValid{
		Email:    user.Email,
		Password: user.MUserPswd,
	}
	if err := validate.Struct(ValidReg); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.auth.RegisterNewUser(ctx, user, appID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return ctx.JSON(id)
}
