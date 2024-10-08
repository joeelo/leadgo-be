// https://stackoverflow.com/questions/76175950/connect-to-mongodb-when-you-have-google-login-as-credentials-for-mongodb-atlas

package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	//run database
	configs.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	routes.UserRoute(app)
	routes.PrgoramRoute(app)

	app.Listen(":6000")
}
