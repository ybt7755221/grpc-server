package service

import (
	"encoding/json"
	"fmt"
	"grpc-server/entity"
	"grpc-server/library/gutil"
	"grpc-server/model"
	gin_users_proto "grpc-server/protos/gin_users"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GinUsersServer struct {
	ginUsersModel *model.GinUsersModel
}

//注册服务
func (s *GinUsersServer) Register(gs *grpc.Server) {
	gin_users_proto.RegisterGinUsersServer(gs, s)
}

//查询服务-带分页
func (s *GinUsersServer) FindByPagination(ctx context.Context, in *gin_users_proto.QuerySchema) (*gin_users_proto.FindRes, error) {
	ginUsers := new(entity.GinUsers)
	gutil.BeanUtil(ginUsers, in.Conditions)
	ginUsersQuery := entity.GinUsersQuery{}
	ginUsersQuery.Conditions = *ginUsers
	ginUsersQuery.PageNum = int(in.PageNum)
	ginUsersQuery.PageSize = int(in.PageSize)
	fmt.Println(ginUsersQuery)
	ginUsersList, err := s.ginUsersModel.Find(ginUsersQuery)
	result := new(gin_users_proto.FindRes)
	if err != nil {
		result.Code = entity.ERROR
		result.Msg = err.Error()
		result.Data = ""
		return result, err
	}
	byteData, err := json.Marshal(ginUsersList)
	if err != nil {
		result.Code = entity.ERROR
		result.Msg = err.Error()
		result.Data = ""
		return result, err
	}
	result.Code = entity.SUCCESS
	result.Msg = entity.GetResultInfo(entity.SUCCESS)
	result.Data = string(byteData)
	return result, nil
}

//查询单条
func (s *GinUsersServer) FindOne(ctx context.Context, in *gin_users_proto.GinUsersSchema) (*gin_users_proto.FindOneRes, error) {
	ginUsers := new(entity.GinUsers)
	gutil.BeanUtil(ginUsers, in)
	ginUsersRes, err := s.ginUsersModel.Get(*ginUsers)
	gutil.BeanUtil(in, &ginUsersRes)
	res := new(gin_users_proto.FindOneRes)
	res.Code = 1000
	res.Msg = "ok"
	res.Data = in
	return res, err
}

//创建
func (s *GinUsersServer) Create(ctx context.Context, in *gin_users_proto.GinUsersSchema) (res *gin_users_proto.FindOneRes, err error) {
	fmt.Println("CREATE")
	ginUsers := new(entity.GinUsers)
	gutil.BeanUtil(ginUsers, in)
	ginUsersRes, err := s.ginUsersModel.Insert(*ginUsers)
	res = new(gin_users_proto.FindOneRes)
	if err != nil {
		res.Code = entity.ERROR
		res.Msg = err.Error()
	} else {
		res.Code = 1000
		res.Msg = "success"
	}
	gutil.BeanUtil(in, &ginUsersRes)
	res.Data = in
	return res, err
}

//更新
func (s *GinUsersServer) Update(ctx context.Context, in *gin_users_proto.UpdateSchema) (*gin_users_proto.FindRes, error) {
	updateForm := new(entity.GinUsersUpdateForm)
	gutil.BeanUtil(&updateForm.Conditions, in.Conditions)
	gutil.BeanUtil(&updateForm.Modifies, in.Modifies)
	aff, err := s.ginUsersModel.Update(updateForm.Conditions, updateForm.Modifies)
	result := new(gin_users_proto.FindRes)
	if err != nil {
		result.Code = entity.ERROR
		result.Msg = err.Error()
		result.Data = ""
		return result, err
	}
	result.Code = entity.SUCCESS
	result.Msg = entity.GetResultInfo(entity.SUCCESS)
	result.Data = fmt.Sprintf("affect lines: %d", aff)
	return result, err
}