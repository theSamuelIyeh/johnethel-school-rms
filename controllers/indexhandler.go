package controllers

import "github.com/gofiber/fiber/v2"

func IndexController(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"name": "samuel"}, "layouts/main")
}

func IndexApiController(c *fiber.Ctx) error {
	return c.SendString("This is backend")
}

func AboutController(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{}, "layouts/main")
}

func SignupController(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{}, "layouts/main")
}

func LoginController(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{}, "layouts/main")
}
