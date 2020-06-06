# go-web

## 项目规范


### git提交规范
- feat：新功能（feature）
- fix：修补bug
 -docs：文档（documentation）
- style： 格式（不影响代码运行的变动）
- refactor：重构（即不是新增功能，也不是修改bug的代码变动）
- test：增加测试
- chore：构建过程或辅助工具的变动

## 目录结构

```
.
├── web
│   ├──conf: 配置
│   ├──dao: 数据访问层
│   ├──library: 公共库
│   ├──models: 数据模型
│   ├──server
│   │   └── http: 提供Restful api
│   └───service: 提供服务
│
└── vendor: 项目所需外部包
```