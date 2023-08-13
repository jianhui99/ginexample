package Pagination

import (
	"os"
	"strconv"
)

type PaginationResponse struct {
	Total          int         `json:"total"`
	TotalPageCount int         `json:"total_page_count"`
	NextPage       string      `json:"next_page"`
	Page           int         `json:"page"`
	Limit          int         `json:"limit"`
	Data           interface{} `json:"data"`
}

func GetNextPage(endpoint string, page int, limit int) string {
	endpoint = os.Getenv("APP_DOMAIN") + endpoint
	nextPage := endpoint + "?page=" + strconv.Itoa(page+1) + "&limit=" + strconv.Itoa(limit)
	return nextPage
}
