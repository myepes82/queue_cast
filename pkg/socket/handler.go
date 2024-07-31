package socket

import (
	"fmt"
	"net/http"
	"queuecast/pkg/core"
	"queuecast/pkg/interfaces"

	"github.com/gorilla/websocket"
)

type SocketHandler struct {
	logger     *core.Logger
	compressor interfaces.Compressor
}

var upgrader = websocket.Upgrader{
	HandshakeTimeout:  1000,
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	Error:             nil,
	CheckOrigin:       func(r *http.Request) bool { return true },
	EnableCompression: true,
}

func NewSocketHandler(logger *core.Logger, compressor interfaces.Compressor) *SocketHandler {
	return &SocketHandler{
		logger:     logger,
		compressor: compressor,
	}
}

func (s *SocketHandler) writeMessage(conn *websocket.Conn, message string) error {
	compressedMessage, err := s.compressor.Compress([]byte(message))

	if err != nil {
		s.logger.Error(fmt.Sprintf("Error decompressing message: %v", err))
		return err
	}

	if err := websocket.Message.Send(ws, compressedMessage); err != nil {
		s.logger.Error(fmt.Sprintf("Error sending message: %v", err))
		break
	}
}

func (s *SocketHandler) readMessage(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			s.logger.Error(fmt.Sprintf("Error receiving message: %v", err))
			return
		}

		s.logger.Info(fmt.Sprintf("message received: %v", messageType))
		s.logger.Info(fmt.Sprintf("message received: %v", string(p)))
	}
}

func (s *SocketHandler) HandleSocketConnections(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error upgrading websocket: %v", err))
		return
	}

	s.logger.Info("New WebSocket connection established")

	s.readMessage(ws)
}
