package main

import (
	"go-framework/example/internal/di"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// todo flag log trace config

	_, cf, err := di.InitApp()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		// todo log info
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cf()
			// todo log info
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
