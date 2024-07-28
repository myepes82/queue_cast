package config

type SocketConfig struct {
	Host string
	Port int
}

type ApplicationConfig struct {
	SocketConfig *SocketConfig
}

var GlobalConfig *ApplicationConfig = &ApplicationConfig{}
