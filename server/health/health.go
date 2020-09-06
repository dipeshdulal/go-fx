package health

import (
	"github.com/dipeshdulal/fxinit/server/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func registerRoutes(handler *handler.Handler) {
	handler.Gin.GET("/health", func(c *gin.Context) {
		c.JSON(200, "Health OK")
	})
}

// Module for go fx
var Module = fx.Options(fx.Invoke(registerRoutes))
