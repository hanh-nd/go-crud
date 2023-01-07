package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"hanhngo.me/m/common"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) GetUserList(c *fiber.Ctx) error {
	var query GetUserListQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid query arguments!"))
	}

	users, err := handler.userService.GetUserList(ParseGetUserListQuery(query))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(users))
}

func (handler *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}
	user, err := handler.userService.GetUserById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(user))
}

func (handler *UserHandler) UpdateUserProfile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}
	var body UpdateUserProfileBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid update body!"))
	}

	updatedUser, err := handler.userService.UpdateUserProfileById(id, body)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(updatedUser))
}

func (handler *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}

	err = handler.userService.DeleteUserById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(fiber.StatusInternalServerError))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Deleted!"))
}
