package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"go-server/httpServer"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s: %v", dbFileName, err)
	}
	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := poker.NewPlayerServer(store)

	port := 8080
	fmt.Printf("Server running on %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
