package controller

import (
	"log"
	"project-sanber/database"
	"project-sanber/middleware"
	"project-sanber/repository"
	"project-sanber/structs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		log.Println("unable to parse body")
	}

	user_info, err1 := repository.Detail_user(database.DbConnection, data["username"].(string))
	if err1 != nil {
		panic(err1)
	}

	id := ""
	nama := ""
	role := ""
	password := ""

	for _, d := range user_info {
		id = strconv.Itoa(int(d.ID))
		nama = d.Nama
		role = d.Role
		password = d.Password
	}

	if nama == "" || role == "" {
		return c.JSON(fiber.Map{
			"message": "login gagal",
		})
	}

	if data["password"].(string) != password {
		return c.JSON(fiber.Map{
			"message": "login gagal",
		})
	}

	token, errt := middleware.GenerateJwt(id)
	if errt != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	user := structs.User_login{
		Username: data["username"].(string),
		Nama:     nama,
		Role:     role,
		Token:    token,
	}

	err2 := repository.Insert_user_login(database.DbConnection, user)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(fiber.Map{
		"message": "berhasil Login",
		"data":    user,
	})

}
