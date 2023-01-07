package auth

import (
	"errors"

	"hanhngo.me/m/modules/users"
	"hanhngo.me/m/plugins/bcrypt"
	"hanhngo.me/m/plugins/jwt"
)

type AuthService struct {
	userService users.UserService
	jwtService  jwt.JwtService
}

func NewAuthService(userService users.UserService, jwtService jwt.JwtService) AuthService {
	return AuthService{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (service *AuthService) Register(body users.CreateUserBody) (string, error) {
	user, err := service.userService.CreateUser(body)

	if err != nil {
		return "", err
	}

	accessToken, err := service.jwtService.Sign(user.ID, user.Username)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (service *AuthService) Login(body LoginBody) (string, error) {
	username := body.Username
	user, err := service.userService.GetUserByUsername(username)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("User not found!")
	}

	if bcrypt.Compare(body.Password, user.Password) != true {
		return "", errors.New("Invalid username or password")
	}

	accessToken, err := service.jwtService.Sign(user.ID, user.Username)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
