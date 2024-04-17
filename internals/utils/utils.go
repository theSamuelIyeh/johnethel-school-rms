package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/iancoleman/strcase"
	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/views/layouts"
)

type DashboardConfig struct {
	Role     string
	UserName string
	Avatar   string
}

func getDashboardConfig(c echo.Context) DashboardConfig {
	return DashboardConfig{
		Role:     c.Get("role").(string),
		UserName: c.Get("userName").(string),
		Avatar:   c.Get("avatar").(string),
	}
}

func CreateCookie(c echo.Context, token, name string, expires int) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = fmt.Sprintf("%s %s", "Bearer", token)
	cookie.HttpOnly = true
	cookie.Path = "/"
	if expires == -1 {
		// Set expiration to a past time to expire the cookie immediately
		cookie.Expires = time.Unix(1, 0)
	} else if expires > 0 {
		// Set expiration based on the provided duration
		cookie.Expires = time.Now().Add(time.Duration(expires) * time.Second)
	} else {
		// No expiration
		cookie.Expires = time.Time{}
	}
	c.SetCookie(cookie)
}

func CreateOrdinaryCookie(c echo.Context, token, name string, expires int) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = fmt.Sprintf("%s %s", "Bearer", token)
	cookie.HttpOnly = false
	cookie.Path = "/"
	if expires == -1 {
		// Set expiration to a past time to expire the cookie immediately
		cookie.Expires = time.Unix(1, 0)
	} else if expires > 0 {
		// Set expiration based on the provided duration
		cookie.Expires = time.Now().Add(time.Duration(expires) * time.Second)
	} else {
		// No expiration
		cookie.Expires = time.Time{}
	}
	c.SetCookie(cookie)
}

func IsHXRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

func HXTriggerName(c echo.Context) string {
	return c.Request().Header.Get("HX-Trigger-Name")
}

func RedirectToUrl(c echo.Context, url string) error {
	if IsHXRequest(c) {
		c.Response().Header().Set("Hx-redirect", url)
		c.Response().Header().Set("hx-trigger", "swap")
		return c.String(http.StatusPermanentRedirect, "")
	}
	return c.Redirect(http.StatusPermanentRedirect, url)
}

func RenderOrdinaryTemplate(c echo.Context, templates TemplateStruct) error {
	_, theme, err := GetTokenFromCookie(c, "theme")
	if err != nil {
		theme = "false"
	}
	if IsHXRequest(c) {
		if strings.HasSuffix(HXTriggerName(c), "-component-trigger") {
			return templates.Component.Render(c.Request().Context(), c.Response().Writer)
		} else if strings.HasSuffix(HXTriggerName(c), "-page-trigger") {
			return templates.Page.Render(c.Request().Context(), c.Response().Writer)
		}
	}
	return layouts.SiteLayout(templates.Page, templates.Title, theme).Render(c.Request().Context(), c.Response().Writer)
}

func RenderDashboardTemplate(c echo.Context, page templ.Component, title string) error {
	config := getDashboardConfig(c)
	_, theme, err := GetTokenFromCookie(c, "theme")
	if err != nil {
		theme = "false"
	}
	if IsHXRequest(c) {
		return page.Render(c.Request().Context(), c.Response().Writer)
	}
	return layouts.DashboardLayout(page, title, theme, config.Role, config.UserName, strcase.ToKebab(title), config.Avatar).Render(c.Request().Context(), c.Response().Writer)
}

func GetTokenFromCookie(c echo.Context, tokenName string) (*http.Cookie, string, error) {
	accessCookie, err := c.Cookie(tokenName)
	if err != nil {
		return nil, "", err
	}
	accessTokenParts := strings.Split(accessCookie.Value, " ")
	if len(accessTokenParts) != 2 || accessTokenParts[0] != "Bearer" {
		return nil, "", errors.New("invalid access token format")
	}
	return accessCookie, accessTokenParts[1], nil
}
