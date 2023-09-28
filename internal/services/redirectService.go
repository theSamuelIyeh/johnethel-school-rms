package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thesamueliyeh/johnethel-rms/internal/database"
)

type redirectService struct {
	db database.Database
}

func NewRedirectService(db database.Database) *redirectService {
	return &redirectService{db: db}
}

func (s *redirectService) RedirectService(c *fiber.Ctx, route string) error {
	if c.Get("HX-Request") == "true" {
		c.Set("HX-Redirect", route)
		return c.Status(fiber.StatusSeeOther).SendString("") // No need to generate a response body
	}

	// For a normal load, perform a redirect
	return c.Redirect(route)
}
