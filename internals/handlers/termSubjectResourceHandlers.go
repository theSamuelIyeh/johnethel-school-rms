package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TermSubject struct {
	Name string
}

// get all term subjects
func (h *Handler) GetAllTermSubjectHandler(c echo.Context) error {
	// err, sessions := services.GetAllSession()
	return c.String(http.StatusOK, "get all sessions route")
}

// get a single term subject
func (h *Handler) GetTermSubjectHandler(c echo.Context) error {
	sessionID := c.Param("id")
	return c.String(http.StatusOK, "get a single session"+sessionID)
}

// create a single term subject
func (h *Handler) CreateTermSubjectHandler(c echo.Context) error {
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

// update a single term subject
func (h *Handler) UpdateTermSubjectHandler(c echo.Context) error {
	return c.String(http.StatusOK, "updated successfully")
}

// delete a single term subject
func (h *Handler) DeleteTermSubjectHandler(c echo.Context) error {
	return c.String(http.StatusOK, "deleted successfully")
}
