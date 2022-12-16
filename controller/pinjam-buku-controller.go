package controller

import (
	"log"
	"project-sanber/database"
	"project-sanber/repository"
	"project-sanber/structs"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Insert_pinjam(c *fiber.Ctx) error {
	var data map[string]any

	if err := c.BodyParser(&data); err != nil {
		log.Println("unable to parse body")
	}

	currentTime := time.Now()
	now := currentTime.Local().Format("2 Jan 2006 15:04")

	borrow := structs.Pinjam_buku{
		Id_buku:          data["id_buku"].(string),
		Id_user_peminjam: data["id_user_peminjam"].(string),
		Created_by:       data["created_by"].(string),
		Updated_at:       now,
		Updated_by:       data["updated_by"].(string),
	}

	err2 := repository.Insert_pinjam(database.DbConnection, borrow)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(fiber.Map{
		"message": "berhasil create data",
	})
}

func Get_peminjam(c *fiber.Ctx) error {

	user, err := repository.Get_peminjam(database.DbConnection)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"data": user,
	})

}
