package main

import (
	"github.com/majorbruteforce/hifive/internal/config"
	log "github.com/majorbruteforce/hifive/pkg/logger"
)

func main() {
	log.Init()
	defer log.Sync()

	cfg := config.Load() // pass it down as DI to

	log.Info(cfg.Env)
}
