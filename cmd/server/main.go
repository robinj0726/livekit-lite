package main

import (
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"livekit-lite/pkg/config"
	"livekit-lite/pkg/routing"
	"livekit-lite/pkg/service"

	serverlogger "livekit-lite/pkg/logger"

	"github.com/livekit/protocol/logger"
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

	serverlogger.InitFromConfig(conf.Logging)

	if conf.Development {
		if conf.BindAddresses == nil {
			conf.BindAddresses = []string{
				"127.0.0.1",
				"[::1]",
			}
		}

	}

	return conf, nil
}

func startServer() error {

	conf, err := getConfig()
	if err != nil {
		return err
	}

	currentNode, err := routing.NewLocalNode(conf)
	if err != nil {
		return err
	}

	server, err := service.InitializeServer(conf, currentNode)
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
