package api

import (
	"github.com/Aibier/go-aml-service/internal/api/router"
	"github.com/Aibier/go-aml-service/internal/pkg/config"
	"github.com/Aibier/go-aml-service/internal/pkg/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
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
			log.WithError( err).Error("starting main API server error")
		}
	}()

	//if conf.Server.Mode != "development" {
	//	go func() {
	//		log.Infof("[PID=%d] starting metrics server on server on %s ...", pid, conf.Server.Port)
	//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//			log.WithError(err).Fatalf("starting metrics server error")
	//		}
	//	}()
	//}
	//
	//term := make(chan os.Signal)
	///* #nosec */
	//signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	//
	//<-term
	//log.Info("received SIGINT/SIGTERM and shutting down gracefully ...")
	//log.Infof("server will shutdown in or before %v", (conf.Server.GracefulShutdown)*time.Second)
	//ctx, cancel := context.WithTimeout(context.Background(), (conf.Server.GracefulShutdown)*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.WithError( err).Error("error shutting down API")
	//}
	//log.Info("exiting")
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