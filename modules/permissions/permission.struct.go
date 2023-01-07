package permissions

import (
	"math"

	"hanhngo.me/m/common"
)

type GetPermissionListQuery struct {
	Page           int
	Limit          int
	OrderBy        string
	OrderDirection string
}

func ParseGetPermissionListQuery(query GetPermissionListQuery) GetPermissionListQuery {
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

	return GetPermissionListQuery{
		Page:           page,
		Limit:          limit,
		OrderBy:        orderBy,
		OrderDirection: orderDirection,
	}
}

type CreatePermissionBody struct {
	Name string
}

type UpdatePermissionBody struct {
	Name string
}
