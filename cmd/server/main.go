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

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("userID")
		if userID == "" {
			http.Error(w, "missing userID", http.StatusBadRequest)
			return
		}
		sm.HandleWS(w, r, userID)
	})

	http.ListenAndServe(":4018", nil)
}
