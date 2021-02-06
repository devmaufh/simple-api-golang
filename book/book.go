package book

import (
	"api-example.com/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

//Book structure for manage books
type Book struct {
	ID   uint
	Name string
	gorm.Model
}

//GetBooks Returns all books
func GetBooks(c *fiber.Ctx) {
	db := database.DB
	var books []Book
	db.Find(&books)
	data := map[string][]Book{
		"data": books,
	}
	c.JSON(data)
}

//GetBook returns given book
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var book Book
	db.Find(&book, id)
	data := map[string]Book{
		"data": book,
	}
	c.JSON(data)
}

//NewBook store a new book
func NewBook(c *fiber.Ctx) {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db := database.DB
	db.Create(book)
	data := map[string]Book{
		"data": *book,
	}
	c.JSON(data)
}

//UpdateBook update given book
func UpdateBook(c *fiber.Ctx) {
	type DataUpdateBook struct {
		Name string `json:"name"`
	}
	var updatedData DataUpdateBook
	if err := c.BodyParser(&updatedData); err != nil || updatedData.Name == "" {
		c.Status(503).Send(err)
	}
	var book Book
	id := c.Params("id")
	db := database.DB
	db.First(&book, id)
	db.Model(&book).Update("name", updatedData.Name)
	data := map[string]string{
		"data": "book updated!",
	}
	c.JSON(data)
}

//DeleteBook delete given book
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	var book Book
	database.DB.Find(&book, id)
	if book.Name == "" {
		c.Status(404).Send("Book not found")
	}
	database.DB.Delete(&book)
	data := map[string]string{
		"data": "Book deleted!",
	}
	c.JSON(data)
}
