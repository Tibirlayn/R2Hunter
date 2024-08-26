package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// NewToken creates new JWT token for given user and app.
func NewToken(user account.Member, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256) // Создает новый токен типа jwt.Token с указанием метода подписи HS256

	claims := token.Claims.(jwt.MapClaims)          // Получает доступ к заявкам (claims) токена как к словарю с помощью jwt.MapClaims, чтобы можно было установить значения для каждой заявки
	claims["uid"] = user.MUserId                    // Устанавливает уникальный идентификатор пользователя в заявке "uid".
	claims["nphone"] = user.Email                   // Устанавливает номер телефона пользователя в заявке "nphone".
	claims["exp"] = time.Now().Add(duration).Unix() // Устанавливает время истечения токена в заявке "exp".
	claims["app_id"] = app.ID                       // Устанавливает идентификатор приложения (app) в заявке "app_id".

	tokenString, err := token.SignedString([]byte(app.Secret)) // Подписывает токен, используя секрет приложения (app.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(ctx *fiber.Ctx, secret string) (int64, error) {
    // Получение токена доступа из заголовка Authorization
    jwtToken := ctx.Get("Authorization")
    if jwtToken == "" {
        return 0, fmt.Errorf("authorization token is required")
    }

    tokenString := strings.TrimPrefix(jwtToken, "Bearer ")

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        return 0, fmt.Errorf("failed to parse token: %v", err)
    }

    if !token.Valid {
        return 0, fmt.Errorf("invalid token")
    }

    var userID int64

    // Проверка времени действия токена
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        exp := time.Unix(int64(claims["exp"].(float64)), 0)
        userID = int64(claims["uid"].(float64))

        if exp.Before(time.Now()) {
            return 0, fmt.Errorf("token has expired")
        }
    }

    // Все проверки успешны - возврат userID
    return userID, nil
}
