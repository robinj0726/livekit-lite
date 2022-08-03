package main

import (
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"livekit-lite/pkg/config"
	"livekit-lite/pkg/logger"
	"livekit-lite/pkg/service"
)

func main() {
	startServer()
}

func getConfig() (*config.Config, error) {
	confString, err := getConfigString("./livekit.yaml")
	if err != nil {
		return nil, err
	}

	conf, err := config.NewConfig(confString)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func startServer() error {

	conf, err := getConfig()
	if err != nil {
		return err
	}

	server, err := service.InitializeServer(conf)
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

func getConfigString(configFile string) (string, error) {
	outConfigBody, err := ioutil.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	return string(outConfigBody), nil
}
