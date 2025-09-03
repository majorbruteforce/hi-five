package main

import (
	log "github.com/majorbruteforce/hi-five/pkg/logger"
)

func main() {
	log.Init()
	defer log.Sync()
	
	log.Info("Live!")
}
