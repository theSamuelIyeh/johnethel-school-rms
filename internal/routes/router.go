package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thesamueliyeh/johnethel-rms/internal/handlers/apihandlers"
	"github.com/thesamueliyeh/johnethel-rms/internal/handlers/webhandlers"
	middlewares "github.com/thesamueliyeh/johnethel-rms/internal/middleware"
)

func SetupRoutes(app *fiber.App) {

	// ---------------- group api routes
	api := app.Group("/api")
	api.Get("/", apihandlers.HandleIndexApi)

	// group auth routes for api
	apiAuth := api.Group("/auth", middlewares.IsNotAuthenticated)
	apiAuth.Post("/login", apihandlers.HandleLoginApi)
	apiAuth.Post("/parent/login", apihandlers.HandleParentLoginApi)
	apiAuth.Post("/logout", apihandlers.HandleLogoutApi)

	/* -
	   -
	   - space for clarity
	   -
	   -
	*/

	// ---------------- group web routes
	web := app.Group("/")

	// web routes
	web.Get("/", webhandlers.HandleIndex)
	web.Get("/about", webhandlers.HandleAbout)
	web.Get("/dashboard", middlewares.IsAuthenticated, webhandlers.HandleDashboard)
	web.Get("/resetpassword", middlewares.IsNotAuthenticated, webhandlers.HandleResetPassword)

	// group auth routes for web
	webAuth := web.Group("/auth", middlewares.IsNotAuthenticated)

	// web auth routes
	webAuth.Get("/signup", webhandlers.HandleSignup)
	webAuth.Get("/parent/login", webhandlers.HandleParentLogin)
	webAuth.Get("/login", webhandlers.HandleLogin)
}
