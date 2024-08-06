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

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type ApplicationConfig struct {
	socketConfig   *SocketConfig
	serverConfig   *ServerConfig
	databaseConfig *DatabaseConfig
}

func (cfg *ApplicationConfig) GetSocketConfig() *SocketConfig {
	return cfg.socketConfig
}

func (cfg *ApplicationConfig) GetServerConfig() *ServerConfig {
	return cfg.serverConfig
}

func (cfg *ApplicationConfig) GetDatabaseConfig() *DatabaseConfig {
	return cfg.databaseConfig
}
