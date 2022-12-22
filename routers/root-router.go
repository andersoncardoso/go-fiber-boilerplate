package routers

import (
	"myapp/models"

	"github.com/gofiber/fiber/v2"
)

func RootRouter(app *fiber.App) {
	db := models.GetDB()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"ok": true})
	})

	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		var result bool
		tx := db.Raw("Select 1").Scan(&result)
		if tx.Error != nil {
			return ctx.Status(500).JSON(fiber.Map{"ok": false})
		}

		return ctx.JSON(fiber.Map{"ok": true})
	})
}
