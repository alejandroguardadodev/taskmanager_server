package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"taskmanagerserver.com/api/database"
)

func migrate() {
	log.Println("Running Migation")
	//database.DB.AutoMigrate(&models.Project{}, &models.Task{})
}

func main() {
	godotenv.Load()

	database.Conect()

	migrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", //",other"
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	url := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Fatal(app.Listen(url))
}