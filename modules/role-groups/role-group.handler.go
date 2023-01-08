package roleGroups

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"hanhngo.me/m/common"
)

type RoleGroupHandler struct {
	roleGroupService RoleGroupService
}

func NewRoleGroupHandler(roleGroupService RoleGroupService) RoleGroupHandler {
	return RoleGroupHandler{
		roleGroupService: roleGroupService,
	}
}

func (handler *RoleGroupHandler) CreateRoleGroup(c *fiber.Ctx) error {
	var body CreateRoleGroupBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	roleGroup, err := handler.roleGroupService.CreateRoleGroup(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(roleGroup, fiber.StatusCreated))
}

func (handler *RoleGroupHandler) GetRoleGroupList(c *fiber.Ctx) error {
	var query GetRoleGroupListQuery
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	roleGroups, err := handler.roleGroupService.GetRoleGroupList(query)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(roleGroups))
}

func (handler *RoleGroupHandler) GetRoleGroupById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "RoleGroup id must be integer!"))
	}

	roleGroup, err := handler.roleGroupService.GetRoleGroupById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(roleGroup))
}

func (handler *RoleGroupHandler) UpdateRoleGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "RoleGroup id must be integer!"))
	}

	var body UpdateRoleGroupBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	roleGroup, err := handler.roleGroupService.UpdateRoleGroup(id, body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(roleGroup))
}

func (handler *RoleGroupHandler) DeleteRoleGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "RoleGroup id must be integer!"))
	}

	err = handler.roleGroupService.DeleteRoleGroup(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("OK"))
}
