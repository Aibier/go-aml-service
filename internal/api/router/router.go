package router

import (
	"fmt"
	"io"
	"os"

	"github.com/Aibier/go-aml-service/internal/api/controllers"
	"github.com/Aibier/go-aml-service/internal/api/middlewares"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, err := os.Create("log/api.log")
	if err !=nil {
		log.WithError(err).Printf("failed to create log file %s", err)
	}
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())
	// Routes
	// ================== Login Routes
	app.POST("/api/login", controllers.Login)
	app.POST("/api/register", controllers.CreateUser)
	// ================== User Routes

	app.GET("/api/users", middlewares.AuthRequired(), controllers.GetUsers)
	app.GET("/api/users/:id", middlewares.AuthRequired(), controllers.GetUserById)
	app.POST("/api/users", middlewares.AuthRequired(), controllers.CreateUser)
	app.PUT("/api/users/:id", middlewares.AuthRequired(),controllers.UpdateUser)
	app.DELETE("/api/users/:id", middlewares.AuthRequired(), controllers.DeleteUser)
	// ================== Tasks Routes
	app.GET("/api/tasks/:id", middlewares.AuthRequired(), controllers.GetTaskById)
	app.GET("/api/tasks", middlewares.AuthRequired(),controllers.GetTasks)
	app.POST("/api/tasks", middlewares.AuthRequired(), controllers.CreateTask)
	app.PUT("/api/tasks/:id", middlewares.AuthRequired(),controllers.UpdateTask)
	app.DELETE("/api/tasks/:id", middlewares.AuthRequired(),controllers.DeleteTask)

	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
