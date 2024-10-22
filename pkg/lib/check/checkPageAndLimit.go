package check

import "github.com/gofiber/fiber/v2"

func CheckPageAndLimit(ctx *fiber.Ctx) (page int, limit int, err error) {
	const op = "pkg.lib.check.CheckPageAndLimit"

	pageNumber := ctx.QueryInt("page")
	if pageNumber == 0 {
		return 0, 0, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message:": "Invalid page number " + op,
		})
	}
	limitCnt := ctx.QueryInt("limit")
	if limitCnt == 0 {
		return 0, 0, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid limitCnt " + op,
		})
	}

	return pageNumber, limitCnt, nil
}