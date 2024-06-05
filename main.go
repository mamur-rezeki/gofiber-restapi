package main

import (
	"gofiber-restapi/controllers/ctlbook"
	"gofiber-restapi/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDB()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Println(c.Method(), c.Path())

		return c.Next()
	})

	api := app.Group("/api")
	books := api.Group("/books")

	books.Get("/", ctlbook.Index)
	books.Get("/:id", ctlbook.Show)
	books.Post("/", ctlbook.Create)
	books.Put("/:id", ctlbook.Update)
	books.Delete("/:id", ctlbook.Delete)
	books.Delete("/:start/:end", ctlbook.DeleteRange)

	books.Put("/faker/:qty", ctlbook.Faker)

	app.Listen(":26000")
}
