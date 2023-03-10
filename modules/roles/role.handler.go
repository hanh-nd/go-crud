package roles

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"hanhngo.me/m/common"
)

type RoleHandler struct {
	roleService RoleService
}

func NewRoleHandler(roleService RoleService) RoleHandler {
	return RoleHandler{
		roleService: roleService,
	}
}

func (handler *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var body CreateRoleBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	role, err := handler.roleService.CreateRole(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(role, fiber.StatusCreated))
}

func (handler *RoleHandler) GetRoleList(c *fiber.Ctx) error {
	var query GetRoleListQuery
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid query!"))
	}

	roles, err := handler.roleService.GetRoleList(query)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(roles))
}

func (handler *RoleHandler) GetRoleById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Role id must be integer!"))
	}

	role, err := handler.roleService.GetRoleById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(role))
}

func (handler *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Role id must be integer!"))
	}

	var body UpdateRoleBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	role, err := handler.roleService.UpdateRole(id, body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(role))
}

func (handler *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Role id must be integer!"))
	}

	err = handler.roleService.DeleteRole(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("OK"))
}

func (handler *RoleHandler) UpdateRolePermissions(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Role id must be integer!"))
	}

	var body UpdateRolePermissionsBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body"))

	}

	role, err := handler.roleService.UpdateRolePermissions(id, body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(role))
}
