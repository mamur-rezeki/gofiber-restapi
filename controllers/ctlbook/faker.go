package ctlbook

import (
	"gofiber-restapi/models"

	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Faker(c *fiber.Ctx) error {
	qty, err := c.ParamsInt("qty", 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "Invalid qty",
			},
		)
	}

	models.DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < qty; i++ {
			book := models.Book{
				Title:       faker.Sentence(),
				Author:      faker.Name(),
				Description: faker.Sentence(),
				PublishDate: faker.Date(),
			}
			gtx := tx.Create(&book)
			if gtx.Error != nil {
				return gtx.Error
			}
		}

		return err
	})

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Faker created",
			"qty":     qty,
		},
	)
}
