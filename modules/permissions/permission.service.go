package permissions

import (
	"errors"

	"gorm.io/gorm"

	"hanhngo.me/m/common"
	"hanhngo.me/m/database"
	"hanhngo.me/m/model"
)

type PermissionService struct{}

func NewPermissionService() PermissionService {
	return PermissionService{}
}

func (this *PermissionService) CreatePermission(body CreatePermissionBody) (*model.Permission, error) {
	db := database.DB
	existedPermission, err := this.GetPermissionByName(body.Name)

	if err != nil {
		return nil, err
	}

	if existedPermission != nil {
		return nil, errors.New("Permission existed")
	}

	permission := model.Permission{
		Name: body.Name,
	}

	if err := db.Create(&permission).Error; err != nil {
		return nil, err
	}

	return &permission, nil
}

func (*PermissionService) GetPermissionList(query GetPermissionListQuery) (common.GetListResponse, error) {
	db := database.DB
	var items []model.Permission
	var totalItems int64

	parsedQuery := ParseGetPermissionListQuery(query)
	page := parsedQuery.Page
	limit := parsedQuery.Limit
	offset := (page - 1) * limit
	err := db.Model(&model.Permission{}).Limit(limit).Offset(offset).Find(&items).Count(&totalItems).Error
	return common.NewGetListResponse(items, totalItems), err
}

func (*PermissionService) GetPermissionById(id int) (*model.Permission, error) {
	db := database.DB

	var permission model.Permission
	err := db.First(&permission, id).Error

	return &permission, err
}

func (*PermissionService) GetPermissionByName(name string) (*model.Permission, error) {
	db := database.DB

	var permission model.Permission
	err := db.Where("name = ?", name).First(&permission).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &permission, err
}

func (this *PermissionService) UpdatePermission(id int, body UpdatePermissionBody) (*model.Permission, error) {
	db := database.DB

	permission, err := this.GetPermissionById(id)

	if err != nil {
		return nil, err
	}

	permission.Name = body.Name
	err = db.Save(&permission).Error

	return permission, err
}

func (this *PermissionService) DeletePermission(id int) error {
	db := database.DB

	permission, err := this.GetPermissionById(id)

	if err != nil {
		return err
	}

	err = db.Delete(&permission).Error

	return err
}
