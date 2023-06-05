package main

import (
	"go-folder-sample/configs"
	"go-folder-sample/database"
	"go-folder-sample/inject"
	"go-folder-sample/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	logger.InitLogger()
	configs.InitialiseConfig()
	db := database.ConnectMongoDB()

	DI := inject.NewDI(app, db)
	DI.Inject()

	app.Listen(":3000")
}
