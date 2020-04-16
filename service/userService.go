package service

import (
	"encoding/json"
	"fmt"
	"grpc-server/entity"
	"grpc-server/library/gutil"
	"grpc-server/model"
	users_proto "grpc-server/protos/users"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type UserServer struct {
	userModel *model.GinUsersModel
}

//注册服务
func (u *UserServer) Register(s *grpc.Server) {
	users_proto.RegisterGinUsersServer(s, u)
}

//查询服务-带分页
func (u *UserServer) FindUsers(ctx context.Context, in *users_proto.QuerySchema) (*users_proto.FindRes, error) {
	ginUsers := new(entity.GinUsers)
	gutil.BeanUtil(ginUsers, in.Conditions)
	ginUsersQuery := entity.GinUsersQuery{}
	ginUsersQuery.Conditions = *ginUsers
	ginUsersQuery.PageNum = int(in.PageNum)
	ginUsersQuery.PageSize = int(in.PageSize)
	fmt.Println(ginUsersQuery)
	ginUsersList, err := u.userModel.Find(ginUsersQuery)
	if err != nil {
		return resultString(entity.ERROR, err.Error(), ""), err
	}
	byteData, err := json.Marshal(ginUsersList)
	if err != nil {
		return resultString(entity.ERROR, err.Error(), ""), err
	}
	return resultString(entity.SUCCESS, entity.GetResultInfo(entity.SUCCESS), string(byteData)), nil
}

//查询单条
func (u *UserServer) FindOne(ctx context.Context, in *users_proto.GinUsersSchema) (*users_proto.FindOneRes, error) {
	ginUsers := new(entity.GinUsers)
	gutil.BeanUtil(ginUsers, in)
	ginUsersRes, err := u.userModel.Get(*ginUsers)
	gutil.BeanUtil(in, &ginUsersRes)
	res := new(users_proto.FindOneRes)
	res.Code = 1000
	res.Msg = "ok"
	res.Data = in
	return res, err
}

//创建
func (u *UserServer) Create(ctx context.Context, in *users_proto.GinUsersSchema) (res *users_proto.FindOneRes, err error) {
	fmt.Println("CREATE")
	ginUsers := new(entity.GinUsers)
	gutil.BeanUtil(ginUsers, in)
	ginUsersRes, err := u.userModel.Insert(*ginUsers)
	res = new(users_proto.FindOneRes)
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
func (u *UserServer) Update(ctx context.Context, in *users_proto.UpdateSchema) (*users_proto.FindRes, error) {
	updateForm := new(entity.GinUsersUpdateForm)
	gutil.BeanUtil(&updateForm.Conditions, in.Conditions)
	gutil.BeanUtil(&updateForm.Modifies, in.Modifies)
	aff, err := u.userModel.Update(updateForm.Conditions, updateForm.Modifies)
	if err != nil {
		return resultString(entity.ERROR, err.Error(), ""), err
	}
	return resultString(entity.SUCCESS, entity.GetResultInfo(entity.SUCCESS), fmt.Sprintf("affect lines: %d", aff)), err
}
