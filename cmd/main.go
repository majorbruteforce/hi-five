package main

import (
	"log"
	"net/http"

	"github.com/majorbruteforce/hi-five/internal/broadcast"
	"github.com/majorbruteforce/hi-five/internal/matchmaker"
)

func main() {

	// r := mux.NewRouter()

	matchConfig := matchmaker.Config{
		TargetBufferSize: 4,
		Strategy:         matchmaker.StrategySize,
	}

	mm := matchmaker.NewManager(matchConfig)
	cm := broadcast.NewConnetionManager(mm)

	http.HandleFunc("/ws", cm.ServeConnections)
	http.HandleFunc("/debug", cm.Debug)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
