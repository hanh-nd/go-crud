package permissions

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"hanhngo.me/m/common"
)

type PermissionHandler struct {
	permissionService PermissionService
}

func NewPermissionHandler(permissionService PermissionService) PermissionHandler {
	return PermissionHandler{
		permissionService: permissionService,
	}
}

func (this *PermissionHandler) CreatePermission(c *fiber.Ctx) error {
	var body CreatePermissionBody
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	permission, err := this.permissionService.CreatePermission(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(permission, fiber.StatusCreated))
}

func (this *PermissionHandler) GetPermissionList(c *fiber.Ctx) error {
	var query GetPermissionListQuery
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	permissions, err := this.permissionService.GetPermissionList(query)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(permissions))
}

func (this *PermissionHandler) GetPermissionById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Permission id must be integer!"))
	}

	permission, err := this.permissionService.GetPermissionById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(permission))
}

func (this *PermissionHandler) UpdatePermission(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Permission id must be integer!"))
	}

	var body UpdatePermissionBody
	if err := c.QueryParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	permission, err := this.permissionService.UpdatePermission(id, body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(permission))
}

func (this *PermissionHandler) DeletePermission(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Permission id must be integer!"))
	}

	err = this.permissionService.DeletePermission(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("OK"))
}
