package auth

import "hanhngo.me/m/modules/users"

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
