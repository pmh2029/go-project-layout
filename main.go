package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"engine/config"
	"engine/internal/app/myapi/router"
	"engine/pkg/shared/database"
	sharedLogger "engine/pkg/shared/logger"
)

func main() {
	logger := sharedLogger.NewLogger()
	gin.SetMode(gin.DebugMode)

	config.LoadConfig(logger)

	engine := gin.New()
	router := &router.Router{
		Engine: engine,
		DBConn: config.LoadDB(logger),
	}

	defer database.CloseDB(config.LoadDB(logger), logger)

	router.InitializeRouter(logger)
	router.SetupHandler()

	engine.Run(":" + os.Getenv("API_PORT"))
}
