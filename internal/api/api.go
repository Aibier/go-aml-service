package api

import (
	"net/http"
	"os"
	"time"

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

// Run server
func Run(configPath string) {
	if configPath == "" {
		configPath = "config/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	pid := os.Getpid()
	srvAddr := ":" + conf.Server.Port
	srv := NewServer(srvAddr, web)
	err := web.Run(":" + conf.Server.Port)
	if err != nil {
		log.WithError(err).Logger.Printf("Something went wrong at %s", conf.Server.Port)
	}
	go func() {
		log.Infof("[PID=%d] starting main API server on %s ...", pid, srvAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Error("starting main API server error")
		}
	}()
}

// NewServer creates a new http.Server object
func NewServer(serverAddr string, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:         serverAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 180 * time.Second,
		Handler:      handler,
	}
	return srv
}
