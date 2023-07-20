//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/jpillora/overseer"
	"github.com/xuender/weigh/proxy"
)

func InitCfg(env *Env) *overseer.Config {
	wire.Build(
		NewOverseerCfg,
		NewConfig,
		proxy.NewService,
		proxy.NewHandler,
	)

	return new(overseer.Config)
}
