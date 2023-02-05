package cli

import "handson/domain"

type CLI struct {
	playerStore domain.PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Cleo")
}
