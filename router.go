package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thesamueliyeh/johnethel-rms/controllers"
)

func setupRouter(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", controllers.IndexHandler)
}
