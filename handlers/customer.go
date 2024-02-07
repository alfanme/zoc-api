package handlers

import (
	"zoc-api/database"
	"zoc-api/models"

	"github.com/gofiber/fiber/v2"
)


func GetCustomers(c *fiber.Ctx) error {
	var data []models.Customer

	query := database.Client.Find(&data)

	if query.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.Response{
				Error: query.Error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		models.Response{
			Data: data,
		},
	)
}


func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.Customer

	query := database.Client.First(&data, id)

	if query.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.Response{
				Error: query.Error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		models.Response{
			Data: data,
		},
	)
}


func CreateCustomer(c *fiber.Ctx) error {
	var mutation models.Request[models.Customer]

	err := c.BodyParser(&mutation)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Error: err.Error(),
		})
	}

	data := models.Customer{
		Name: mutation.Data.Name,
		Phone: mutation.Data.Phone,
		Address: mutation.Data.Address,
	}

	query := database.Client.Create(&data)

	if query.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.Response{
				Error: query.Error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		models.Response{
			Data: data,
			Message: "Successfully created",
		},
	)
}


func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	var mutation models.Request[models.CustomerMutation]

	err := c.BodyParser(&mutation)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Error: err.Error(),
		})
	}

	var data models.Customer

	query := database.Client.First(&data, id)

	if query.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.Response{
				Error: query.Error.Error(),
			},
		)
	}

	data.Name = mutation.Data.Name
	data.Phone = mutation.Data.Phone
	data.Address = mutation.Data.Address

	query = database.Client.Save(&data)

	if query.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.Response{
				Error: query.Error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		models.Response{
			Data: data,
			Message: "Successfully updated",
		},
	)
}


func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.Customer

	query := database.Client.First(&data, id)

	if query.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.Response{
				Error: query.Error.Error(),
			},
		)
	}

	query = database.Client.Delete(&data)

	if query.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.Response{
				Error: query.Error.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		models.Response{
			Data: data,
			Message: "Successfully deleted",
		},
	)
}