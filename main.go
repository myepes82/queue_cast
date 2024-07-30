package main

import (
	"queuecast/pkg/config"
	"queuecast/pkg/core"
	"queuecast/pkg/socket"
)

func main() {

	logger, err := core.NewLogger("info")
	if err != nil {
		panic(err)
	}

	configManager := config.NewConfigManager(logger)

	configManager.InitConfig()

	socketConfig := configManager.GetSocketConfig()

	messageCompressor := core.NewMessageCompressor(logger)

	socketHandler := socket.NewSocketHandler(logger, messageCompressor)

	server, err := socket.NewServer(socketConfig, logger, socketHandler)

	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		return
	}
}
