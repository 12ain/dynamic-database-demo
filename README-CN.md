# 动态数据库演示

这是一个使用 Golang 接口功能的演示，可在不更改业务代码的情况下切换数据库引擎。目前支持 MySQL 和 MongoDB。


## 本地运行

克隆项目

```bash
  git clone git@github.com:12ain/dynamic-database-demo.git
```

转到项目目录

```bash
  cd dynamic-database-demo
```

安装依赖项

```bash
  go mod tidy
```

在 Docker 中运行 MySQL 和 MongoDB

```bash
  docker-compose up -d
```

启动服务器

```bash
  go run main.go
```

### 切换数据库

application.yaml

```yaml
#   driver: mysql
  driver: mongodb
```
