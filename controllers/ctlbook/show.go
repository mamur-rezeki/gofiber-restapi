package ctlbook

import (
	"gofiber-restapi/models"

	"github.com/gofiber/fiber/v2"
)

func Show(c *fiber.Ctx) error {

	id := c.Params("id")

	book := models.Book{}
	if err := models.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}
