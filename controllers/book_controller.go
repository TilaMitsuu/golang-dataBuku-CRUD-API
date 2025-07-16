package controllers

import (
	"book-api/config"
	"book-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	config.DB.Find(&books)

	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "No books found",
		})
	}

	return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	if c.Method() != "POST" {
		return c.Status(405).JSON(fiber.Map{
			"error":   "Method not gabisa digunakan",
			"message": "Use POST method to create a book",
		})
	}
	book := new(models.Book)
	if err := c.BodyParser(&book); err != nil {
		return err
	}
	config.DB.Create(&book)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	config.DB.First(&book, id)

	if book.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	if err := c.BodyParser(&book); err != nil {
		return err
	}

	config.DB.Save(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	config.DB.First(&book, id)

	if book.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	config.DB.Delete(&book)
	return c.JSON(fiber.Map{
		"message": "Book deleted",
	})
}
