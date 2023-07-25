package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/samber/lo"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/xuender/weigh/app"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	cfg := app.NewConfig(&app.Env{Cfg: "weight.toml"})

	if cfg.LogLevel == 0 {
		cfg.LogLevel = uint32(logs.Info)
	}

	if len(cfg.Serial) == 0 {
		cfg.Serial = []string{"serial"}
	}

	if len(cfg.QPS) == 0 {
		cfg.QPS = map[string]uint32{"qps1": base.Ten, "qps2": base.TwentyFour}
	}

	if len(cfg.Timeout) == 0 {
		cfg.Timeout = map[string]uint32{"timeout1": base.Five, "timeout2": base.Seven}
	}

	encoder := toml.NewEncoder(lo.Must1(os.Create("weigh.toml")))

	lo.Must0(encoder.Encode(cfg))
}

func usage() {
	fmt.Fprintf(os.Stderr, "cfg\n\n")
	fmt.Fprintf(os.Stderr, "Create config file.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
