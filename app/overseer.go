package app

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"github.com/samber/lo"
)

const _timeout = time.Second * time.Duration(3)

func NewOverseerCfg(env *Env, handler http.Handler) *overseer.Config {
	cfg := &overseer.Config{
		Program: func(state overseer.State) {
			sev := &http.Server{
				ReadHeaderTimeout: _timeout,
				Handler:           handler,
			}

			lo.Must0(sev.Serve(state.Listener))
		},
		Address: fmt.Sprintf("0.0.0.0:%d", env.Port),
	}

	if strings.HasPrefix(strings.ToLower(env.Upgrade), "http") {
		cfg.Fetcher = &fetcher.HTTP{
			URL:      env.Upgrade,
			Interval: time.Minute,
		}
	} else {
		cfg.Fetcher = &fetcher.File{Path: env.Upgrade}
	}

	return cfg
}
