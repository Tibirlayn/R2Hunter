package member

import (
	"fmt"
	"log/slog"
	"time"

	account "github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/query/account"
	"github.com/Tibirlayn/R2Hunter/internal/service/parm"
	"github.com/Tibirlayn/R2Hunter/internal/service/account/auth"
	"github.com/Tibirlayn/R2Hunter/internal/service/game"
	"github.com/gofiber/fiber/v2"
)

type Member struct {
	log *slog.Logger
	usrMemberProvider UserMemberProvider
	pc *pc.Pc
	parm *parm.Parm
	auth *auth.Auth
	tokenTTL time.Duration
}

type UserMemberProvider interface {
	Member(ctx *fiber.Ctx, name string) (query.MemberParm, error)
	MemberAll(ctx *fiber.Ctx, name string) (query.MemberPcItem, error)
	MemberBil(ctx *fiber.Ctx, email string) (account.User, error)
	UserSearch(ctx *fiber.Ctx, name string) ([]account.User, error)
	UserLastLogin(ctx *fiber.Ctx) ([]int, error)
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

func (m *Member) UserSearch(ctx *fiber.Ctx, name string) ([]account.User, error) {
	const op = "service.account.member.User"

	// проверка на авторизацию 
	_, err := m.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, err
	}

	res, err := m.usrMemberProvider.UserSearch(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil	
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

func (m *Member) MemberAll(ctx *fiber.Ctx, name string) (query.MemberPcItem, error) {
	const op = "service.account.member.Member"

	// проверка на авторизацию 
	userID, err := m.auth.ValidJWT(ctx, op)
	if err != nil {
		return query.MemberPcItem{}, err
	}

	m.log.Info(fmt.Sprintf("admin %d checking user data", userID))

	resMember, err := m.usrMemberProvider.MemberAll(ctx, name)
	if err != nil {
		return query.MemberPcItem{}, err // Вернуть ошибку, если не удалось получить данные
	}

	return resMember, nil
}

func (m *Member) MemberBil(ctx *fiber.Ctx, email string) (account.User, error) {
	const op = "service.account.member.MemberBil"

	// проверка на авторизацию 
	_, err := m.auth.ValidJWT(ctx, op)
	if err != nil {
		return account.User{}, fmt.Errorf("%s, %w", op, err)
	}

	res, err := m.usrMemberProvider.MemberBil(ctx, email)
	if err != nil {
		return account.User{}, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}

func (m *Member) UserLastLogin(ctx *fiber.Ctx) ([]int, error) {
	const op = "service.account.member.MemberBil"

	// проверка на авторизацию 
	_, err := m.auth.ValidJWT(ctx, op)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	res, err := m.usrMemberProvider.UserLastLogin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}