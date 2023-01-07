package users

import (
	"errors"

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

func (this *UserService) CreateUser(body CreateUserBody) (*model.User, error) {
	db := database.DB
	existedUser, err := this.GetUserByUsername(body.Username)
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

	parsedQuery := ParseGetUserListQuery(query)
	page := parsedQuery.Page
	limit := parsedQuery.Limit
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

func (this *UserService) UpdateUserProfileById(id int, body UpdateUserProfileBody) (*model.User, error) {
	db := database.DB
	user, err := this.GetUserById(id)
	if err != nil {
		return user, err
	}

	if body.Email != "" {
		user.Email = body.Email
	}

	err = db.Save(&user).Error

	return user, err
}

func (this *UserService) DeleteUserById(id int) error {
	db := database.DB
	user, err := this.GetUserById(id)

	if err != nil {
		return err
	}

	err = db.Delete(&user).Error
	return err
}
