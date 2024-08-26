package member

import (
	"debug/macho"
	"fmt"
	"strings"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	routersMember "github.com/Tibirlayn/R2Hunter/internal/routers/member"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Member interface {
	// name = email, login, nikname
	Member(ctx *fiber.Ctx, ) (
			member account.Member, 
			user account.User, 
			userAdmin account.UserAdmin, 
			pc game.Pc, 
			pcInv game.PcInventory, 
			pcState game.PcState, err error)
}

type ServiceMemberAPI struct {
	member Member
}

func RegisterMember(RestAPI *fiber.App, member Member) {
	api := &ServiceMemberAPI{member: member}

	routersMember.NewRoutersMember(RestAPI, api)
}

func (s *ServiceMemberAPI) Member(c *fiber.Ctx) error {
	const (
		op = "restapi.account.member.Member"
		empty = "empty"
	)

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}
	
	if data["name"] == "" {
		return fmt.Errorf("%s, %s", op, "empty")
	}

	// TODO: проверить на валидацию логин или никнейм
	login := removeEmailSymbols(data["name"])
	validMember := account.Member{
		Email: data["name"],
		MUserId: login,
	}

	validNikname := game.Pc{
		MNm: data["name"],
	}

	// TODO: проверка на авторизацию 


	return nil
}

func removeEmailSymbols(email string) string {
	var result strings.Builder

	for _, char := range email {
		if char != '@' && char != '.' {
			result.WriteRune(char)
		}
	}

	return result.String()
}