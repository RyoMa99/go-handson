package infrastructure

import (
	"handson/domain"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []domain.Player {
	league, _ := domain.NewLeague(f.database)
	return league
}
