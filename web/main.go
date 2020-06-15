package main

import (
	"flag"

	"go-web/web/conf"
	"go-web/web/library/gateway"
	"go-web/web/service/user"
)

func main(){
	flag.Parse()
	conf.Init()

	apis:=make([]gateway.Api,0)
	apis=append(apis, gateway.Api{
		ID:       1,
		Method:   "GET",
		Path:     "/demo",
		Service:  "UserService",
		Function: "Demo",
		ParamStruct: []string{"*models.BaseReq"},
		BodyType: "JSON",
	},gateway.Api{
		ID:       2,
		Method:   "GET",
		Path:     "/demo2",
		Service:  "UserService",
		Function: "Demo2",
		ParamStruct: []string{"[]models.BaseReq"},
		BodyType: "JSON",
	})
	gateway.Init(apis,user.NewUser()).Run()
}
