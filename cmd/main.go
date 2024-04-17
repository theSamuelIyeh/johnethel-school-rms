package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/auth"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/database"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/handlers"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/middlewares"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/routes"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/utils"
)

func main() {
	// load env variables
	godotenv.Load()
	dbUrl := os.Getenv("SUPABASE_DB_URL")
	if dbUrl == "" {
		log.Fatal("no database url found")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("no port found")
	}

	// initialise Supabase auth
	auth.InitSupabaseAuth()

	// initilise db service
	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// create new apiconfig
	apiCfg := &utils.ApiCfg{
		DB: database.New(conn),
	}

	// Create a new instance of Handler struct with the config
	Handler := handlers.NewHandler(apiCfg)

	// Create a new instance of Middleware struct with config
	Middleware := middlewares.NewMiddleware(apiCfg)

	// initialise echo
	e := echo.New()
	e.Static("/static", "static")

	// initialise validator
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// initialise router
	routes.InitRouter(e, Handler, Middleware)

	e.Logger.Fatal(e.Start(":" + port))
}
