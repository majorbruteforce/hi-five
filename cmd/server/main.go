package main

import (
	"net/http"

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

	log.Log.Infoln("Starting server on", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, nil)
}
