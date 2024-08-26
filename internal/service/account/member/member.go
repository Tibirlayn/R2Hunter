package member

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/auth"
	"github.com/gofiber/fiber/v2"
)

type Member struct {
	log *slog.Logger
	usrMemberProvider UserMemberProvider
	auth *auth.Auth
	tokenTTL time.Duration
}

type UserMemberProvider interface {
	Member(ctx *fiber.Ctx, mp query.MemberParm) (query.MemberParm, error)
}

func New(log *slog.Logger, userMemberProvider UserMemberProvider, auth *auth.Auth, tokenTTL time.Duration) *Member {
	return &Member{
		log: log,
		usrMemberProvider: userMemberProvider,
		auth: auth,
		tokenTTL: tokenTTL,
	}
}


func (m *Member) Member(ctx *fiber.Ctx, mp query.MemberParm) (memberParm query.MemberParm, err error) {
		const op = "service.account.member.Member"

		// TODO: проверка на авторизацию 
		userID, err := m.auth.ValidJWT(ctx, op)
		if err != nil {
			return memberParm, err
		}

		m.log.Info(fmt.Sprintf("admin %d checking user data", userID))

		result, err := m.Member(ctx, mp)
		if err != nil {
			return memberParm, fmt.Errorf("%s, %w", op, err)
		}

		return result, nil
}
