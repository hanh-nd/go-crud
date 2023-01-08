package userGroups

import (
	"errors"

	"gorm.io/gorm"

	"hanhngo.me/m/common"
	"hanhngo.me/m/database"
	"hanhngo.me/m/model"
)

type UserGroupService struct{}

func NewUserGroupService() UserGroupService {
	return UserGroupService{}
}

func (service *UserGroupService) CreateUserGroup(body CreateUserGroupBody) (*model.UserGroup, error) {
	db := database.DB
	existedUserGroup, err := service.GetUserGroupByName(body.Name)

	if err != nil {
		return nil, err
	}

	if existedUserGroup != nil {
		return nil, errors.New("UserGroup existed")
	}

	userGroup := model.UserGroup{
		Name: body.Name,
	}

	if err := db.Create(&userGroup).Error; err != nil {
		return nil, err
	}

	return &userGroup, nil
}

func (*UserGroupService) GetUserGroupList(query GetUserGroupListQuery) (common.GetListResponse, error) {
	db := database.DB
	var items []model.UserGroup
	var totalItems int64

	parsedQuery := ParseGetUserGroupListQuery(query)
	page := parsedQuery.Page
	limit := parsedQuery.Limit
	offset := (page - 1) * limit
	err := db.Model(&model.UserGroup{}).Limit(limit).Offset(offset).Find(&items).Count(&totalItems).Error
	return common.NewGetListResponse(items, totalItems), err
}

func (*UserGroupService) GetUserGroupById(id int) (*model.UserGroup, error) {
	db := database.DB

	var userGroup model.UserGroup
	err := db.Preload("Manager").Preload("Users").First(&userGroup, id).Error

	return &userGroup, err
}

func (*UserGroupService) GetUserGroupByName(name string) (*model.UserGroup, error) {
	db := database.DB

	var userGroup model.UserGroup
	err := db.Where("name = ?", name).First(&userGroup).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &userGroup, err
}

func (service *UserGroupService) UpdateUserGroup(id int, body UpdateUserGroupBody) (*model.UserGroup, error) {
	db := database.DB

	userGroup, err := service.GetUserGroupById(id)

	if err != nil {
		return nil, err
	}

	userGroup.Name = body.Name
	err = db.Save(&userGroup).Error

	return userGroup, err
}

func (service *UserGroupService) DeleteUserGroup(id int) error {
	db := database.DB

	userGroup, err := service.GetUserGroupById(id)

	if err != nil {
		return err
	}

	err = db.Delete(&userGroup).Error

	return err
}
