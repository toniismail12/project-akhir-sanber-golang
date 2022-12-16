package main

import (
	"database/sql"
	"fmt"
	"os"
	"project-sanber/database"
	"project-sanber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Faild load file env")
	} else {
		fmt.Println("Success read file env")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Connection Failed")
		panic(err)
	} else {
		fmt.Println("Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	app := fiber.New()

	// router
	routes.Setup(app)

	// get port
	port := os.Getenv("PORT")

	// run
	app.Listen(":" + port)
}
