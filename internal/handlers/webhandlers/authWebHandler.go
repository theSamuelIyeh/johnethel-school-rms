package webhandlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"name": "samuel"}, "layouts/main")
}

func HandleAbout(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{}, "layouts/main")
}

func HandleSignup(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{}, "layouts/main")
}

func HandleLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{}, "layouts/main")
}

func HandleDashboard(c *fiber.Ctx) error {
	user := c.Locals("user")
	fmt.Println(user)
	return c.Render("dashboard", fiber.Map{"user": user}, "layouts/main")
}

func HandleParentLogin(c *fiber.Ctx) error {
	return c.Render("parentlogin", fiber.Map{}, "layouts/main")
}

func HandleResetPassword(c *fiber.Ctx) error {
	return c.Render("resetpassword", fiber.Map{}, "layouts/main")
}
