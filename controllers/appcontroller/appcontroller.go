package appcontroller

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"title": "WDB Studios",
	}, "layouts/master")
}

func About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"title": "About WDB Studios",
	}, "layouts/master")
}

func Services(c *fiber.Ctx) error {
	return c.Render("services", fiber.Map{
		"title": "What We Do At WDB Studios",
	}, "layouts/master")
}

func Contact(c *fiber.Ctx) error {
	return c.Render("contact", fiber.Map{
		"title": "Talk to Us-WDB Studios",
	}, "layouts/master")
}
