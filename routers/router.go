package routers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{"error": "Not Found"})
}

func ErrResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.
		Status(status).
		JSON(fiber.Map{"error": message})
}

func IdParam(ctx *fiber.Ctx) (uint64, error) {
	ctx.Accepts("application/json")
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return id, ErrResponse(ctx, 422, "Error: could not parse id. "+err.Error())
	}
	return id, nil
}
