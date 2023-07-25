package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
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

	if oss.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
		logs.Log(logs.SetLogFile("/var/tmp", "proxy.log"))

		defer logs.Close()
	}

	overseer.Run(*app.InitCfg(env))
}

func usage() {
	fmt.Fprintf(os.Stderr, "weigh\n\n")
	fmt.Fprintf(os.Stderr, "Simple HTTP proxy library for Go.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
