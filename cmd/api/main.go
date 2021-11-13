package main

import (
	_ "github.com/Aibier/go-aml-service/docs"
	"github.com/Aibier/go-aml-service/internal/api"
)

// @Golang API REST
// @version 1.0
// @description API REST in Golang with Gin Framework

// @contact.name Tony Aizize

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	api.Run("")
}
