package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page     int    `json:"page" query:"page"`
	PageSize int    `json:"page_size" query:"page_size"`
	Sort     string `json:"sort" query:"sort"`
	Order    string `json:"order" query:"order"`
}

// PaginationResponse represents a paginated response
type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
}

// GetPaginationParams extracts pagination parameters from Echo context
func GetPaginationParams(c echo.Context) PaginationParams {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize <= 0 || pageSize > 100 { // Max 100 items per page
		pageSize = 10
	}

	sort := c.QueryParam("sort")
	if sort == "" {
		sort = "id"
	}

	order := c.QueryParam("order")
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
		Sort:     sort,
		Order:    order,
	}
}

// ApplyPagination applies pagination to a GORM query
func ApplyPagination(db *gorm.DB, params PaginationParams) *gorm.DB {
	offset := (params.Page - 1) * params.PageSize

	// Apply sorting
	orderClause := params.Sort + " " + params.Order

	return db.Order(orderClause).Limit(params.PageSize).Offset(offset)
}

// CreatePaginationResponse creates a paginated response
func CreatePaginationResponse(data interface{}, params PaginationParams, totalItems int64) PaginationResponse {
	totalPages := int((totalItems + int64(params.PageSize) - 1) / int64(params.PageSize))

	return PaginationResponse{
		Data:       data,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
		HasNext:    params.Page < totalPages,
		HasPrev:    params.Page > 1,
	}
}
