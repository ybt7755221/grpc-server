package entity

import (
	"time"
)

type GinUsers struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username   string    `json:"username" xorm:"not null default '' comment('用户名') unique VARCHAR(50)"`
	Fullname   string    `json:"fullname" xorm:"not null default '' comment('用户中文名') VARCHAR(50)"`
	Password   string    `json:"password" xorm:"not null default '' comment('密码') CHAR(34)"`
	Mobile     string    `json:"mobile" xorm:"not null default '' comment('手机号') unique CHAR(20)"`
	Email      string    `json:"email" xorm:"not null default 'example@example.com' comment('邮箱') VARCHAR(128)"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	UpdateTime time.Time `json:"update_time" xorm:"updated"`
}

type GinUsersQuery struct {
	Conditions GinUsers `json:"conditions"`
	PageNum    int      `json:"page_num"`
	PageSize   int      `json:"page_size"`
}

type GinUsersUpdateForm struct {
	Conditions GinUsers
	Modifies   GinUsers
}
