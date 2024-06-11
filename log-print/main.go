package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "log-print", log.LstdFlags)

	logger.Println("Log print is starting...")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-ctx.Done():
			logger.Println("Log print is stopping...")
			return
		case <-ticker.C:
			logger.Printf("blob blob vol. %d\n", i)
			i++
		}
	}
}
