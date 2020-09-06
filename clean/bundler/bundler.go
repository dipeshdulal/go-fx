package bundler

import (
	"context"
	"fmt"

	"github.com/dipeshdulal/fxinit/clean/controllers"
	"github.com/dipeshdulal/fxinit/clean/routes"
	"github.com/dipeshdulal/fxinit/clean/services"
	"go.uber.org/fx"
)

// Module exported for go-fx depdency injection
var Module = fx.Options(
	controllers.Module,
	services.Module,
	routes.Module,
	fx.Invoke(runApplication),
)

// runApplication runs the gin application
func runApplication(lifecycle fx.Lifecycle, handler routes.Handler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting Application")
			fmt.Println("---------------------")
			fmt.Println("------- CLEAN -------")
			fmt.Println("---------------------")
			go handler.Gin.Run(":5000")
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Stopping Application")
			return nil
		},
	})
}
