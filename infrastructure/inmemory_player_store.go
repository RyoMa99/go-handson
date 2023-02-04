package infrastructure

import (
	"handson/domain"
	"sync"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		Store: map[string]int{},
		lock:  sync.RWMutex{},
	}
}

type InMemoryPlayerStore struct {
	Store map[string]int
	lock  sync.RWMutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.Store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []domain.Player {
	var league []domain.Player
	for name, wins := range i.Store {
		league = append(league, domain.Player{Name: name, Wins: wins})
	}
	return league
}
