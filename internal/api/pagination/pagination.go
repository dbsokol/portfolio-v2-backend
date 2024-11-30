package pagination

import (
	"log"
	"net/url"
	"portfolio/internal/api/exception"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PaginatedResponse struct {
	Count   int         `json:"count"`
	Next    *string     `json:"next" nullable:"true"`
	Prev    *string     `json:"prev" nullable:"true"`
	Results interface{} `json:"results"`
}

func GetPaginationParams(c *gin.Context) (limit int, offset int) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		log.Printf("Failed to parse limit: %v", err)
		limit = 10
	}

	offset, err = strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		log.Printf("Failed to parse offset: %v", err)
		offset = 0
	}

	return limit, offset
}

func GetPaginatedResponse(c *gin.Context, limit *int, offset *int, count int, data interface{}) *PaginatedResponse {
	// Apply default values if limit or offset is nil
	limitValue := 10 // Default limit
	offsetValue := 0 // Default offset

	if limit != nil {
		limitValue = *limit
	}

	if offset != nil {
		offsetValue = *offset
	}

	// Validate the calculated values
	if offsetValue > count {
		exception.AddAPIException(
			c,
			exception.APIException{
				Name:   "offset",
				Detail: "offset must be less than the total number of results",
			},
		)
		return nil
	}

	if limitValue <= 0 {
		exception.AddAPIException(
			c,
			exception.APIException{
				Name:   "limit",
				Detail: "limit must be greater than 0",
			},
		)
		return nil
	}

	if offsetValue < 0 {
		exception.AddAPIException(
			c,
			exception.APIException{
				Name:   "offset",
				Detail: "offset must be greater than or equal to 0",
			},
		)
		return nil
	}

	// Generate Next and Prev URLs
	baseURL := c.Request.URL.String()
	next := updatePaginationParams(c, baseURL, limitValue, offsetValue+limitValue, count)
	prev := updatePaginationParams(c, baseURL, limitValue, offsetValue-limitValue, count)

	// Return the paginated response
	return &PaginatedResponse{
		Count:   count,
		Results: data,
		Next:    next,
		Prev:    prev,
	}
}

func updatePaginationParams(c *gin.Context, baseURL string, limit, offset, count int) *string {
	if offset < 0 || offset >= count {
		return nil
	}

	// Parse the base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Printf("Failed to parse URL: %v", err)
		return nil
	}

	// Update query parameters
	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(offset))
	u.RawQuery = q.Encode()

	// Determine the scheme
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	finalURL := decodeURLString(scheme + "://" + c.Request.Host + u.Path + "?" + u.RawQuery)

	return finalURL
}

func decodeURLString(u string) *string {
	/* method to safely decode encoded URL characters */

	result := strings.ReplaceAll(u, "%2C", ",")
	return &result
}
