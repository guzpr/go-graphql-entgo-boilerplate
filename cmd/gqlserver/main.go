package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sekalahita/epirus/internal/ent/gen/runtime"
	"github.com/sekalahita/epirus/internal/gqlserver"
)

const defaultPort = 8080

func main() {
	server := gqlserver.NewServer()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	go func() {
		oscall := <-ch
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	server.Run(ctx, defaultPort)
}
