package main

import (
	"queuecast/pkg/config"
	"queuecast/pkg/core"
	"queuecast/pkg/socket"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	configManager := config.NewConfigManager(logger)

	configManager.InitConfig()

	serverConfig := configManager.GetServerConfig()
	socketConfig := configManager.GetSocketConfig()

	messageCompressor := core.NewMessageCompressor(logger)

	socketHandler := socket.NewSocketHandler(
		socketConfig,
		logger,
		messageCompressor)

	server, err := socket.NewServer(serverConfig, logger, socketHandler)

	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		return
	}
}
