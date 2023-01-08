package roleGroups

import (
	"math"

	"hanhngo.me/m/common"
)

type GetRoleGroupListQuery struct {
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
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
	Name string `json:"name"`
}

type UpdateRoleGroupBody struct {
	Name string `json:"name"`
}
