package main

import (
	"myapp/configs"
	"myapp/models"
	"myapp/routers"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	flag.Parse()

	configs.Load()
	models.GetDB()

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	// Bind routers/handlers
	routers.RootRouter(app)
	routers.UserRouter(app)

	// Handle not founds
	app.Use(routers.NotFound)

	//
	log.Fatal(app.Listen(*port))
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(logger.New())

	router.Listen(":3000")
}
