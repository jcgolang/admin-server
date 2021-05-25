package routes

import (
	authController "admin-server/controllers/auth.controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/auth/register", authController.Register)
}
