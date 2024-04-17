package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Session struct {
	Name string
}

// get all sessions handler
func (h *Handler) GetAllSessionHandler(c echo.Context) error {
	// err, sessions := services.GetAllSession()
	return c.String(http.StatusOK, "get all sessions route")
}

// get a single session
func (h *Handler) GetSessionHandler(c echo.Context) error {
	sessionID := c.Param("id")
	return c.String(http.StatusOK, "get a single session"+sessionID)
}

// create a single session
func (h *Handler) CreateSessionHandler(c echo.Context) error {
	session := new(Session)
	if err := c.Bind(session); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	if err := c.Validate(session); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	return c.String(http.StatusOK, "session created successfully")
}

// update a single session
func (h *Handler) UpdateSessionHandler(c echo.Context) error {
	return c.String(http.StatusOK, "updated successfully")
}

// delete a single session
func (h *Handler) DeleteSessionHandler(c echo.Context) error {
	return c.String(http.StatusOK, "deleted successfully")
}
