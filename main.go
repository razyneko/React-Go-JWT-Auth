package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/razyneko/React-Go-JWT-Auth/database"
	"github.com/razyneko/React-Go-JWT-Auth/routes"
)

func main() {
    database.Connect()
   // Initialize a new Fiber app
    app := fiber.New()

    app.Use(cors.New(cors.Config{
        // for frontend to use cookie
        AllowCredentials: true,
    }))

    routes.Setup(app)
    // Start the server on port 8000
    log.Fatal(app.Listen(":8000"))
}