package routes

import (
	controllers "fiber-mongo-api/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/user/:userId", controllers.GetAUser)
	app.Post("/user", controllers.CreateUser)
	app.Get("/users", controllers.GetAllUsers)
}
