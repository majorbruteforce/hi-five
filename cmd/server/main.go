package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/majorbruteforce/hifive/internal/config"
	"github.com/majorbruteforce/hifive/internal/sockets"
	log "github.com/majorbruteforce/hifive/pkg/logger"
)

func main() {
	log.Init()
	defer log.Sync()

	cfg := config.Load()

	sm := sockets.NewSocketManager(cfg)
	go sm.Run()
	sm.RegisterWSHandler()

	go func(sm *sockets.SocketManager) {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			sm.Broadcast(fmt.Appendf(nil, "Broadcast @ %s", t.Format(time.Kitchen)))
		}

		defer ticker.Stop()
	}(sm)

	log.Log.Infoln("Starting server on", cfg.Port)
	http.ListenAndServe(cfg.Port, nil)
}
