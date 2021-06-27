# 基本介绍
gobrief是一个基于gin开发的脚手架。本着不造轮子的思想，目前脚手架具备了基本快速开发的功能。提供了多种示例文件，让您把更多时间专注在业务开发上。

# 目录结构
```
➜  gobrief tree .
gobrief
├── README.md
├── app  项目app
│   ├── api    主要接受用户请求层
│   │   ├── dashboard.go
│   │   └── sys_user.go
│   ├── dao    数据处理层
│   │   └── sys_user.go
│   ├── model  模型定义层不含调用方法
│   │   └── sys_user.go
│   └── service 业务处理层
├── go.mod
├── go.sum
├── gobrief   脚手架主目录
│   ├── dbs   数据库驱动
│   │   ├── init.go     初始化数据库驱动
│   │   └── migrate.go  初始化数据库表
│   ├── form_validation  表单校验驱动
│   │   ├── form_validation.go   表单校验的基本方法
│   │   ├── init.go              初始化表单校验驱动
│   │   └── sys_user.go          
│   ├── logger         日志驱动
│   │   └── logger.go  日志基本方法
│   ├── middleware     中间件
│   │   ├── init.go    初始化中间件
│   │   ├── logger.go  日志中间件
│   │   └── recover.go 捕捉致命错误中间件
│   ├── result         统一返回
│   │   ├── result.go  统一返回基本方法
│   │   └── result_code.go  业务码
│   └── router         路由驱动
│       ├── init.go    初始化路由驱动
│       └── v1         v1版路由
│           ├── dashboard.go
│           ├── router.go   v1路由基本方法
│           ├── sys_base.go
│           └── sys_user.go
├── logs               日志
│   └── go-brief.log
├── main.go            启动文件
└── test
    └── router_example
        └── main.go
```

# gobrief思想
我本人致力于不依赖框架开发，架构只是简单的构建出来结构。并不做实际内容的输入，本框架会很轻量。并不会像python里面的Django。
对于框架不同人有不同的想法，有人喜欢使用像Django一样的自带管理的快速开发框架。可我认为框架本身会限制自己很多的想象与实现。
所以，我本人也一直致力于非常轻量的框架。但我认为Gin过于轻量，Gin很适合一个团队用来构建自己的框架体系，但并不是所有团队都
有时间去一点一点研究MVC或MTV等这种架构与目录结构，甚至还要为日志和路由分组浪费时间。所以带着这样的想法gobrief就诞生了。


# 技术选型
- 后端：基于`Gin`快速搭建基础restful风格API，Gin是一个go语言编写的Web框架。
- 数据库： 采用`gorm`实现对数据库的基本操作与驱动。虽然我在项目里使用了gorm，但绝对不建议依赖orm框架！！！！
- 日志：使用`zap`和`lumberjack`记录日志与日志切割。

# 使用说明
- golang版本 >= v1.14.14
- IDE推荐：Goland
- 推荐使用：GoMod

# 计划任务
- [X] 配置管理，区分dev于prod。
- [X] 增加热重启基于`air`，开发过程中使用`air`，生产不要使用。
- [X] 增加命令行gobrief-tool，使用命令直接构建出一个gobrief。项目地址：https://github.com/heyangguang/wisdom-client

# 后期规划
等开发完计划任务以后，我会使用gobrief来开发一个带有`权限管理、用户管理、角色管理、API管理`的参考项目。

1. 可以一起完善学习gobrief的开发思想。
2. 作为example项目一起学习。
3. 带有`权限管理、用户管理、角色管理、API管理`的项目，也可以开箱既用。

# 启动项目
`air  -c .air.toml `