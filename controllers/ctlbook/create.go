package ctlbook

import (
	"gofiber-restapi/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {

	var book models.Book

	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	book.PublishDate = time.Now().Format("2006-01-02")

	gtx := models.DB.Create(&book)
	if gtx.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": gtx.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}
