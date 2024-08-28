package member

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query/account"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/auth"
	"github.com/Tibirlayn/R2Hunter/internal/service/game/pc"
	"github.com/gofiber/fiber/v2"
)

type Member struct {
	log *slog.Logger
	usrMemberProvider UserMemberProvider
	pc *pc.Pc
	auth *auth.Auth
	tokenTTL time.Duration
}

type UserMemberProvider interface {
	Member(ctx *fiber.Ctx, mp query.MemberParm) (query.MemberParm, error)
}

func New(log *slog.Logger, userMemberProvider UserMemberProvider, auth *auth.Auth, pc *pc.Pc, tokenTTL time.Duration) *Member {
	return &Member{
		log: log,
		usrMemberProvider: userMemberProvider,
		pc: pc,
		auth: auth,
		tokenTTL: tokenTTL,
	}
}


func (m *Member) Member(ctx *fiber.Ctx, mp query.MemberParm) (memberParm query.MemberParm, err error) {
		const op = "service.account.member.Member"

		// TODO: проверка на авторизацию 
		userID, err := m.auth.ValidJWT(ctx, op)
		if err != nil {
			return mp, err
		}

		m.log.Info(fmt.Sprintf("admin %d checking user data", userID))

		// посмотреть по имени персонажа 
		// Если по персонажу есть данные получить все данные по account

		

		result, err := m.usrMemberProvider.Member(ctx, mp)
		if err != nil {
			return mp, fmt.Errorf("%s, %w", op, err)
		}

		return result, nil
}
