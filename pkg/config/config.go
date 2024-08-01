package config

import "time"

type SocketConfig struct {
	TimeOut     time.Duration
	RBufferSize int
	WBufferSize int
	Topics      []string
}

type ServerConfig struct {
	Port int
}

type ApplicationConfig struct {
	SocketConfig *SocketConfig
	ServerConfig *ServerConfig
}
