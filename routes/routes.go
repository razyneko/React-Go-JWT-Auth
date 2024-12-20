package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/razyneko/React-Go-JWT-Auth/controllers"
)

func Setup(app *fiber.App) {
	
    // Define a route for the GET method on the root path '/'
    app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.LogOut)


}