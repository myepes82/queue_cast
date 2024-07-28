package main

import (
	"queuecast/pkg/config"
	"queuecast/pkg/core"
)

func main() {

	logger, err := core.NewLogger("debug")
	if err != nil {
		panic(err)
	}

	configManager := config.NewConfigManager(logger)

	configManager.InitConfig()

}
