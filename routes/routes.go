package routes

import (
	"project-sanber/controller"
	"project-sanber/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", controller.Main)

	app.Post("/login", controller.Login)

	// wajib login untuk akses route dibawah ini
	app.Use(middleware.IsAuthenticate)

	// buku
	app.Get("/buku", controller.Get_buku)

	app.Post("/buku", controller.Insert_buku)

	app.Put("/buku", controller.Insert_buku)

	app.Delete("/buku/:id", controller.Delete_buku)

	// pinjam
	app.Get("/pinjam", controller.Get_peminjam)

	app.Post("/pinjam", controller.Insert_pinjam)

	// user
	app.Get("/user", controller.Get_user)

	app.Post("/user", controller.Insert_user)

	app.Delete("/user/:id", controller.Delete_user)

}
