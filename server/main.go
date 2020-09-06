package main

import (
	"github.com/dipeshdulal/fxinit/server/bundlefx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundlefx.Module,
	).Run()
}
