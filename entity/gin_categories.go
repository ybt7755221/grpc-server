package entity

type GinCategories struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Catename    string `json:"catename" xorm:"not null VARCHAR(50)"`
	Description string `json:"description" xorm:"not null VARCHAR(255)"`
	CreateTime  string `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdateTime  string `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
