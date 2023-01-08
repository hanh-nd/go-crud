package auth

import (
	"github.com/gofiber/fiber/v2"

	"hanhngo.me/m/common"
	"hanhngo.me/m/modules/users"
)

type AuthHandler struct {
	authService AuthService
	userService users.UserService
}

func NewAuthHandler(authService AuthService, userService users.UserService) AuthHandler {
	return AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (handler *AuthHandler) Register(c *fiber.Ctx) error {
	var body users.CreateUserBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	accessToken, err := handler.authService.Register(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(TokenResponse{
		AccessToken: accessToken,
	}, fiber.StatusCreated))
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	var body LoginBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, "Invalid body!"))
	}

	accessToken, err := handler.authService.Login(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(TokenResponse{
		AccessToken: accessToken,
	}))
}
