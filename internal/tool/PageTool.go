package tool

import (
	"net/http"
	"strconv"
)

func ParsePageParams(r *http.Request) (int, int) {
	// 页码（默认1，最小1）
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	// 每页条数（默认32，最小1，最大100）
	pageSizeStr := r.URL.Query().Get("page_size")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 32
	}
	if pageSize > 100 {
		pageSize = 100
	}

	return page, pageSize
}
