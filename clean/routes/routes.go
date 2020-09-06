package routes

import (
	"github.com/dipeshdulal/fxinit/clean/controllers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module exports depdendency
var Module = fx.Options(
	fx.Provide(NewHandler),
	fx.Invoke(registerRoutes),
)

// Handler is request handler
type Handler struct {
	Gin *gin.Engine
}

// NewHandler creates a new handler
func NewHandler() Handler {
	engine := gin.Default()
	return Handler{Gin: engine}
}

func registerRoutes(
	userController controllers.UserController,
	handler Handler,
) {
	userRoutes := handler.Gin.Group("/user")
	{
		userRoutes.GET("/", userController.Get)
		userRoutes.POST("/", userController.Post)
	}
}
