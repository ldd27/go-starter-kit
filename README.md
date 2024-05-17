# go-starter-kit

## 工具安装

```shell
# swagger文档生成工具
go install github.com/swaggo/swag/cmd/swag@latest
# 依赖注入工具
go install github.com/google/wire/cmd/wire@latest
```

## swagger

```shell
# 生成swagger.yaml文档
make docs
# 启动应用程序后通过一下地址访问
# https://petstore.swagger.io/?url=http://localhost:8080/swagger.yaml
```

## 依赖注入

```shell
# 生成依赖注入代码
make wire
```

## 编辑程序

```shell
make build
```

## 启动程序

```shell
# 1.复制conf.toml.example并修改成conf.toml
# 2.修改conf.toml配置
# 3.启动程序
./bin/go-starter-kit -c conf.toml serve 
```