package member

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query/account"
	gen "github.com/Tibirlayn/R2Hunter/pkg/lib/genlogin"
	routersMember "github.com/Tibirlayn/R2Hunter/internal/routers/account/member"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Member interface {
	Member(ctx *fiber.Ctx, name string) (memberParm query.MemberParm, err error)
	MemberAll(ctx *fiber.Ctx, name string) (memberPcItem query.MemberPcItem, err error)
	UserSearch(ctx *fiber.Ctx, name string) ([]account.User, error)
}

type ServiceMemberAPI struct {
	member     Member
}

func RegisterMember(RestAPI *fiber.App, member Member) {
	api := &ServiceMemberAPI{member: member}

	routersMember.NewRoutersMember(RestAPI, api)
}

func (s *ServiceMemberAPI) UserSearch(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.account.member.User"
		empty = "empty"
	)

	name := ctx.Query("name")

	if name == "" {
		return fmt.Errorf("%s, %s", op, empty)
	}

	login := gen.RemoveEmailSymbols(name)

	res, err := s.member.UserSearch(ctx, login)
	if err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	return ctx.JSON(res)
}

func (s *ServiceMemberAPI) Member(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.account.member.Member"
		empty = "empty"
	)

	name := ctx.Query("name")

	if name == "" {
		return fmt.Errorf("%s, %s", op, "empty")
	}

 	validMember := account.Member{
		Email: name,
		MUserId: name,
	}
	
	if err := validate.Struct(validMember); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	resMemberParm, err := s.member.Member(ctx, name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return ctx.JSON(resMemberParm)
}

func (s *ServiceMemberAPI) MemberAll(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.account.member.Member"
		empty = "empty"
	)

	name := ctx.Query("name")

	if name == "" {
		return fmt.Errorf("%s, %s", op, "empty")
	}

 	validMember := account.Member{
		Email: name,
		MUserId: name,
	}
	
	if err := validate.Struct(validMember); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	resMember, err := s.member.MemberAll(ctx, name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return ctx.JSON(resMember)
}