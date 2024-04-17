package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TermClass struct {
	Name string
}

// get all term class
func (h *Handler) GetAllTermClassHandler(c echo.Context) error {
	// err, sessions := services.GetAllSession()
	return c.String(http.StatusOK, "get all sessions route")
}

// get a single term class
func (h *Handler) GetTermClassHandler(c echo.Context) error {
	termClassID := c.Param("id")
	return c.String(http.StatusOK, "get a single term class "+termClassID)
}

// create a single term class
func (h *Handler) CreateTermClassHandler(c echo.Context) error {
	termClass := new(TermClass)
	if err := c.Bind(termClass); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	if err := c.Validate(termClass); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	return c.String(http.StatusOK, "term class created successfully")
}

// update a single term class
func (h *Handler) UpdateTermClassHandler(c echo.Context) error {
	return c.String(http.StatusOK, "updated successfully")
}

// delete a single term class
func (h *Handler) DeleteTermClassHandler(c echo.Context) error {
	return c.String(http.StatusOK, "deleted successfully")
}
