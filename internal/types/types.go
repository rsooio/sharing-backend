// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Type   string `json:"type,options=sms|pwd,default=sms"`
	Mobile string `json:"mobile"`
	Secret string `json:"secret"`
}

type LoginResp struct {
	Token     string `json:"token"`
	ExpireAt  int64  `json:"expire"`
	Resources uint64 `json:"res"`
}

type OtpReq struct {
	Mobile string `form:"mobile"`
}

type OtpResp struct {
	Limit int `json:"limit"`
}

type ID struct {
	ID uint `json:"uid",path:"uid"`
}

type Resource struct {
	Role     int    `json:"role,optional"`
	Mobile   string `json:"mobile,optional"`
	Username string `json:"username,optional"`
}

type Detail struct {
	ID
	Resource
}

type Paginator struct {
	PageSize   int `form:"ps"`
	PageNumber int `form:"pn"`
}

type PasswordChange struct {
	ID
	Password    string `json:"pwd"`
	NewPassword string `json:"new_pwd"`
}

type DetailList struct {
	DetailList []Detail `json:"list"`
	Total      int64    `json:"total"`
}
