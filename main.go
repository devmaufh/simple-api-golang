package main

import (
	"api-example.com/book"
	"api-example.com/database"
	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World from fiber refactorized")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/book", book.GetBooks)
	v1.Get("/book/:id", book.GetBook)
	v1.Post("/book", book.NewBook)
	v1.Patch("/book/:id", book.UpdateBook)
	v1.Delete("/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=dev password=12345678 dbname=books_api_golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	database.DB.AutoMigrate(&book.Book{})
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

}
