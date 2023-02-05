package cli

import (
	"handson/domain"
	"io"
)

type CLI struct {
	playerStore domain.PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Chris")
}
