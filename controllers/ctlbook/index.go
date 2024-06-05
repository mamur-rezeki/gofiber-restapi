package ctlbook

import (
	"gofiber-restapi/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book

	models.DB.Find(&books)

	return c.Status(fiber.StatusOK).JSON(books)

}
