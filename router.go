package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thesamueliyeh/johnethel-rms/controllers"
)

func setupRouter(app *fiber.App) {
	api := app.Group("/api")
	web := app.Group("/")

	web.Get("/", controllers.IndexController)
	web.Get("/about", controllers.AboutController)
	web.Get("/signup", controllers.SignupController)
	web.Get("/login", controllers.LoginController)
	api.Get("/", controllers.IndexApiController)
}
