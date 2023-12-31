package app

import (
	"github.com/BurntSushi/toml"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/xuender/weigh/pb"
)

type Env struct {
	Port    uint
	Cfg     string
	Upgrade string
}

func NewEnv() *Env {
	return &Env{8080, "weigh.toml", "upgrade"}
}

func NewConfig(env *Env) *pb.Config {
	cfg := &pb.Config{}

	if _, err := toml.DecodeFile(env.Cfg, cfg); err != nil {
		logs.SetLevel(logs.Info)
		logs.E.Println(env.Cfg, err)
	} else {
		logs.SetLevel(logs.Level(cfg.LogLevel))
	}

	if cfg.PoolSize == 0 {
		cfg.PoolSize = base.Kilo * base.Ten
	}

	if cfg.TimeoutSecond < 1 {
		cfg.TimeoutSecond = 300
	}

	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = base.Kilo
	}

	if cfg.MaxIdleConnsPerHost == 0 {
		cfg.MaxIdleConnsPerHost = base.Hundred
	}

	logs.D.Println("pool size:", cfg.PoolSize)

	return cfg
}
