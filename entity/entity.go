package entity

import (
	"grpc-server/config"
)

const (
	SUCCESS          = 1000
	ERROR            = 1100
	MISSINGPARAMETER = 1101
	UNKNOWN          = 1200
)

var EN = map[int32]string{
	1000: "Success",
	1100: "Error",
	1101: "Missing Parameter",
	1200: "Unknown Error",
}

var CN = map[int32]string{
	1000: "操作成功",
	1100: "操作失败",
	1101: "缺少必要参数",
	1200: "未知错误",
}

func GetResultInfo(code int32) (str string) {
	switch config.LANG {
	case 1:
		str = CN[code]
	case 2:
		str = EN[code]
	}
	return
}
