package bundlefx

import (
	"context"
	"fmt"

	"github.com/dipeshdulal/fxinit/server/handler"
	"github.com/dipeshdulal/fxinit/server/health"
	"go.uber.org/fx"
)

// Module exports for fx
var Module = fx.Options(
	handler.Module,
	health.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *handler.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting application in :5000")
				go h.Gin.Run(":5000")
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
