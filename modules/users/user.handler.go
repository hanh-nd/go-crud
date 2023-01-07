package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"hanhngo.me/m/common"
)

func GetUserListHandler(c *fiber.Ctx) error {
	var query GetUserListQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid query arguments!"))
	}

	users, err := GetUserListService(ParseGetUserListQuery(query))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(users))
}

func GetUserByIdHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}
	user, err := GetUserByIdService(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(user))
}

func UpdateUserProfileHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}
	var body UpdateUserProfileBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid update body!"))
	}

	updatedUser, err := UpdateUserProfileByIdService(id, body)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(updatedUser))
}

func DeleteUserHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}

	err = DeleteUserByIdService(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Deleted!"))
}
