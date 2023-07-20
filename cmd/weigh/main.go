package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpillora/overseer"
	"github.com/xuender/weigh/app"
)

func main() {
	env := app.NewEnv()

	flag.UintVar(&env.Port, "port", env.Port, "server port")
	flag.StringVar(&env.Cfg, "config", env.Cfg, "config file")
	flag.StringVar(&env.Upgrade, "upgrade", env.Upgrade, "upgrade file")
	flag.Usage = usage
	flag.Parse()

	overseer.SanityCheck()
	overseer.Run(*app.InitCfg(env))
}

func usage() {
	fmt.Fprintf(os.Stderr, "weigh\n\n")
	fmt.Fprintf(os.Stderr, "Simple HTTP proxy library for Go.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
