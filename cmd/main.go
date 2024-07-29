package main

import (
	"log"
	"net/http"

	"github.com/majorbruteforce/hi-five/internal/broadcast"
)

func main() {
	cm := broadcast.NewConnetionManager()
	http.HandleFunc("/ws", cm.ServeConnections)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
