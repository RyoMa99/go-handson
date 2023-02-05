package cli

import (
	"bufio"
	"handson/domain"
	"io"
	"strings"
)

func NewCLI(store domain.PlayerStore, in io.Reader) *CLI {
	return &CLI{store, bufio.NewScanner(in)}
}

type CLI struct {
	playerStore domain.PlayerStore
	in          *bufio.Scanner
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
