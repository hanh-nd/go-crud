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

func (this *UserHandler) GetUserList(c *fiber.Ctx) error {
	var query GetUserListQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid query arguments!"))
	}

	users, err := this.userService.GetUserList(query)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(users))
}

func (this *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}
	user, err := this.userService.GetUserById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(user))
}

func (this *UserHandler) UpdateUserProfile(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}
	var body UpdateUserProfileBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid update body!"))
	}

	updatedUser, err := this.userService.UpdateUserProfileById(id, body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(updatedUser))
}

func (this *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "User id must be number!"))
	}

	err = this.userService.DeleteUserById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Deleted!"))
}
