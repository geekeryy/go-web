package models

type BaseReq struct {
	Name string `form:"name" json:"name"`
	Data []Data `form:"data" json:"data"`
}

type Data struct {
	ID  int64  `form:"id" json:"id"`
	Msg string `form:"msg" json:"msg"`
}
