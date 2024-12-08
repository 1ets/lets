package loader

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/1ets/lets"
)

// List of stop function
var Stopper = []func(){}

// Hold the thread for exitting
func RunningForever() {
	go gracefulShutdown()

	forever := make(chan int)
	<-forever
}

// TODO: Fatal handling
func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		lets.LogD("Shutdown gracefully. ...zzZ")

		OnShutdown()
		os.Exit(0)
	}()
}

func OnShutdown() {
	for _, stop := range Stopper {
		stop()
	}
}
