package api

import (
	"github.com/Aibier/go-aml-service/internal/api/router"
	"github.com/Aibier/go-aml-service/internal/pkg/config"
	"github.com/Aibier/go-aml-service/internal/pkg/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "config/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	err := web.Run(":" + conf.Server.Port)
	if err != nil {
		log.WithError(err).Logger.Printf("Something went wrong at %s", conf.Server.Port)
	}
	log.Printf("go aml server is running at %s", conf.Server.Port)
}
