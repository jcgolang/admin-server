package main

import (
	db "admin-server/database"
	"admin-server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	defer db.GetManager().Disconnect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":3001")
}
