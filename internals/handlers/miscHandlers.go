package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/utils"
	"github.com/theSamuelIyeh/johnethel-school-rms/views/pages"
)

type Handler struct {
	Config *utils.ApiCfg
}

func NewHandler(cfg *utils.ApiCfg) *Handler {
	return &Handler{Config: cfg}
}

// index page handler
func (h *Handler) GetIndexPageHandler(c echo.Context) error {
	title := "Johnethel School RMS"
	templates := utils.TemplateStruct{Title: title, Page: pages.IndexPage(title)}
	return utils.RenderOrdinaryTemplate(c, templates)
}

func (h *Handler) SetTheme(c echo.Context) error {
	value := c.FormValue("checkbox")
	if value == "myDarkTheme" {
		theme := "true"
		utils.CreateOrdinaryCookie(c, theme, "theme", 0)
	} else {
		theme := "false"
		utils.CreateOrdinaryCookie(c, theme, "theme", 0)
	}
	return c.String(200, "Theme set successfully")
}
