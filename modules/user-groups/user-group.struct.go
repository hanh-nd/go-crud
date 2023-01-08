package userGroups

import (
	"math"

	"hanhngo.me/m/common"
)

type GetUserGroupListQuery struct {
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}

func ParseGetUserGroupListQuery(query GetUserGroupListQuery) GetUserGroupListQuery {
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

	return GetUserGroupListQuery{
		Page:           page,
		Limit:          limit,
		OrderBy:        orderBy,
		OrderDirection: orderDirection,
	}
}

type CreateUserGroupBody struct {
	Name string `json:"name"`
}

type UpdateUserGroupBody struct {
	Name string `json:"name"`
}
