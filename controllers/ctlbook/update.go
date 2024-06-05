package ctlbook

import (
	"gofiber-restapi/models"

	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {

	id := c.Params("id")

	var book models.Book

	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "No rows updated",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(book)
}
