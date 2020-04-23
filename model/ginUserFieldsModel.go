package model

import (
	"errors"
	. "grpc-server/entity"
	DB "grpc-server/library/database"
)

type GinUserFieldsModel struct {
}

//查找多条数据
func (u *GinUserFieldsModel) Find(ginUserFieldsQuery GinUserFieldsQuery) ([]GinUserFields, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	ginUserFields := make([]GinUserFields, 0)
	//limit
	if ginUserFieldsQuery.PageNum > 0 && ginUserFieldsQuery.PageSize > 0 {
		limitSlic := getLimit(ginUserFieldsQuery.PageNum, ginUserFieldsQuery.PageSize)
		dbConn.Limit(limitSlic[0], limitSlic[1])
	}
	err := dbConn.Find(&ginUserFields, ginUserFieldsQuery.Conditions)
	return ginUserFields, err
}

//根据id查找单条数据
func (u *GinUserFieldsModel) Get(ginUserFields GinUserFields) (GinUserFields, error) {
	dbConn := DB.GetDB(Gin)
	_, err := dbConn.Get(&ginUserFields)
	defer dbConn.Close()
	return ginUserFields, err
}

//插入
func (u *GinUserFieldsModel) Insert(ginUserFields GinUserFields) (GinUserFields, error) {
	dbConn := DB.GetDB(Gin)
	defer dbConn.Close()
	affected, err := dbConn.Insert(&ginUserFields)
	if err != nil {
		return ginUserFields, err
	}
	if affected < 1 {
		err = errors.New("插入影响行数: 0")
		return ginUserFields, err
	}
	return ginUserFields, err
}

//更新
func (u *GinUserFieldsModel) Update(conditions GinUserFields, ginUserFields GinUserFields) (affected int64, err error) {
	dbConn := DB.GetDB(Gin)
	affected, err = dbConn.Update(ginUserFields, conditions)
	defer dbConn.Close()
	return
}
