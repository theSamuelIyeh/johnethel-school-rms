package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func setup() {
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		AppName:           "Johnethel School Result Management System",
		PassLocalsToViews: true,
		Views:             engine,
	})
	app.Static("/", "./public")

	setupRouter(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
