package roleGroups

import (
	"math"

	"hanhngo.me/m/common"
)

type GetRoleGroupListQuery struct {
	Page           int
	Limit          int
	OrderBy        string
	OrderDirection string
}

func ParseGetRoleGroupListQuery(query GetRoleGroupListQuery) GetRoleGroupListQuery {
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

	return GetRoleGroupListQuery{
		Page:           page,
		Limit:          limit,
		OrderBy:        orderBy,
		OrderDirection: orderDirection,
	}
}

type CreateRoleGroupBody struct {
	Name string
}

type UpdateRoleGroupBody struct {
	Name string
}
