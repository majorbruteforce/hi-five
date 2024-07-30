package main

import (
	"fmt"

	"github.com/majorbruteforce/hi-five/pkg/graph"
)

// import (
// 	"log"
// 	"net/http"

// 	"github.com/majorbruteforce/hi-five/internal/broadcast"
// )

// func main() {
// 	cm := broadcast.NewConnetionManager()
// 	http.HandleFunc("/ws", cm.ServeConnections)
// 	http.HandleFunc("/debug", cm.Debug)

// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

func main() {
	type Edge = graph.Edge
	edges := []Edge{
		{Node1: 1, Node2: 2, Weight: 2},
		{Node1: 1, Node2: 3, Weight: 3},
		{Node1: 1, Node2: 4, Weight: 3},
		{Node1: 1, Node2: 5, Weight: 2},
		{Node1: 2, Node2: 4, Weight: 4},
		{Node1: 2, Node2: 5, Weight: 2},
		{Node1: 3, Node2: 4, Weight: 1},
		{Node1: 3, Node2: 5, Weight: 3},
		{Node1: 4, Node2: 5, Weight: 1},
	}

	result, err := graph.MaxWeightedMatching(edges)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
