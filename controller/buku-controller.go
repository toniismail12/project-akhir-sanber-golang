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

func Get_buku(c *fiber.Ctx) error {

	buku, err := repository.GetAllBuku(database.DbConnection)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"data": buku,
	})

}

func Insert_buku(c *fiber.Ctx) error {
	var data map[string]any

	if err := c.BodyParser(&data); err != nil {
		log.Println("unable to parse body")
	}

	tahun_terbit, _ := strconv.Atoi(data["tahun_terbit"].(string))

	currentTime := time.Now()
	now := currentTime.Local().Format("2 Jan 2006 15:04")

	buku := structs.Buku{
		Judul_buku:     data["judul_buku"].(string),
		Author:         data["author"].(string),
		Tahun_terbit:   tahun_terbit,
		Publisher:      data["publisher"].(string),
		ISBN:           data["isbn"].(string),
		Jumlah_halaman: data["jumlah_halaman"].(string),
		Is_dipinjam:    "no",
		Created_at:     now,
		Created_by:     data["created_by"].(string),
		Updated_at:     now,
		Updated_by:     data["updated_by"].(string),
	}

	err2 := repository.InsertBuku(database.DbConnection, buku)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(fiber.Map{
		"message": "berhasil create data",
	})
}

func Update_buku(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		log.Println("unable to parse body")
	}

	tahun_terbit, _ := strconv.Atoi(data["tahun_terbit"].(string))

	currentTime := time.Now()
	now := currentTime.Local().Format("2 Jan 2006 15:04")

	buku := structs.Buku{
		Judul_buku:     data["judul_buku"].(string),
		Author:         data["author"].(string),
		Tahun_terbit:   tahun_terbit,
		Publisher:      data["publisher"].(string),
		ISBN:           data["isbn"].(string),
		Jumlah_halaman: data["jumlah_halaman"].(string),
		Is_dipinjam:    data["is_dipinjam"].(string),
		Updated_at:     now,
		Updated_by:     data["updated_by"].(string),
		ID:             id,
	}

	err2 := repository.UpdateBuku(database.DbConnection, buku)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(fiber.Map{
		"message": "berhasil update data",
	})

}

func Delete_buku(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err2 := repository.DeleteBuku(database.DbConnection, id)
	if err2 != nil {
		panic(err2)
	}

	return c.JSON(fiber.Map{
		"message": "berhasil delete",
	})
}
