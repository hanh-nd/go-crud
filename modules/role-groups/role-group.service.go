package roleGroups

import (
	"errors"

	"gorm.io/gorm"

	"hanhngo.me/m/common"
	"hanhngo.me/m/database"
	"hanhngo.me/m/model"
)

type RoleGroupService struct{}

func NewRoleGroupService() RoleGroupService {
	return RoleGroupService{}
}

func (this *RoleGroupService) CreateRoleGroup(body CreateRoleGroupBody) (*model.RoleGroup, error) {
	db := database.DB
	existedRoleGroup, err := this.GetRoleGroupByName(body.Name)

	if err != nil {
		return nil, err
	}

	if existedRoleGroup != nil {
		return nil, errors.New("RoleGroup existed")
	}

	roleGroup := model.RoleGroup{
		Name: body.Name,
	}

	if err := db.Create(&roleGroup).Error; err != nil {
		return nil, err
	}

	return &roleGroup, nil
}

func (*RoleGroupService) GetRoleGroupList(query GetRoleGroupListQuery) (common.GetListResponse, error) {
	db := database.DB
	var items []model.RoleGroup
	var totalItems int64

	parsedQuery := ParseGetRoleGroupListQuery(query)
	page := parsedQuery.Page
	limit := parsedQuery.Limit
	offset := (page - 1) * limit
	err := db.Model(&model.RoleGroup{}).Limit(limit).Offset(offset).Find(&items).Count(&totalItems).Error
	return common.NewGetListResponse(items, totalItems), err
}

func (*RoleGroupService) GetRoleGroupById(id int) (*model.RoleGroup, error) {
	db := database.DB

	var roleGroup model.RoleGroup
	err := db.First(&roleGroup, id).Error

	return &roleGroup, err
}

func (*RoleGroupService) GetRoleGroupByName(name string) (*model.RoleGroup, error) {
	db := database.DB

	var roleGroup model.RoleGroup
	err := db.Where("name = ?", name).First(&roleGroup).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &roleGroup, err
}

func (this *RoleGroupService) UpdateRoleGroup(id int, body UpdateRoleGroupBody) (*model.RoleGroup, error) {
	db := database.DB

	roleGroup, err := this.GetRoleGroupById(id)

	if err != nil {
		return nil, err
	}

	roleGroup.Name = body.Name
	err = db.Save(&roleGroup).Error

	return roleGroup, err
}

func (this *RoleGroupService) DeleteRoleGroup(id int) error {
	db := database.DB

	roleGroup, err := this.GetRoleGroupById(id)

	if err != nil {
		return err
	}

	err = db.Delete(&roleGroup).Error

	return err
}
