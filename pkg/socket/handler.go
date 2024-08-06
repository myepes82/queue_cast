package socket

import (
	"fmt"
	"net/http"
	"queuecast/pkg/config"
	"queuecast/pkg/definitions"
	"queuecast/pkg/port/in"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Handler struct {
	logger           *zap.Logger
	upgrader         *websocket.Upgrader
	compressor       definitions.Compressor
	eventPersistence in.SaveEventCommand
	broadCast        *Broadcast
}

func NewSocketHandler(
	config *config.SocketConfig,
	logger *zap.Logger,
	broadcast *Broadcast,
	compressor definitions.Compressor,
	eventPersistence in.SaveEventCommand) *Handler {
	return &Handler{
		logger: logger,
		upgrader: &websocket.Upgrader{
			HandshakeTimeout: config.TimeOut,
			ReadBufferSize:   config.RBufferSize,
			WriteBufferSize:  config.WBufferSize,
			CheckOrigin:      func(r *http.Request) bool { return true },
		},
		broadCast:        broadcast,
		eventPersistence: eventPersistence,
		compressor:       compressor,
	}
}

func (s *Handler) writeMessage(conn *websocket.Conn, messageType int, message string) error {
	compressedMessage, err := s.compressor.Compress(message)
	if err != nil {
		s.logger.Error("Error compressing message", zap.Error(err))
		return err
	}

	if err := conn.WriteMessage(messageType, compressedMessage); err != nil {
		s.logger.Error("Error sending message", zap.Error(err))
		return err
	}
	return nil
}

func (s *Handler) readMessage(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			s.logger.Error("Error receiving message", zap.Error(err))
			return
		}

		message := string(p)

		s.logger.Info("Message received",
			zap.Int("messageType", messageType),
			zap.String("message", message))

		eventDto := in.SaveEventCommandDto{
			Time:    time.Now(),
			Origin:  conn.RemoteAddr().String(),
			Content: message,
		}
		if err := s.eventPersistence.Execute(eventDto); err != nil {
			s.logger.Error("error at saving input event", zap.Error(err))
		}
	}
}

func (s *Handler) HandleSocketConnections(w http.ResponseWriter, r *http.Request) {

	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error upgrading websocket: %v", err))
		return
	}

	defer func(ws *websocket.Conn) {
		if err := ws.Close(); err != nil {
			s.logger.Error("error closing websocket", zap.Error(err))
		}
	}(ws)

	//Socket connection data
	socketKey := r.Header.Get("Sec-WebSocket-Key")
	topic := r.URL.Query().Get("topic")

	s.logger.Info("client connected", zap.String("client", socketKey))
	s.logger.Info("client topic connected", zap.String("topic", topic))

	err = ws.WriteMessage(1, []byte("connected"))
	if err != nil {
		s.logger.Error("error sending welcome message", zap.Error(err))
	}

	err = ws.WriteMessage(1, []byte(fmt.Sprintf("subscribed: %s", topic)))
	if err != nil {
		s.logger.Error("error sending subscribed message", zap.Error(err))
	}

	s.readMessage(ws)
}
