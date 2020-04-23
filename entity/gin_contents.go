package entity

type GinContents struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Topic      string `json:"topic" xorm:"not null comment('主题') VARCHAR(255)"`
	Content    string `json:"content" xorm:"not null comment('详细内容') TINYTEXT"`
	Category   int    `json:"category" xorm:"not null default 0 comment('分类') TINYINT(3)"`
	TestTime   string `json:"test_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('测试时间') DATETIME"`
	PulishTime string `json:"pulish_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('上线时间') DATETIME"`
	OpTime     string `json:"op_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	JiraUrl    string `json:"jira_url" xorm:"not null default '' comment('jira地址') VARCHAR(255)"`
}
