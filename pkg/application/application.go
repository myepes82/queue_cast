package application

import (
	"go.uber.org/zap"
	"net/http"
	"queuecast/pkg/config"
	"queuecast/pkg/definitions"
	"queuecast/pkg/port/in"
	"queuecast/pkg/socket"
)

type SocketApplication struct {
	//Config
	config *config.ApplicationConfig

	//BroadCast
	broadCast *socket.Broadcast

	//Logger
	logger *zap.Logger

	//Server
	server *socket.Server

	//Services
	compressorService definitions.Compressor

	//Commands
	saveEventUseCase in.SaveEventCommand
}

func NewSocketApplication(
	logger *zap.Logger,
	config *config.ApplicationConfig,
	server *socket.Server,
	compressor definitions.Compressor,
	saveEventUseCase in.SaveEventCommand) *SocketApplication {
	application := &SocketApplication{
		logger:            logger,
		config:            config,
		broadCast:         loadBroadCastChannels(logger, config.GetSocketConfig()),
		server:            server,
		compressorService: compressor,
		saveEventUseCase:  saveEventUseCase,
	}
	return application
}

func loadBroadCastChannels(logger *zap.Logger, config *config.SocketConfig) *socket.Broadcast {
	logger.Info("loading broadcast channels")
	broadCast := &socket.Broadcast{
		Channels: make(map[string]*socket.Channel),
	}
	for _, topic := range config.Topics {
		broadCast.Channels[topic] = socket.NewEmptyChannel(topic)
		logger.Info("broadcasting topic", zap.String("topic", topic))
	}
	logger.Info("broadcast channels loaded")
	return broadCast
}

func (app *SocketApplication) registerSocketHandler() {
	app.logger.Info("registering socket handler")
	socketHandler := socket.NewSocketHandler(
		app.config.GetSocketConfig(),
		app.logger,
		app.broadCast,
		app.compressorService,
		app.saveEventUseCase)

	http.HandleFunc("/ws", socketHandler.HandleSocketConnections)
}

func (app *SocketApplication) Start() error {
	app.logger.Info("starting application")
	app.registerSocketHandler()
	if err := app.server.Start(); err != nil {
		app.logger.Error("failed to start application server", zap.Error(err))
	}
	app.logger.Error("application started")
	return nil
}
