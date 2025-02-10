package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	port := 8080
	fmt.Printf("Server running on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
