package model

import (
	"encoding/json"
	"grpc-server/config"
)

const (
	Gin = config.Gin
)

//获取分页信息
func getLimit(pageNum int, pageSize int) []int {
	limitSlice := make([]int, 2)
	if pageSize == 0 {
		limitSlice[1] = 20
	}
	if pageNum > 0 {
		limitSlice[0] = (pageNum - 1) * pageSize
	}
	return limitSlice
}

//获取排序信息
func getSort(sortStr string) (sort map[string]string) {
	if err := json.Unmarshal([]byte(sortStr), &sort); err != nil {
		return nil
	}
	return
}
