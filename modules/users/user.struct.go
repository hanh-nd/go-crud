package users

import (
	"math"

	"hanhngo.me/m/common"
)

type GetUserListQuery struct {
	Page           int    `json:"page"`
	Limit          int    `json:"limit"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}

func ParseGetUserListQuery(query GetUserListQuery) GetUserListQuery {
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

	return GetUserListQuery{
		Page:           page,
		Limit:          limit,
		OrderBy:        orderBy,
		OrderDirection: orderDirection,
	}
}

type CreateUserBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserProfileBody struct {
	Email string `json:"email"`
}
