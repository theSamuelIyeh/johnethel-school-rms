package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/views/fragments"
)

type Term struct {
	Name string
}

// get all terms
func (h *Handler) GetAllTermHandler(c echo.Context) error {
	// err, sessions := services.GetAllSession()
	return c.String(http.StatusOK, "get all term route")
}

// get a single term
func (h *Handler) GetTermDropdownHandler(c echo.Context) error {
	role := c.Param("role")
	if role == "admin" {
		result, err := h.Config.DB.GetAllTerms(c.Request().Context())
		if err != nil {
			return fragments.ErrorMessageComponent("Error Loading Terms").Render(c.Request().Context(), c.Response().Writer)
		}
		return fragments.AdminTermDropdown(result).Render(c.Request().Context(), c.Response().Writer)
	} else {
		result, err := h.Config.DB.GetAllUnlockedTerms(c.Request().Context())
		if err != nil {
			return fragments.ErrorMessageComponent("Error Loading Terms").Render(c.Request().Context(), c.Response().Writer)
		}
		return fragments.TermDropdown(result).Render(c.Request().Context(), c.Response().Writer)
	}
}

// create a single term
func (h *Handler) CreateTermHandler(c echo.Context) error {
	term := new(Term)
	if err := c.Bind(term); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	if err := c.Validate(term); err != nil {
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showError\">Bad Request</p>")
	}
	return c.String(http.StatusOK, "term created successfully")
}

// update a single term
func (h *Handler) UpdateTermHandler(c echo.Context) error {
	return c.String(http.StatusOK, "updated successfully")
}

// delete a single term
func (h *Handler) DeleteTermHandler(c echo.Context) error {
	return c.String(http.StatusOK, "deleted successfully")
}
