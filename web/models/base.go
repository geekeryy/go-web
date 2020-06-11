package models

type BaseReq struct {
	Name string `form:"name" json:"name"`
}
