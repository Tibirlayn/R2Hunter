package member

import (
	"fmt"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/game"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query/account"
	routersMember "github.com/Tibirlayn/R2Hunter/internal/routers/account/member"
//	gen "github.com/Tibirlayn/R2Hunter/pkg/lib/genlogin"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Member interface {
	// name = email, login, nikname
	//Member(ctx *fiber.Ctx, mp query.MemberParm) (memberParm query.MemberParm, err error)
	Member(ctx *fiber.Ctx, name string) (memberParm query.MemberParm, err error)
}

type ServiceMemberAPI struct {
	member     Member
}

func RegisterMember(RestAPI *fiber.App, member Member) {
	api := &ServiceMemberAPI{member: member}

	routersMember.NewRoutersMember(RestAPI, api)
}

func (s *ServiceMemberAPI) Member(ctx *fiber.Ctx) error {
	const (
		op    = "restapi.account.member.Member"
		empty = "empty"
	)

	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	if data["name"] == "" {
		return fmt.Errorf("%s, %s", op, "empty")
	}

	// TODO: проверить на валидацию логин или никнейм
	// login := gen.RemoveEmailSymbols(data["name"])

	validMember := query.MemberParm{ 
		Member: account.Member{ 
			Email:   data["name"],
			MUserId: data["name"], // login
		},
		Pc: game.Pc{ 
			MNm: data["name"],
		},
	}
	
	if err := validate.Struct(validMember); err != nil {
		return fmt.Errorf("%s, %w", op, err)
	}

	resultMemberParm, err := s.member.Member(ctx, data["name"])
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return ctx.JSON(resultMemberParm)
}