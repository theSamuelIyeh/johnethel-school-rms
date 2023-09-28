package apihandlers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/thesamueliyeh/johnethel-rms/internal/services"
)

func HandleIndexApi(c *fiber.Ctx) error {
	return c.SendString("This is backend")
}

func HandleLoginApi(c *fiber.Ctx) error {
	splittedAdmissionNo := strings.Split(c.FormValue("admissionNo"), "/")
	if splittedAdmissionNo == nil || len(splittedAdmissionNo) != 3 {
		c.Type(".html")
		return c.SendString("<h1>*Invalid Admission Number or Password</h1>")
	}
	if splittedAdmissionNo[0] != "js" || splittedAdmissionNo[1] != "ogb" {
		c.Type(".html")
		return c.SendString("<h1>*Invalid Admission Number or Password</h1>")
	}
	pattern := `([ps])(\d+)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	subStrings := regex.FindStringSubmatch(splittedAdmissionNo[2])
	if len(subStrings) != 3 {
		c.Type(".html")
		return c.SendString("<h1>*Invalid Admission Number or Password</h1>")
	}
	role := subStrings[1]
	if role == "p" {
		role = "pupil"
	} else if role == "s" {
		role = "staff"
	} else {
		c.Type(".html")
		return c.SendString("<h1>*Invalid Admission Number or Password</h1>")
	}
	admissionNo := subStrings[2]
	password := c.FormValue("password")

	user, err := services.CredentialsService.VerifyCredentials(admissionNo, password, role)
	if err != nil {
		c.Type(".html")
		return c.SendString("<h1>*Invalid Admission Number or Password</h1>")
	}
	token, err := services.TokenService.GenerateToken(user)
	if err != nil {
		c.Type(".html")
		return c.SendString("<h1>*Error logging in</h1>")
	}
	err = services.TokenService.SetRefreshToken(token.RefreshToken, token.RefreshTokenExp, token.UserId)
	if err != nil {
		c.Type(".html")
		fmt.Println(err)
		return c.SendString("<h1>*Error logging in</h1>")
	}
	c.Cookie(&fiber.Cookie{
		Name:  "access_token",
		Value: token.AccessToken,
		// Secure:   true,
		Path:     "/",
		HTTPOnly: true,
		Expires:  token.AccessTokenExp,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "id",
		Value: token.UserId,
		// Secure:   true,
		Path:     "/",
		HTTPOnly: true,
		Expires:  token.AccessTokenExp,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "role",
		Value: token.Role,
		// Secure:   true,
		Path:     "/",
		HTTPOnly: true,
		Expires:  token.AccessTokenExp,
	})
	return services.RedirectService.RedirectService(c, "/dashboard")
	// c.Set("HX-Redirect", "/dashboard")
	// return c.Status(fiber.StatusSeeOther).SendString("")
}

func HandleParentLoginApi(c *fiber.Ctx) error {
	role := "parent"
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := services.CredentialsService.VerifyCredentials(email, password, role)
	if err != nil {
		c.Type(".html")
		return c.SendString("<h1>*Invalid Admission Number or Password</h1>")
	}
	token, err := services.TokenService.GenerateToken(user)
	if err != nil {
		c.Type(".html")
		return c.SendString("<h1>*Error logging in</h1>")
	}
	err = services.TokenService.SetRefreshToken(token.RefreshToken, token.RefreshTokenExp, token.UserId)
	if err != nil {
		c.Type(".html")
		return c.SendString("<h1>*Error logging in</h1>")
	}
	c.Cookie(&fiber.Cookie{
		Name:  "access_token",
		Value: token.AccessToken,
		// Secure:   true,
		Path:     "/",
		HTTPOnly: true,
		Expires:  token.AccessTokenExp,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "id",
		Value: token.UserId,
		// Secure:   true,
		Path:     "/",
		HTTPOnly: true,
		Expires:  token.AccessTokenExp,
	})
	c.Cookie(&fiber.Cookie{
		Name:  "role",
		Value: token.Role,
		// Secure:   true,
		Path:     "/",
		HTTPOnly: true,
		Expires:  token.AccessTokenExp,
	})
	return services.RedirectService.RedirectService(c, "/dashboard")
	// c.Set("HX-Redirect", "/dashboard")
	// return c.Status(fiber.StatusSeeOther).SendString("")
}

func HandleLogoutApi(c *fiber.Ctx) error {
	err := services.TokenService.ClearTokens(c)
	if err != nil {
		return err
	}
	return services.RedirectService.RedirectService(c, "/auth/login")
	// c.Set("HX-Redirect", "/auth/login")
	// return c.Status(fiber.StatusSeeOther).SendString("")
}
