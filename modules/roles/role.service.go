package roles

import (
	"errors"

	"gorm.io/gorm"

	"hanhngo.me/m/common"
	"hanhngo.me/m/database"
	"hanhngo.me/m/model"
	"hanhngo.me/m/modules/permissions"
)

type RoleService struct {
	permissionService permissions.PermissionService
}

func NewRoleService(permissionService permissions.PermissionService) RoleService {
	return RoleService{
		permissionService: permissionService,
	}
}

func (service *RoleService) CreateRole(body CreateRoleBody) (*model.Role, error) {
	db := database.DB
	existedRole, err := service.GetRoleByName(body.Name)

	if err != nil {
		return nil, err
	}

	if existedRole != nil {
		return nil, errors.New("role existed")
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
	err := db.Preload("Permissions").First(&role, id).Error

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

func (service *RoleService) UpdateRole(id int, body UpdateRoleBody) (*model.Role, error) {
	db := database.DB

	role, err := service.GetRoleById(id)

	if err != nil {
		return nil, err
	}

	role.Name = body.Name
	err = db.Save(&role).Error

	return role, err
}

func (service *RoleService) DeleteRole(id int) error {
	db := database.DB

	role, err := service.GetRoleById(id)

	if err != nil {
		return err
	}

	err = db.Delete(&role).Error

	return err
}

func (service *RoleService) UpdateRolePermissions(id int, body UpdateRolePermissionsBody) (*model.Role, error) {
	db := database.DB

	permissions, err := service.permissionService.GetPermissionByIds(body.PermissionIds)

	if err != nil {
		return nil, err
	}

	if len(*permissions) < len(body.PermissionIds) {
		return nil, errors.New("some permissions not existed")
	}

	role, err := service.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	toAddPermissions, toDeletePermissions := common.Difference(*permissions, role.Permissions)
	err = db.Transaction(func(tx *gorm.DB) error {
		toDeletePermissionIds := common.Map(toDeletePermissions, func(p model.Permission) uint {
			return p.ID
		})

		toAddRolePermissions := common.Map(toAddPermissions, func(p model.Permission) model.RolePermission {
			return model.RolePermission{
				RoleID:       role.ID,
				PermissionID: p.ID,
			}
		})

		if len(toDeletePermissionIds) > 0 {
			if err := db.Where("role_id = ?", role.ID).Where("permission_id IN ?", toDeletePermissionIds).Delete(&model.RolePermission{}).Error; err != nil {
				return err
			}
		}

		if len(toAddRolePermissions) > 0 {
			if err := db.Create(&toAddRolePermissions).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	updatedRole, err := service.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	return updatedRole, nil
}
