package main

import (
	"fmt"
	"handson/cli"
	"handson/infrastructure"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := infrastructure.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	cli.NewCLI(store, os.Stdin, os.Stdout, cli.BlindAlerterFunc(cli.StdOutAlerter)).PlayPoker()
}
