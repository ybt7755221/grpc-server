package model

import (
	"errors"
	. "grpc-server/entity"
	DB "grpc-server/library/database"
)

type GinUsersModel struct {
}

//查找多条数据
func (u *GinUsersModel) Find(ginUsersQuery GinUsersQuery) ([]GinUsers, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	ginUsers := make([]GinUsers, 0)
	//limit
	if ginUsersQuery.PageNum > 0 && ginUsersQuery.PageSize > 0 {
		limitSlic := getLimit(ginUsersQuery.PageNum, ginUsersQuery.PageSize)
		dbConn.Limit(limitSlic[0], limitSlic[1])
	}
	err := dbConn.Find(&ginUsers, ginUsersQuery.Conditions)
	return ginUsers, err
}

//根据id查找单条数据
func (u *GinUsersModel) Get(ginUsers GinUsers) (GinUsers, error) {
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(&ginUsers)
	defer dbConn.Close()
	return ginUsers, err
}

//插入
func (u *GinUsersModel) Insert(ginUsers GinUsers) (GinUsers, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	affected, err := dbConn.Insert(&ginUsers)
	if err != nil {
		return ginUsers, err
	}
	if affected < 1 {
		err = errors.New("插入影响行数: 0")
		return ginUsers, err
	}
	return ginUsers, err
}

//更新
func (u *GinUsersModel) Update(conditions GinUsers, ginUsers GinUsers) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Update(ginUsers, conditions)
	defer dbConn.Close()
	return
}
