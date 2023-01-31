package main

import (
	"handson/server"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	handler := &server.PlayerServer{Store: &InMemoryPlayerStore{}}

	log.Fatal(http.ListenAndServe(":5000", handler))
}
