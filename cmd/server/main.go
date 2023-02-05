package main

import (
	"handson/infrastructure"
	"handson/server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := infrastructure.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := server.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
