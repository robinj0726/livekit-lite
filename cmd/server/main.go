package main

import (
	"os"
	"os/signal"
	"syscall"

	"livekit-lite/pkg/logger"
	"livekit-lite/pkg/service"
)

func main() {
	startServer()
}

func startServer() error {

	server, err := service.NewLivekitServer()
	if err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan
		logger.Infow("exit requested, shutting down", "signal", sig)
		server.Stop(false)
	}()

	return server.Start()
}
