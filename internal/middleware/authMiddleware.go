package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thesamueliyeh/johnethel-rms/internal/services"
)

func IsAuthenticated(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		services.TokenService.ClearTokens(c)
		return services.RedirectService.RedirectService(c, "/auth/login")
		// c.Set("HX-Redirect", "/auth/login")
		// return c.Redirect("/auth/login")
	}
	shouldRefreshAccessToken, err := services.TokenService.VerifyToken(accessToken)
	if err != nil {
		services.TokenService.ClearTokens(c)
		return services.RedirectService.RedirectService(c, "/auth/login")
		// c.Set("HX-Redirect", "/auth/login")
		// return c.Redirect("/auth/login")
	}
	if shouldRefreshAccessToken {
		token, err := services.TokenService.RefreshAccessToken(accessToken)
		if err != nil {
			services.TokenService.ClearTokens(c)
			return services.RedirectService.RedirectService(c, "/auth/login")
			// c.Set("HX-Redirect", "/auth/login")
			// return c.Redirect("/auth/login")
		}
		c.Cookie(&fiber.Cookie{
			Name:  "access_token",
			Value: token.AccessToken,
			// Secure: true,
			Path:     "/",
			HTTPOnly: true,
			Expires:  token.AccessTokenExp,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "id",
			Value: token.UserId,
			// Secure: true,
			Path:     "/",
			HTTPOnly: true,
			Expires:  token.AccessTokenExp,
		})
		c.Cookie(&fiber.Cookie{
			Name:  "role",
			Value: token.Role,
			// Secure: true,
			Path:     "/",
			HTTPOnly: true,
			Expires:  token.AccessTokenExp,
		})
	}
	user, err := services.UserService.GetUserById(c.Cookies("id"), c.Cookies("role"))
	if err != nil {
		services.TokenService.ClearTokens(c)
		return services.RedirectService.RedirectService(c, "/auth/login")
		// c.Set("HX-Redirect", "/auth/login")
		// return c.Redirect("/auth/login")
	}
	c.Locals("user", user)
	return c.Next()
}

func IsNotAuthenticated(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		return c.Next()
	}
	shouldRefreshAccessToken, err := services.TokenService.VerifyToken(accessToken)
	if err != nil {
		services.TokenService.ClearTokens(c)
		return c.Next()
	}
	if shouldRefreshAccessToken {
		if c.Path() == "/api/auth/logout" {
			return c.Next()
		}
		return services.RedirectService.RedirectService(c, "/dashboard")
		// c.Set("HX-Redirect", "/dashboard")
		// return c.Redirect("/dashboard")
	}
	if c.Path() == "/api/auth/logout" {
		return c.Next()
	}
	return services.RedirectService.RedirectService(c, "/dashboard")
	// c.Set("HX-Redirect", "/dashboard")
	// return c.Redirect("/dashboard")
}
