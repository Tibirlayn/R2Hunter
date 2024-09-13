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
//	Member(ctx *fiber.Ctx, mp query.MemberParm) (query.MemberParm, error)
	Member(ctx *fiber.Ctx, name string) (query.MemberParm, error)
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

func (m *Member) Member(ctx *fiber.Ctx, name string) (query.MemberParm, error) {
		const op = "service.account.member.Member"

		var memberParm query.MemberParm

		// проверка на авторизацию 
		userID, err := m.auth.ValidJWT(ctx, op)
		if err != nil {
			return memberParm, err
		}

		m.log.Info(fmt.Sprintf("admin %d checking user data", userID))

		resultMember, err := m.usrMemberProvider.Member(ctx, name)
		if err != nil {
			m.log.Info("%s, %w", op, err)
			// errorList = append(errorList, err)
			return resultMember, err // Вернуть ошибку, если не удалось получить данные
		}


		return resultMember, nil
}

/* 		memberParm = resultMember

		// Получаем срез pcCard, содержащий несколько PcParm
		pcCards, err := m.pc.PcCard(ctx, name, int64(memberParm.User.MUserNo))
		if err != nil {
			m.log.Info("%s, %w", op, err)
			errorList = append(errorList, err)
			return memberParm, fmt.Errorf("%s, %w", op, err)
		}

		// Присваиваем срез pcCards к полю PcCards структуры memberParm
		memberParm.PcCards = pcCards */


/* 			
			pcCard, err := m.pc.PcCard(ctx, name, int64(resultMember.User.MUserNo))
			if err != nil {
				return memberParm, fmt.Errorf("%s, %w", op, err)
			} 
*/

/*  			
			resultMember.Pc = pcCard.Pc
			resultMember.PcInv = pcCard.PcInv
			resultMember.PcState = pcCard.PcState 
*/