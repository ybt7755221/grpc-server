package entity

type GinUserFields struct {
	Uid       int    `json:"uid"`
	Sexy      string `json:"sexy"`
	Birthdy   string `json:"birthdy"`
	IdCard    string `json:"id_card"`
	TruthName string `json:"truth_name"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

type GinUserFieldsQuery struct {
	Conditions GinUserFields `json:"conditions"`
	PageNum    int           `json:"page_num"`
	PageSize   int           `json:"page_size"`
}

type GinUserFieldsUpdateForm struct {
	Conditions GinUserFields `json:"conditions"`
	Modifies   GinUserFields `json:"modifies"`
}
