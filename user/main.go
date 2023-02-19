package main

import (
	"handson/user/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	closer := server.Init()
	defer closer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
