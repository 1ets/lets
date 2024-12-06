package loader

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// List of stop function
var Stopper = []func(){}

// Hold the thread for exitting
func RunningForever() {
	go gracefulShutdown()

	forever := make(chan int)
	<-forever
}

// TODO: Create stopper
// TODO: Fatal handling
func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Shutdown gracefully. ...zzZ")
		OnShutdown()
		os.Exit(0)
	}()
}

func OnShutdown() {
	for _, stop := range Stopper {
		stop()
	}
}
