package handlers

import (
	"PZN/config"
	"PZN/entities"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var users []entities.User

	config.Database.Find(&users)
	return c.Status(200).JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.User

	result := config.Database.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{
			"message": "User Not Found",
		})
	}

	return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&user)
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

func DeleteUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	var user entities.User

	result := config.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(map[string]string{
			"message": "Data Uer not found, please check again",
		})
	}

	return c.Status(200).JSON(map[string]string{
		"message": "User success deleted",
	})
}
