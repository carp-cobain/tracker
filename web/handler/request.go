package handler

import (
	"strconv"

	"github.com/carp-cobain/tracker/domain"
	"github.com/gin-gonic/gin"
)

// Get and return bounded query parameters for paging.
// If no query params are found, default values are returned.
func getPageParams(c *gin.Context) domain.PageParams {
	cursor, limit := uint64(0), 10
	if cursorQuery, ok := c.GetQuery("cursor"); ok {
		cursor, _ = strconv.ParseUint(cursorQuery, 10, 64)
	}
	if limitQuery, ok := c.GetQuery("limit"); ok {
		limit, _ = strconv.Atoi(limitQuery)
	}
	return domain.NewPageParams(cursor, clamp(limit))
}

// Ensure limit is between 10 and 1000
func clamp(limit int) int {
	if limit >= 10 && limit <= 1000 {
		return limit
	}
	return 10
}
