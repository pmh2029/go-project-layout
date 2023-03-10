package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"engine/internal/pkg/domains/models/dtos"
)

// Router is application struct
type Router struct {
	Engine *gin.Engine
	DBConn *gorm.DB
	Logger *logrus.Logger
}

// InitializeRouter initializes Engine and middleware
func (r *Router) InitializeRouter(logger *logrus.Logger) {
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())
	r.Engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))
	r.Logger = logger
}

func (r *Router) SetupHandler() {
	// ping
	r.Engine.GET("/ping", func(c *gin.Context) {
		data := dtos.BaseResponse{
			Status: "success",
			Data:   gin.H{"message": "Pong!"},
			Error:  nil,
		}
		c.JSON(http.StatusOK, data)
	})
}
