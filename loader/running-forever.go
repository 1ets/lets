package loader

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

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
		fmt.Println("Sutting down gracefully.")
		// clean up here
		os.Exit(0)
	}()
}
