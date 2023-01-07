package users

import (
	"math"

	"hanhngo.me/m/common"
	"hanhngo.me/m/database"
	"hanhngo.me/m/model"
)

var UserOmit = []string{"password"}

func GetUserListService(query GetUserListQuery) (common.GetListResponse, error) {
	db := database.DB
	var items []model.User
	var totalItems int64
	page := int(math.Max(common.DEFAULT_PAGE_VALUE, float64(query.Page)))
	limit := query.Limit
	offset := (page - 1) * limit
	err := db.Model(&model.User{}).Limit(limit).Offset(offset).Omit(UserOmit...).Find(&items).Count(&totalItems).Error
	return common.NewGetListResponse(items, totalItems), err
}

func GetUserByIdService(id int) (model.User, error) {
	db := database.DB
	var user model.User
	err := db.Model(&model.User{}).Omit(UserOmit...).First(&user, id).Error
	return user, err
}

func UpdateUserProfileByIdService(id int, body UpdateUserProfileBody) (model.User, error) {
	db := database.DB
	user, err := GetUserByIdService(id)
	if err != nil {
		return user, err
	}

	if body.Email != "" {
		user.Email = body.Email
	}

	db.Save(&user)

	return user, nil
}

func DeleteUserByIdService(id int) error {
	db := database.DB
	user, err := GetUserByIdService(id)

	if err != nil {
		return err
	}

	db.Delete(&user)
	return nil
}
