# Dynamic Database Demo

This is a demo that uses Golang Interface features to switch database engines without changing business code. Currently supports MySQL, MongoDB.


## Run Locally

Clone the project

```bash
  git clone git@github.com:12ain/dynamic-database-demo.git
```

Go to the project directory

```bash
  cd dynamic-database-demo
```

Install dependencies

```bash
  go mod tidy
```

Run MySQL And MongoDB In Docker

```bash
  docker-compose up -d
```

Start the server

```bash
  go run main.go
```

### Switch Database

Edit application.yaml

```yaml
#   driver: mysql
  driver: mongodb
```
