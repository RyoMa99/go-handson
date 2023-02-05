package infrastructure

import (
	"encoding/json"
	"handson/domain"
	"os"
)

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := domain.NewLeague(database)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{database}),
		league:   league,
	}
}

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   domain.League
}

func (f *FileSystemPlayerStore) GetLeague() domain.League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, domain.Player{Name: name, Wins: 1})
	}

	f.database.Encode(f.league)
}
