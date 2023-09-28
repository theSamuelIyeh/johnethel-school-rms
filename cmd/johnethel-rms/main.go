package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
	"github.com/thesamueliyeh/johnethel-rms/internal/database"
	"github.com/thesamueliyeh/johnethel-rms/internal/routes"
	"github.com/thesamueliyeh/johnethel-rms/internal/services"
)

func main() {
	// load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// create django template engine
	engine := django.New("./templates", ".html")

	// create fiber app
	app := fiber.New(fiber.Config{
		AppName:           "Johnethel School Result Management System",
		PassLocalsToViews: true,
		Views:             engine,
	})

	app.Static("/", "./static")

	// initialize database connection
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()
	db := database.NewDatabase(database.Engine)

	// init services here
	services.InitServices(db)

	//register routes
	routes.SetupRoutes(app)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)

	// log server start
	log.Println("Server started already...")
}
