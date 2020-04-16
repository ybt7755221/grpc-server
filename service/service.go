package service

import (
	users_proto "grpc-server/protos/users"
)

func resultString(code int32, msg string, data string) *users_proto.FindRes {
	res := new(users_proto.FindRes)
	res.Code = code
	res.Msg = msg
	res.Data = data
	return res
}
