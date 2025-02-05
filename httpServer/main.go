package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	port := 8080
	fmt.Printf("Server running on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
