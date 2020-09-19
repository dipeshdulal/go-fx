package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Handler type
type Handler struct {
	Gin *gin.Engine
}

// NewHandler returns a new gin router
func NewHandler() *Handler {
	handler := Handler{Gin: gin.Default()}
	return &handler
}

// Module for fx
var Module = fx.Options(fx.Provide(NewHandler))
