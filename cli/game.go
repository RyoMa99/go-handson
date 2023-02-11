package cli

import (
	"handson/domain"
	"time"
)

func NewGame(alerter BlindAlerter, store domain.PlayerStore) *game {
	return &game{
		alerter: alerter,
		store:   store,
	}
}

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type game struct {
	alerter BlindAlerter
	store   domain.PlayerStore
}

func (p *game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (p *game) Finish(winner string) {
	p.store.RecordWin(winner)
}
