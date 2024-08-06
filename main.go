package main

import (
	"fmt"
	"go.uber.org/zap"
	persistence "queuecast/pkg/adapters/out/persistente"
	"queuecast/pkg/application"
	"queuecast/pkg/config"
	"queuecast/pkg/core/service"
	"queuecast/pkg/core/usecases"
	"queuecast/pkg/socket"
	"queuecast/pkg/utils"
)

func main() {

	logger := utils.GetLoggerFactory()

	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			fmt.Println("failed to sync zap logger")
		}
	}(logger)

	configManager := config.NewConfigManager(logger)

	configManager.InitConfig()

	applicationConfig := configManager.GetConfig()

	serverConfig := applicationConfig.GetServerConfig()
	databaseConfig := applicationConfig.GetDatabaseConfig()

	//Persistence
	database := persistence.NewRedisPersistence(databaseConfig, logger)

	//Services
	messageCompressorService := service.NewMessageCompressor(logger)

	//Adapters - Out
	eventRepository := persistence.NewEventPersistenceAdapter(logger, database)

	//Adapter - In

	//use_cases
	saveEventUseCase := usecases.NewSaveEventUseCase(eventRepository, logger)

	server, err := socket.NewServer(serverConfig, logger)

	if err != nil {
		panic(err)
	}

	newApplication := application.NewSocketApplication(logger, applicationConfig, server, messageCompressorService, saveEventUseCase)

	if err = newApplication.Start(); err != nil {
		return
	}
}
