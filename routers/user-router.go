package routers

import (
	"myapp/models"
	"myapp/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	app.Get("/users", func(ctx *fiber.Ctx) error {
		users, err := services.UserList()
		if err != nil {
			return ErrResponse(ctx, 500, "Can't get all users. "+err.Error())
		}
		return ctx.JSON(users)
	})

	// REMOVE-ME
	// Dummy endpoint for testin
	app.Get("/users/_test", func(ctx *fiber.Ctx) error {
		user := &models.User{
			FirstName: "John3",
			LastName:  "Doe",
			Email:     "john3@example.com",
			Password:  "12345",
		}
		err := services.UserCreate(user)
		if err != nil {
			return ErrResponse(ctx, 422, "Could not create user. "+err.Error())
		}
		return ctx.JSON(user)
	})

	app.Get("/users/:id", func(ctx *fiber.Ctx) error {
		id, err := IdParam(ctx)
		if err != nil {
			return err
		}

		user, err := services.UserFind(id)
		if err != nil {
			return ErrResponse(ctx, 404, err.Error())
		}
		return ctx.JSON(user)
	})

	app.Post("/users", func(ctx *fiber.Ctx) error {
		var user models.User

		err := ctx.BodyParser(user)
		if err != nil {
			return ErrResponse(ctx, 422, "Could not parse JSON. "+err.Error())
		}

		err = services.UserCreate(&user)
		if err != nil {
			return ErrResponse(ctx, 422, "Could not create user. "+err.Error())
		}
		return ctx.Status(201).JSON(user)
	})

	app.Post("/users/:id", func(ctx *fiber.Ctx) error {
		id, err := IdParam(ctx)
		if err != nil {
			return err
		}

		// First we ensure the User exists
		user, err := services.UserFind(id)
		if err != nil {
			return ErrResponse(ctx, 404, err.Error())
		}

		err = ctx.BodyParser(user)
		if err != nil {
			return ErrResponse(ctx, 422, "Could not parse JSON. "+err.Error())
		}

		if strconv.FormatUint(user.ID, 10) != ctx.Params("id") {
			return ErrResponse(ctx, 400, "IDs mismatch ")
		}

		err = services.UserUpdate(&user)
		if err != nil {
			return ErrResponse(ctx, 422, "Could not update. "+err.Error())
		}
		return ctx.JSON(user)
	})

	app.Delete("/users/:id", func(ctx *fiber.Ctx) error {
		id, err := IdParam(ctx)
		if err != nil {
			return err
		}
		err = services.UserDelete(id)
		if err != nil {
			return ErrResponse(ctx, 422, "Could not delete user. "+err.Error())
		}

		return ctx.JSON(fiber.Map{"message": "user deleted"})
	})
}
