# GO-WEB 业务框架

## TODO
- [ ] 应用程序上下文
- [ ] 数据访问层
- [ ] 定时任务
- [ ] 工具库

## Git Flow 分支管理
- master：主分支（用于版本发布，始终与线上一致）
- dev：开发分支（用于开发，提测时，从dev检出release-1.0.0分支）
- release: 预发布（用于测试，测试中有问题直接修改，测试完成后合并入master和dev）
- feature-*：功能分支（用于功能开发，完成后合并到dev）
- hotfix-*：修复bug（从master分出来，完成后合并到master和dev）
![](http://assets.processon.com/chart_image/5f93a2e15653bb06ef13def8.png)

## 项目规范

### go文件模板（IDEA）
```
// @Description  TODO
// @Author  	 ${USER}  
// @Created  	 ${DATE} ${TIME}
package ${GO_PACKAGE_NAME}

#if (${GO_PACKAGE_NAME}=="controllers" || ${GO_PACKAGE_NAME}=="services" || ${GO_PACKAGE_NAME}=="dao" )
	
	#set($file_name = ${StringUtils.removeAndHump($NAME)})
	#set($package_name = ${StringUtils.removeAndHump(${GO_PACKAGE_NAME})})
	
	#set($end = $package_name.length()+(-1))
	#set($package_name = $package_name.substring(0,$end))
	
	#if (${GO_PACKAGE_NAME}=="dao")
		#set ($package_name = "Repo")
	#end
	
	type ${file_name}${package_name} struct{
		// TODO
	}
	
	func New${file_name}${package_name}() *${file_name}${package_name} {
		return &${file_name}${package_name}{}
	}

#end 

```

### API规范 
[Restful](http://kaelzhang81.github.io/2019/05/24/Restful-API%E8%AE%BE%E8%AE%A1%E6%9C%80%E4%BD%B3%E5%AE%9E%E8%B7%B5/)
- URL
```
一类资源两个URL

# 资源集合：
/epics
# 资源元素：
/epics/5

定制有效载荷大小
GET /stories?include=subTasks
``` 
- HTTP Method
```
增（POST：非幂等性）: 使用POST方法创建新的资源。
删（DELETE：幂等性）: 使用DELETE方法删除存在的资源。
改（PUT：幂等性）: 使用PUT或PATCH方法来更新已存在的资源。
查: 使用GET方法读取资源。（GET：幂等性）
```
- HTTP 状态码
``` 
200: Success 成功
304: Not Modified 未修改
400: Bad Request 客户端请求语法错误
401: Unauthorized 未授权
403: Forbidden 禁止访问
404: Not Found 资源无法找到
429: Too Many Requests 客户端的请求次数超过限额
500: Internal Server Error 服务器错误
503: Service Unavailable 系统维护
```

- 返回值结构
错误示例:
```json
{
  "code": 1001,
  "msg": "数据更新失败"
}
```
成功示例：
```json
{
  "code": 0,
  "msg": "ok",
  "data": {

  }
}
```

### Git提交规范
- feat：新功能（feature）
- fix：修补bug
- doc:：文档（documentation）
- style： 格式（不影响代码运行的变动）
- refactor：重构（即不是新增功能，也不是修改bug的代码变动）
- test：增加测试
- chore：构建过程或辅助工具的变动

## 项目结构

```
.
├── cmd: 服务命令
├── dao: 数据访问层
├── models: 数据模型
├── router: 路由
├── middlewares: http中间件
├── services: 业务逻辑层
├── util: 工具包
│   ├── database: 数据库
│   ├── jwt: 令牌
│   ├── log: 日志
│   └── email: 邮件
└── vendor: 项目所需外部包
```

## 项目分层图
![](http://assets.processon.com/chart_image/5ee9840fe0b34d4dba40cfb7.png)

## License
© JiangYang, 2020~time.Now

Released under the Apache [License](https://github.com/comeonjy/go-web/blob/master/LICENSE)