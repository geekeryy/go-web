package gateway

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Api struct {
	ID          int64
	Method      string
	Path        string
	Service     string
	Function    string
	ParamStruct []string
	BodyType    string
	Auth        uint
	Status      int8
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

// 网关
type Gateway struct {
	apis []Api
	r    *gin.Engine
	smap map[string]interface{}     // 服务结构体
	pmap map[string]interface{}     // 参数结构体
	bmap map[string]binding.Binding // 参数绑定映射
}

// 服务初始化
func Init(apis []Api, service ...interface{}) *Gateway {
	g := &Gateway{
		r:    gin.Default(),
		smap: make(map[string]interface{}),
		pmap: make(map[string]interface{}),
		bmap: make(map[string]binding.Binding),
	}

	g.bind()
	g.makeStructMap(service...)

	g.router(apis)

	return g
}

// 启动服务
func (g *Gateway) Run(addr ...string) {
	if err := g.r.Run(); err != nil {
		log.Fatal(err)
	}
}

// 构造结构体映射
func (g *Gateway) makeStructMap(service ...interface{}) {
	for _, value := range service {
		t := reflect.TypeOf(value)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		g.smap[t.Name()] = value

		v := reflect.ValueOf(value)
		for i := 0; i < v.NumMethod(); i++ {
			for j := 1; j < v.Method(i).Type().NumIn(); j++ {
				st := v.Method(i).Type().In(j)
				g.pmap[st.String()] = reflect.New(st).Interface()
			}
		}
	}
}

// 请求参数类型
func (g *Gateway) bind() {
	g.bmap["JSON"] = binding.JSON
	g.bmap["XML"] = binding.XML
	g.bmap["Form"] = binding.Form
	g.bmap["Query"] = binding.Query
	g.bmap["FormPost"] = binding.FormPost
	g.bmap["FormMultipart"] = binding.FormMultipart
	g.bmap["ProtoBuf"] = binding.ProtoBuf
	g.bmap["MsgPack"] = binding.MsgPack
	g.bmap["YAML"] = binding.YAML
	g.bmap["Header"] = binding.Header
}

// 设置路由
func (g *Gateway) router(apis []Api) {
	for _, api := range apis {
		switch api.Method {
		case "GET":
			g.r.GET(api.Path, g.Handler(api))
		case "POST":
			g.r.POST(api.Path, g.Handler(api))
		case "PUT":
			g.r.PUT(api.Path, g.Handler(api))
		case "DELETE":
			g.r.DELETE(api.Path, g.Handler(api))
		case "PATCH":
			g.r.PATCH(api.Path, g.Handler(api))
		case "HEAD":
			g.r.PATCH(api.Path, g.Handler(api))
		case "OPTIONS":
			g.r.OPTIONS(api.Path, g.Handler(api))
		case "Any":
			g.r.Any(api.Path, g.Handler(api))
		}

	}
}

// TODO::服务函数多参数情况
func (g *Gateway) Handler(api Api) func(ctx *gin.Context) {

	var service interface{}
	if value, ok := g.smap[api.Service]; ok {
		service = value
	} else {
		log.Fatalln("api.Service not exist :", api.Service)
	}

	var param interface{}
	if value, ok := g.pmap[api.ParamStruct[0]]; ok {
		param = value
	} else {
		log.Fatalln("api.Struct not exist :", api.ParamStruct)
	}

	var bind binding.Binding
	if value, ok := g.bmap[api.BodyType]; ok {
		bind = value
	} else {
		log.Fatalln("api.BodyType not exist :", api.BodyType)
	}

	return func(ctx *gin.Context) {
		params := make([]reflect.Value, 0)
		params = append(params, reflect.ValueOf(ctx))
		if err := ctx.ShouldBindWith(param, bind); err != nil {
			ctx.JSON(http.StatusOK, err.Error())
			return
		}
		params = append(params, reflect.ValueOf(param).Elem())

		res := reflect.ValueOf(service).MethodByName(api.Function).Call(params)

		if len(res) != 2 {
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		if res[1].Interface().(error).Error() != "" {
			log.Println("err:", res[1].Interface().(error).Error())
		}

		ctx.JSON(http.StatusOK, res[0].Interface())
		return
	}
}
