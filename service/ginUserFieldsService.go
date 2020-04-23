package service

import (
	"encoding/json"
	"fmt"
	"grpc-server/entity"
	"grpc-server/library/gutil"
	"grpc-server/model"
	gin_user_fields_proto "grpc-server/protos/gin_user_fields"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GinUserFieldsServer struct {
	ginUserFieldsModel *model.GinUserFieldsModel
}

//注册服务
func (s *GinUserFieldsServer) Register(gs *grpc.Server) {
	gin_user_fields_proto.RegisterGinUserFieldsServer(gs, s)
}

//查询服务-带分页
func (s *GinUserFieldsServer) FindByPagination(ctx context.Context, in *gin_user_fields_proto.QuerySchema) (*gin_user_fields_proto.FindRes, error) {
	ginUserFields := new(entity.GinUserFields)
	gutil.BeanUtil(ginUserFields, in.Conditions)
	ginUserFieldsQuery := entity.GinUserFieldsQuery{}
	ginUserFieldsQuery.Conditions = *ginUserFields
	ginUserFieldsQuery.PageNum = int(in.PageNum)
	ginUserFieldsQuery.PageSize = int(in.PageSize)
	fmt.Println(ginUserFieldsQuery)
	ginUserFieldsList, err := s.ginUserFieldsModel.Find(ginUserFieldsQuery)
	result := new(gin_user_fields_proto.FindRes)
	if err != nil {
		result.Code = entity.ERROR
		result.Msg = err.Error()
		result.Data = ""
		return result, err
	}
	byteData, err := json.Marshal(ginUserFieldsList)
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
func (s *GinUserFieldsServer) FindOne(ctx context.Context, in *gin_user_fields_proto.GinUserFieldsSchema) (*gin_user_fields_proto.FindOneRes, error) {
	ginUserFields := new(entity.GinUserFields)
	gutil.BeanUtil(ginUserFields, in)
	ginUserFieldsRes, err := s.ginUserFieldsModel.Get(*ginUserFields)
	gutil.BeanUtil(in, &ginUserFieldsRes)
	res := new(gin_user_fields_proto.FindOneRes)
	res.Code = 1000
	res.Msg = "ok"
	res.Data = in
	return res, err
}

//创建
func (s *GinUserFieldsServer) Create(ctx context.Context, in *gin_user_fields_proto.GinUserFieldsSchema) (res *gin_user_fields_proto.FindOneRes, err error) {
	fmt.Println("CREATE")
	ginUserFields := new(entity.GinUserFields)
	gutil.BeanUtil(ginUserFields, in)
	ginUserFieldsRes, err := s.ginUserFieldsModel.Insert(*ginUserFields)
	res = new(gin_user_fields_proto.FindOneRes)
	if err != nil {
		res.Code = entity.ERROR
		res.Msg = err.Error()
	} else {
		res.Code = 1000
		res.Msg = "success"
	}
	gutil.BeanUtil(in, &ginUserFieldsRes)
	res.Data = in
	return res, err
}

//更新
func (s *GinUserFieldsServer) Update(ctx context.Context, in *gin_user_fields_proto.UpdateSchema) (*gin_user_fields_proto.FindRes, error) {
	updateForm := new(entity.GinUserFieldsUpdateForm)
	gutil.BeanUtil(&updateForm.Conditions, in.Conditions)
	gutil.BeanUtil(&updateForm.Modifies, in.Modifies)
	aff, err := s.ginUserFieldsModel.Update(updateForm.Conditions, updateForm.Modifies)
	result := new(gin_user_fields_proto.FindRes)
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