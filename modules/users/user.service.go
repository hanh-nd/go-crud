package users

import (
	"errors"
	"math"

	"gorm.io/gorm"

	"hanhngo.me/m/common"
	"hanhngo.me/m/database"
	"hanhngo.me/m/model"
	"hanhngo.me/m/plugins/bcrypt"
)

type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

var UserOmit = []string{"password"}

func (service *UserService) CreateUser(body CreateUserBody) (*model.User, error) {
	db := database.DB
	existedUser, err := service.GetUserByUsername(body.Username)
	if err != nil {
		return nil, err
	}

	if existedUser != nil {
		return nil, errors.New("User existed!")
	}

	hashedPassword, err := bcrypt.Hash(body.Password)

	if err != nil {
		return nil, err
	}

	user := model.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
	}

	if err = db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (*UserService) GetUserList(query GetUserListQuery) (common.GetListResponse, error) {
	db := database.DB
	var items []model.User
	var totalItems int64
	page := int(math.Max(common.DEFAULT_PAGE_VALUE, float64(query.Page)))
	limit := query.Limit
	offset := (page - 1) * limit
	err := db.Model(&model.User{}).Limit(limit).Offset(offset).Omit(UserOmit...).Find(&items).Count(&totalItems).Error
	return common.NewGetListResponse(items, totalItems), err
}

func (*UserService) GetUserById(id int) (*model.User, error) {
	db := database.DB
	var user model.User
	err := db.Model(&model.User{}).Omit(UserOmit...).First(&user, id).Error
	return &user, err
}

func (*UserService) GetUserByUsername(username string) (*model.User, error) {
	db := database.DB
	var user model.User
	err := db.Model(&model.User{}).Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (service *UserService) UpdateUserProfileById(id int, body UpdateUserProfileBody) (*model.User, error) {
	db := database.DB
	user, err := service.GetUserById(id)
	if err != nil {
		return user, err
	}

	if body.Email != "" {
		user.Email = body.Email
	}

	db.Save(&user)

	return user, nil
}

func (service *UserService) DeleteUserById(id int) error {
	db := database.DB
	user, err := service.GetUserById(id)

	if err != nil {
		return err
	}

	db.Delete(&user)
	return nil
}
