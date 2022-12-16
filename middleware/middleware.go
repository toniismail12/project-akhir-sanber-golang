package middleware

import (
	"project-sanber/database"
	"project-sanber/repository"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(c *fiber.Ctx) error {

	// Get token auth
	token := c.Get("Authorization")

	info_user, err1 := repository.Info_user_login(database.DbConnection, token)
	if err1 != nil {
		panic(err1)
	}

	token_tersimpan := ""

	for _, d := range info_user {
		token_tersimpan = d.Token
	}

	// claim token, aktif atau kadaluarsa
	if _, err := Parsejwt(token_tersimpan); err != nil {

		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated || Access Denied 1",
		})
	}

	// cek ke database login
	if token_tersimpan == "" || token_tersimpan != token {

		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated || Access Denied 2",
		})
	}

	return c.Next()
}
