package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Start... (Enter <Ctrl+C> for exit)")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	<-ctx.Done()

	log.Println("Done")
}
