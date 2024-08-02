package routes

import (
	controllers "fiber-mongo-api/controllers/program"

	"github.com/gofiber/fiber/v2"
)

func PrgoramRoute(app *fiber.App) {
	app.Get("/program", controllers.GetAProgram)
	app.Post("/program", controllers.CreateProgram)
	app.Get("/programs", controllers.GetAllPrograms)
}
