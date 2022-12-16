package controller

import "github.com/gofiber/fiber/v2"

func Main(c *fiber.Ctx) error {

	c.Status(200)
	return c.JSON(fiber.Map{
		"app_name":   "Data Buku Perpustakaan Sekolah",
		"descripsi":  "project sanber",
		"created_by": "Toni Ismail",
	})

}
