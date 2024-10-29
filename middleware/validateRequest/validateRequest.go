package validaterequest

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/Tibirlayn/R2Hunter/pkg/jwt"
)

func validateRequest(ctx *fiber.Ctx) error {
	// Пример проверки токена или другого параметра запроса

/* 	app, err := a.appProvider.App(ctx, 2491)
	if err != nil {
		a.log.Error(fmt.Sprintf("%s: %v", op, err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	// валидация токена
	userID, err := jwt.ValidateToken(ctx, app.Secret)
	if err != nil {
		a.log.Error(fmt.Sprintf("Token validation error: %v", err))
		return 0, fmt.Errorf("%s: %w", op, err)
	} */

	// Дополнительные проверки можно добавить здесь
	return ctx.Next() // Переход к следующему обработчику
}