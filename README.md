# RBAC implementation in Golang

### Local Dev Specification
    - go: v1.18
    - mysql: 8.0.27

### Libs
    - gorm.io/gorm                  
    - gorm.io/driver/mysq           
    - github.com/joho/godotenv      
    - github.com/jaswdr/faker       
    - github.com/golang-jwt/jwt/v4
    - github.com/gin-gonic/gin
    - github.com/google/wire

### Account
- root access
    username: roots
    password: roots
- admin
    username: admins
    password: admins
- client
    common client username and password is same

### Setup documentation di local
- sql schema & dump in ./docs/db.rar
- postmancollection in ./docs/**.json

- install dependency
```go
go mod tidy
```

- running app
```go
go run main.go
```

- build executable
```go
go build main.go
```
