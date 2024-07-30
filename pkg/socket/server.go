package socket

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"queuecast/pkg/config"
	"queuecast/pkg/core"
	"syscall"
	"time"
)

type Server struct {
	port          int
	logger        *core.Logger
	server        *http.Server
	socketHandler *SocketHandler
}

func NewServer(
	config *config.SocketConfig,
	logger *core.Logger,
	handler *SocketHandler) (*Server, error) {

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
	s.logger.Info(fmt.Sprintf("Starting socket server on port %d", s.port))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.logger.Fatal(fmt.Sprintf("Failed to start socket server: %v", err))
		}
	}()

	s.logger.Info("Server started")

	<-stop

	s.logger.Info("Shutting down socket server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error(fmt.Sprintf("Failed to shutdown socket server: %v", err))
	}

	return nil
}
