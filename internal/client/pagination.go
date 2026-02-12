package client

import (
	"net/http"
	"strconv"
)

type PaginationInfo struct {
	Page       int `json:"page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	PerPage    int `json:"per_page"`
}

func parsePaginationHeaders(resp *http.Response) *PaginationInfo {
	info := &PaginationInfo{}
	info.TotalPages, _ = strconv.Atoi(resp.Header.Get("X-Pagination-Total-Pages"))
	info.TotalItems, _ = strconv.Atoi(resp.Header.Get("X-Pagination-Result-Count"))
	info.Page, _ = strconv.Atoi(resp.Header.Get("X-Pagination-Page"))
	info.PerPage, _ = strconv.Atoi(resp.Header.Get("X-Pagination-Limit"))
	return info
}
