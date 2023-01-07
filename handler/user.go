package handler

import (
	"github.com/gofiber/fiber/v2"

	"hanhngo.me/m/database"
	"hanhngo.me/m/model"
)

func GetUserList(c *fiber.Ctx) error {
	db := database.DB
	var userList []model.User
	db.Find(&userList)
	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"data": userList,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.First(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{
			"code":  404,
			"error": "user not found",
			"data":  nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"data": user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":  400,
			"data":  nil,
			"error": err,
		})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":  500,
			"data":  nil,
			"error": "Couldn't create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"code": 201,
		"data": user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUsername struct {
		Names string
	}
	db := database.DB
	id := c.Params("id")
	var existedUser model.User
	db.First(&existedUser, id)

	var updateUsername UpdateUsername

	if err := c.BodyParser(&updateUsername); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":  400,
			"error": "Invalid arguments",
		})
	}

	existedUser.Names = updateUsername.Names
	db.Save(&existedUser)

	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"data": existedUser,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var existedUser model.User
	db.First(&existedUser, id)
	db.Delete(&existedUser)

	return c.Status(200).JSON(fiber.Map{
		"code": 200,
		"data": "OK",
	})
}
