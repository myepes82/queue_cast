package socket

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"queuecast/pkg/config"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	port          int
	logger        *zap.Logger
	server        *http.Server
	socketHandler *Handler
}

func NewServer(
	config *config.ServerConfig,
	logger *zap.Logger,
	handler *Handler) (*Server, error) {

	logger.Info("Creating new socket server")

	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", config.Port),
		},
		logger:        logger,
		port:          config.Port,
		socketHandler: handler,
	}, nil
}

func (s *Server) Start() error {
	s.logger.Info("Starting socket server", zap.Int("port", s.port))

	http.HandleFunc("/ws", s.socketHandler.HandleSocketConnections)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatal("Failed to start socket server", zap.Error(err))
		}
	}()

	s.logger.Info("Server started")

	<-stop

	s.logger.Info("Shutting down socket server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("Failed to shutdown socket server", zap.Error(err))
	}

	return nil
}
