package services

import "go.uber.org/fx"

// Module exports depdendency
var Module = fx.Options(
	fx.Provide(NewUserService),
)
