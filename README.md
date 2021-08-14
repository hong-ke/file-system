<h1 align="center">Gin Web</h1>

<div align="center">
 基于 GIN + XORM + MySQL + DIG 实现的web脚手架，目的是提供一套轻量的中后台开发框架，方便、快速的完成业务需求的开发。
<br/>

</div>

## 依赖框架

- [Gin](https://gin-gonic.com/) -- The fastest full-featured web framework for Go.
- [XORM](https://github.com/go-xorm/xorm/) -- Xorm is a simple and powerful ORM for Go.
- [Godis](https://github.com/piaohao/godis/) -- redis client implement by golang.
- [Dig](https://github.com/uber-go/dig/) -- A reflection based dependency injection toolkit for Go.
- [Swagger](https://github.com/swaggo/swag/) -- Automatically generate RESTful API documentation with Swagger 2.0 for Go.

## 快速开始

```bash
$ git clone https://github.com/LyricTian/gin-admin

$ cd gin-admin

# 下载依赖
$ go mod tidy
```

> 启动成功之后，可在浏览器中输入地址进行访问：[http://127.0.0.1:8888/swagger/index.html](http://127.0.0.1:8888/swagger/index.html)

## 生成`swagger`文档

```bash
# 基于Makefile
make swag

# OR 使用swag命令
swag init --propertyStrategy pascalcase
```


## 目录结构

```
├── LICENSE
├── Makefile
├── README.md
├── clients
│   └── redis.go
├── config
│   └── config.go
├── config.yml
├── controller
│   └── hello_controller.go
├── dig
│   ├── container.go
│   └── invoke.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── entity
│   └── user.go
├── go.mod
├── go.sum
├── main.go
├── middlewares
│   ├── error.go
│   ├── header.go
│   ├── logger.go
│   └── request_id.go
├── model
├── repository
│   ├── hello_repository.go
│   ├── impl
│   │   ├── hello_repository.go
│   │   └── repository_factory.go
│   ├── repository_factory.go
│   ├── transaction_scope.go
│   └── transaction_scope_with_cancel.go
├── routers
│   ├── api_router.go
│   └── router.go
└── service
    └── hello_service.go
```


## MIT License

    Copyright (c) 2021 hong-ke

