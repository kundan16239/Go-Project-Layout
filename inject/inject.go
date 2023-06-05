package inject

import (
	"go-folder-sample/app/controllers"
	"go-folder-sample/app/repositories"
	"go-folder-sample/app/routes"
	"go-folder-sample/app/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type DI struct {
	Ctx *fiber.App
	Db  *mongo.Database
}

func NewDI(Ctx *fiber.App, Db *mongo.Database) *DI {
	return &DI{
		Ctx: Ctx,
		Db:  Db,
	}
}

func (DI *DI) Inject() {
	// User Routes
	userRepo := repositories.NewUserRepository(DI.Db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController)
	userRoute.RegisterRoutes(DI.Ctx)

}
