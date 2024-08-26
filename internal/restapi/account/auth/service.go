package auth

import (
	"fmt"
	"strconv"

	"github.com/Tibirlayn/R2Hunter/internal/routers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	email string 
	password string 
	app_id int
}

var validate *validator.Validate

type LoginRequestValid struct {
	Email   string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type RegisterRequestValid struct {
	Email   string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func init() {
	validate = validator.New()
}

type Auth interface {
	Login(ctx *fiber.Ctx, email string, password string, appID int) (token string, err error)
	RegisterNewUser(ctx *fiber.Ctx, email string, password string) (userID int64, err error)
}

type ServerAPI struct {
	auth Auth
}

func Register(RestAPI *fiber.App, auth Auth) {
	api := &ServerAPI{auth: auth}

	// Передача api как реализации интерфейса AuthHandler
	routers.NewRoutersAuth(RestAPI, api) 
}

func (s *ServerAPI) Login(c *fiber.Ctx) error {
	const (
		op = "restapi.account.auth.service.Login"
		appId = "Unknown AppID"
	)

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	appID, err := strconv.Atoi(data["app_id"])
	if err != nil {
		return fmt.Errorf("%s: %s", op, appId)
	}

	user := User{
		email: data["email"],
		password: data["password"],
		app_id: appID,
	}

	ValidLogin := LoginRequestValid{
		Email: user.email,
		Password: user.password,
	}
	if err := validate.Struct(ValidLogin); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if user.app_id == 0 {
		return fmt.Errorf("%s: %s", op, appId)
	}

	token, err := s.auth.Login(c, user.email, user.password, user.app_id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.JSON(token)
}

func (s *ServerAPI) Register(c *fiber.Ctx) error {
	const op = "restapi.account.auth.service.Register"

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := User{
		email: data["email"],
		password: data["password"],
	}

	ValidReg := RegisterRequestValid{
		Email: user.email,
		Password: user.password,
	}
	if err := validate.Struct(ValidReg); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.auth.RegisterNewUser(c, user.email, user.password)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return c.JSON(id)
}

