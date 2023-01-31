package main

import (
	"handson/infrastructure"
	"handson/server"
	"log"
	"net/http"
)

func main() {
	handler := &server.PlayerServer{Store: infrastructure.NewInMemoryPlayerStore()}

	log.Fatal(http.ListenAndServe(":5000", handler))
}
