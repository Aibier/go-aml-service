# go-rest-api-aml-service
Go (Golang) API REST with Gin Framework
## 1. Project Description 
1. Build REST APIs to support AML service with the support of external third party.
2. Setup github actions to run testcases
3. Support microservices
4. JWT authentication
5. Secret Manager


## 2. Run with Docker

1. **Build**

```shell script
make build
docker build . -t api-rest
```

2. **Run**

```shell script
docker-compose up 
or 
docker run -p 3000:3000 api-rest
```

3. **Test**

```shell script
go test -v ./test/...
```

_______

## 3. Generate Docs

```shell script
# Get swag go 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest
# Get swag go bellow 1.16
go get -u github.com/swaggo/swag/cmd/swag
# Generate docs
swag init --dir cmd/api --parseDependency --output docs
or Make sure to import the generated docs/docs.go so that your specific configuration gets init'ed. If your General API annotations do not live in main.go, you can let swag know with -g flag.
swag init -g http/api.go 
```

## 3. Docs and References
1. Swaggo official docs **https://github.com/swaggo/swag**
2. Run and go to **http://localhost:3000/docs/index.html**
