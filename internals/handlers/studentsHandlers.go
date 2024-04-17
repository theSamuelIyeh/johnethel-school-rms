package handlers

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/database"
	"github.com/theSamuelIyeh/johnethel-school-rms/views/fragments"
)

func (h *Handler) GetAllStudentsHandler(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		return fragments.ErrorMessageComponent("Invalid Page Figure").Render(c.Request().Context(), c.Response().Writer)
	}
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		return fragments.ErrorMessageComponent("Invalid Limit Figure").Render(c.Request().Context(), c.Response().Writer)
	}
	page = (page - 1) * limit
	result, err := h.Config.DB.GetAllStudent(c.Request().Context(), database.GetAllStudentParams{
		Offset: int32(page),
		Limit:  int32(limit),
	})
	if err != nil {
		fmt.Println(err)
		return fragments.ErrorMessageComponent("Error Getting Students").Render(c.Request().Context(), c.Response().Writer)
	}
	return fragments.StudentsTableBody(result, "512", "Students Data Loaded Successfully").Render(c.Request().Context(), c.Response().Writer)
}
