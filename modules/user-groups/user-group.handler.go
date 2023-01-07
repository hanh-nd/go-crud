package userGroups

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"hanhngo.me/m/common"
)

type UserGroupHandler struct {
	userGroupService UserGroupService
}

func NewUserGroupHandler(userGroupService UserGroupService) UserGroupHandler {
	return UserGroupHandler{
		userGroupService: userGroupService,
	}
}

func (this *UserGroupHandler) CreateUserGroup(c *fiber.Ctx) error {
	var body CreateUserGroupBody
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	userGroup, err := this.userGroupService.CreateUserGroup(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(userGroup, fiber.StatusCreated))
}

func (this *UserGroupHandler) GetUserGroupList(c *fiber.Ctx) error {
	var query GetUserGroupListQuery
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	userGroups, err := this.userGroupService.GetUserGroupList(query)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(userGroups))
}

func (this *UserGroupHandler) GetUserGroupById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "UserGroup id must be integer!"))
	}

	userGroup, err := this.userGroupService.GetUserGroupById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(userGroup))
}

func (this *UserGroupHandler) UpdateUserGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "UserGroup id must be integer!"))
	}

	var body UpdateUserGroupBody
	if err := c.QueryParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	userGroup, err := this.userGroupService.UpdateUserGroup(id, body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(userGroup))
}

func (this *UserGroupHandler) DeleteUserGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "UserGroup id must be integer!"))
	}

	err = this.userGroupService.DeleteUserGroup(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("OK"))
}
