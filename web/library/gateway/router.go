package gateway

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	errno "go-web/web/library/error"
	"go-web/web/models"
)

type Api struct {
	ID       int64
	Method   string
	Path     string
	Service  string
	Function string
	Param    string
	Auth     uint
	Status   int8
}

const (
	// 无需权限
	NoAuth = 0
	// 管理员访问权限
	AdminAuth = 1 << 1
	// 用户访问权限
	UserAuth = 1 << 2
	// 超级管理员访问权限
	SuperAdminAuth = 1 << 3
)

const ()

type Gateway struct {
	apis []Api
	r    *gin.Engine
	smap map[string]interface{}
}

func Init(apis []Api, service ...interface{}) *Gateway {
	g := &Gateway{
		r:    gin.Default(),
		smap: make(map[string]interface{}),
	}

	g.makeServiceMap(service...)

	g.router(apis)

	return g
}

func (g *Gateway) Run(addr ...string) {
	if err := g.r.Run(); err != nil {
		log.Fatal(err)
	}
}

func (g *Gateway) makeServiceMap(service ...interface{}) {
	for _, v := range service {
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		g.smap[t.Name()] = v
	}
}

func (g *Gateway) router(apis []Api) {
	for _, api := range apis {
		switch api.Method {
		case "GET":
			g.r.GET(api.Path, g.Handler(api))
		}
	}
}

func CallBackend(service interface{}, function string, params []reflect.Value) {

	res := reflect.ValueOf(service).MethodByName(function).Call(params)
	fmt.Printf("%+v \n", res)
}

func (g *Gateway) Handler(api Api) func(ctx *gin.Context) {
	var service interface{}
	if value, ok := g.smap[api.Service]; ok {
		service = value
	} else {
		log.Fatalln("api.Service not exist :", api.Service)
	}

	return func(ctx *gin.Context) {
		params := make([]reflect.Value, 0)
		params = append(params, reflect.ValueOf(ctx))

		 param:= models.BaseReq{}
		if err := ctx.ShouldBindQuery(param); err != nil {
			ctx.JSON(http.StatusOK, errno.ParamErr)
			return
		}
		params = append(params, reflect.ValueOf(param))
		CallBackend(service, api.Function, params)
	}
}
