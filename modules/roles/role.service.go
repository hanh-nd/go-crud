package roles

import (
	"errors"

	"gorm.io/gorm"

	"hanhngo.me/m/common"
	"hanhngo.me/m/database"
	"hanhngo.me/m/model"
)

type RoleService struct{}

func NewRoleService() RoleService {
	return RoleService{}
}

func (this *RoleService) CreateRole(body CreateRoleBody) (*model.Role, error) {
	db := database.DB
	existedRole, err := this.GetRoleByName(body.Name)

	if err != nil {
		return nil, err
	}

	if existedRole != nil {
		return nil, errors.New("Role existed")
	}

	role := model.Role{
		Name: body.Name,
	}

	if err := db.Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (*RoleService) GetRoleList(query GetRoleListQuery) (common.GetListResponse, error) {
	db := database.DB
	var items []model.Role
	var totalItems int64

	parsedQuery := ParseGetRoleListQuery(query)
	page := parsedQuery.Page
	limit := parsedQuery.Limit
	offset := (page - 1) * limit
	err := db.Model(&model.Role{}).Limit(limit).Offset(offset).Find(&items).Count(&totalItems).Error
	return common.NewGetListResponse(items, totalItems), err
}

func (*RoleService) GetRoleById(id int) (*model.Role, error) {
	db := database.DB

	var role model.Role
	err := db.First(&role, id).Error

	return &role, err
}

func (*RoleService) GetRoleByName(name string) (*model.Role, error) {
	db := database.DB

	var role model.Role
	err := db.Where("name = ?", name).First(&role).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &role, err
}

func (this *RoleService) UpdateRole(id int, body UpdateRoleBody) (*model.Role, error) {
	db := database.DB

	role, err := this.GetRoleById(id)

	if err != nil {
		return nil, err
	}

	role.Name = body.Name
	err = db.Save(&role).Error

	return role, err
}

func (this *RoleService) DeleteRole(id int) error {
	db := database.DB

	role, err := this.GetRoleById(id)

	if err != nil {
		return err
	}

	err = db.Delete(&role).Error

	return err
}
