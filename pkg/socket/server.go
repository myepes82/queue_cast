package socket

import (
	"fmt"
	"queuecast/pkg/config"
	"queuecast/pkg/core"
)

type Server struct {
	Port   int
	logger *core.Logger
}

func NewServer(config *config.SocketConfig, logger *core.Logger) (*Server, error) {
	return &Server{
		Port:   config.Port,
		logger: logger,
	}, nil
}

func (s *Server) Start() error {
	s.logger.Info(fmt.Sprintf("Starting socket server on port %d", s.Port))
	return nil
}
