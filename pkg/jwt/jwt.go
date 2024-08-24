package jwt

import (
	"time"

	"github.com/Tibirlayn/R2Hunter/internal/domain/models"
	"github.com/Tibirlayn/R2Hunter/internal/domain/models/account"
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

/*func ValidateToken(ctx *fiber.Ctx, secret string) (int64, error) {
	// Получение метаданных с токеном доступа из контекста запроса
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, status.Errorf(codes.InvalidArgument, "Metadata not provided")
	}

	// Проверка токена доступа
	jwtToken := md.Get("authorization")
	if len(jwtToken) < 1 {
		return 0, status.Errorf(codes.Unauthenticated, "Authorization token is required")
	}

	tokenString := strings.TrimPrefix(jwtToken[0], "Bearer ") // Удаление "Bearer " из jwtToken

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Предоставьте секретный ключ для верификации токена
		return []byte(secret), nil // secret - Предоставьте секретный ключ для вашего токена
	})

	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	var userID int64
	// Проверка времени истечения токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp := time.Unix(int64(claims["exp"].(float64)), 0)
		userID = int64(claims["uid"].(float64))
		if exp.Before(time.Now()) {
			return 0, fmt.Errorf("token has expired")
		}
	}
	// Валидация успешна
	return userID, nil
}*/
