package main

import (
	"context"
	"handson/config"
	"handson/user/db"
	"handson/user/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := config.Init()
	if err != nil {
		panic(err)
	}

	userDB, err := db.NewUserDB(context.Background(), config)
	if err != nil {
		panic(err)
	}

	closer := server.Init(userDB)
	defer closer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
