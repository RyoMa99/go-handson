package infrastructure

import (
	"handson/domain"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []domain.Player {
	f.database.Seek(0, 0)
	league, _ := domain.NewLeague(f.database)
	return league
}
