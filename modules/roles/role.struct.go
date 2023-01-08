package roles

import (
	"math"

	"hanhngo.me/m/common"
)

type GetRoleListQuery struct {
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}

func ParseGetRoleListQuery(query GetRoleListQuery) GetRoleListQuery {
	page := int(math.Max(common.DEFAULT_PAGE_VALUE, float64(query.Page)))
	limit := query.Limit
	if limit == 0 {
		limit = common.DEFAULT_PAGE_LIMIT
	}
	orderBy := query.OrderBy
	if orderBy == "" {
		orderBy = "id"
	}

	orderDirection := query.OrderDirection
	if orderDirection == "" {
		orderDirection = common.DESC
	}

	return GetRoleListQuery{
		Page:           page,
		Limit:          limit,
		OrderBy:        orderBy,
		OrderDirection: orderDirection,
	}
}

type CreateRoleBody struct {
	Name string `json:"name"`
}

type UpdateRoleBody struct {
	Name string `json:"name"`
}

type UpdateRolePermissionsBody struct {
	PermissionIds []int `json:"permissionIds"`
}
