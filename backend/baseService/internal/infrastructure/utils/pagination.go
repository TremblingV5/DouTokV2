package utils

import (
	"github.com/cloudzenith/DouTok/backend/baseService/api"
)

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 10
	}

	return size, (page - 1) * size
}

func getPageInfo(total int64, size int32) int32 {
	if total < 0 {
		total = 0
	}

	totalPage := total / int64(size)
	if total%int64(size) != 0 {
		totalPage++
	}

	return int32(totalPage)
}

func GetPageResponse(count int64, page, size int32) *api.PaginationResponse {
	return &api.PaginationResponse{
		Total: getPageInfo(count, size),
		Page:  page,
		Count: int32(count),
	}
}
