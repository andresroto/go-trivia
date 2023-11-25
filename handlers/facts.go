package handlers

import (
	"github.com/andresroto/go-trivia/database"
	"github.com/andresroto/go-trivia/models"
	"github.com/gofiber/fiber/v2"
)

// ListFacts retrieves a list of facts.
func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(fiber.StatusOK).JSON(facts)
}

// ShowFact retrieves a fact from the database based on the given ID.
func ShowFact(c *fiber.Ctx) error {
	id := c.Params("id")

	var fact = models.Fact{}

	result := database.DB.Db.First(&fact, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fact)
}

// CreateFact creates a new fact by parsing the request body and saving it to the database.
func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Fact created successfully",
		"data":    fact,
	})
}

// UpdateFact updates a fact in the database.
func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")

	var fact = models.Fact{}

	result := database.DB.Db.First(&fact, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result = database.DB.Db.Save(&fact)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fact updated successfully",
		"data":    fact,
	})
}

// DeleteFact deletes a fact from the database based on the provided ID.
func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")

	var fact = models.Fact{}

	result := database.DB.Db.First(&fact, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	result = database.DB.Db.Delete(&fact)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fact deleted successfully",
		"data":    fact,
	})
}
