package socket

import (
	"fmt"
	"log"
	"queuecast/pkg/core"
	"queuecast/pkg/interfaces"

	"golang.org/x/net/websocket"
)

type SocketHandler struct {
	logger     *core.Logger
	compressor interfaces.Compressor
}

func NewSocketHandler(logger *core.Logger, compressor interfaces.Compressor) *SocketHandler {
	return &SocketHandler{
		logger:     logger,
		compressor: compressor,
	}
}

func (s *SocketHandler) HandleSocketConnections(ws *websocket.Conn) {
	s.logger.Info("New WebSocket connection established")

	defer func() {
		s.logger.Info("WebSocket connection closed")
		err := ws.Close()
		if err != nil {
			return
		}
	}()

	for {
		var compressedMessage []byte
		if err := websocket.Message.Receive(ws, &compressedMessage); err != nil {
			s.logger.Error(fmt.Sprintf("Error receiving message: %v", err))
			break
		}

		message, err := s.compressor.Compress(compressedMessage)
		if err != nil {
			s.logger.Error(fmt.Sprintf("Error decompressing message: %v", err))
			break
		}
		s.logger.Info(fmt.Sprintf("Received message: %v", string(message)))
		log.Println("Received message:", string(message))

		compressedMessage, err = s.compressor.Compress(message)
		if err != nil {
			s.logger.Error(fmt.Sprintf("Error compressing message: %v", err))
			break
		}

		if err := websocket.Message.Send(ws, compressedMessage); err != nil {
			s.logger.Error(fmt.Sprintf("Error sending message: %v", err))
			break
		}
	}
}
