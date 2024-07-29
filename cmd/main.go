package main

import (
	"log"
	"net/http"

	"github.com/majorbruteforce/hi-five/internal/broadcast"
)

func main() {
	cm := broadcast.NewConnetionManager()
	go cm.CreateRandomSession()
	http.HandleFunc("/ws", cm.ServeConnections)
	http.HandleFunc("/debug", cm.Debug)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
