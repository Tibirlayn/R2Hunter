package cors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func Cors(appFiber *fiber.App) {
	appFiber.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Разрешите ваш фронтенд-домен
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS", // Разрешить нужные методы
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization", // Разрешить заголовки
		ExposeHeaders:    "Authorization, Refresh-Token", // Экспонирование нужных заголовков
		AllowCredentials: true, // Если отправляете куки или токены
	}))

	// Добавление обработки OPTIONS запросов для всех маршрутов
	// Для теста 
/* 	appFiber.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})  */

	// curl -i -X OPTIONS http://localhost:8087/login
	// curl -i -X POST http://localhost:8087/login -H "Content-Type: application/json" -d "{\"email\": \"test@example.com\", \"password\": \"password123\", \"app_id\": \"2491\"}"
}