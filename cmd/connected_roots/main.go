package main

import (
	_ "expvar"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kortivex/connected_roots/internal/connected_roots/app"
	_ "go.uber.org/automaxprocs"
)

func main() {
	quit := CloseHandler()
	StartApp(quit)
}

func CloseHandler() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	return quit
}

func StartApp(quit chan os.Signal) {
	go func() {
		app.Start()
		quit <- os.Interrupt
	}()

	<-quit
}
