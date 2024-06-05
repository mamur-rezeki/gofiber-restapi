package ctlbook

import (
	"gofiber-restapi/models"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id book not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "book deleted",
	})
}

func DeleteRange(c *fiber.Ctx) error {

	var start, end = c.Params("start"), c.Params("end")
	if models.DB.Delete(&models.Book{}, "id BETWEEN ? AND ?", start, end).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id book not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "books deleted",
	})
}
