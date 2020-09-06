package main

import (
	"github.com/dipeshdulal/fxinit/clean/bundler"
	"go.uber.org/fx"
)

func main() {
	fx.New(bundler.Module).Run()
}
