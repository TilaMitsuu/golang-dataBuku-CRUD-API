package routes

import (
	"book-api/controllers"
	"book-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	book := app.Group("/books", middlewares.Protected())
	book.Get("/ambil", controllers.GetBooks)
	book.Post("/buat", controllers.CreateBook)
	book.Put("/:id", controllers.UpdateBook)
	book.Delete("/:id", controllers.DeleteBook)
}	