package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := &LightsServer{
		port: 3000,
	}
	server.Init()

	go func() {
		server.Start()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
