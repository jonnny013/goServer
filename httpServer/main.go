package main

import (
	"fmt"
	"log"
	"net/http"
)



type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	port := 8080
	fmt.Printf("Server running on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
