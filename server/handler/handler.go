package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Handler type
type Handler struct {
	Gin *gin.Engine
}

func (h *Handler) registerRoutes() {
	h.Gin.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Working": true})
	})
}

// NewHandler returns a new gin router
func NewHandler() *Handler {
	handler := Handler{Gin: gin.Default()}

	handler.registerRoutes()

	return &handler
}

// Module for fx
var Module = fx.Options(fx.Provide(NewHandler))
