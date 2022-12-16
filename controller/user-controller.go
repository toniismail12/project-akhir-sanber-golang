package controller

import (
	"log"
	"project-sanber/database"
	"project-sanber/repository"
	"project-sanber/structs"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Get_user(c *fiber.Ctx) error {

	user, err := repository.Get_User(database.DbConnection)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"data": user,
	})

}

func Insert_user(c *fiber.Ctx) error {
	var data map[string]any

	if err := c.BodyParser(&data); err != nil {
		log.Println("unable to parse body")
	}

	currentTime := time.Now()
	now := currentTime.Local().Format("2 Jan 2006 15:04")

	user := structs.Users{
		Username:   data["username"].(string),
		Password:   data["password"].(string),
		Email:      data["email"].(string),
		Wa:         data["wa"].(string),
		Nama:       data["nama"].(string),
		Alamat:     data["alamat"].(string),
		Role:       data["role"].(string),
		Created_at: now,
		Created_by: data["created_by"].(string),
		Updated_at: now,
		Updated_by: data["updated_by"].(string),
	}

	err2 := repository.Insert_user(database.DbConnection, user)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(fiber.Map{
		"message": "berhasil create data",
	})
}

func Delete_user(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err2 := repository.Delete_user(database.DbConnection, id)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(fiber.Map{
		"message": "berhasil delete",
	})
}
