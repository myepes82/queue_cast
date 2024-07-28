package config

import (
	"fmt"
	"queuecast/pkg/core"
	"queuecast/pkg/errors"

	"github.com/spf13/viper"
)

type ConfigManager struct {
	logger *core.Logger
	config *ApplicationConfig
}

const (
	//Config vars
	HOST = "WS_HOST"
	PORT = "WS_PORT"
)

func NewConfigManager(logger *core.Logger) *ConfigManager {
	return &ConfigManager{
		logger: logger,
		config: &ApplicationConfig{},
	}
}

func (c *ConfigManager) InitConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.myapp")
	viper.AddConfigPath("/etc/myapp/")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func (c *ConfigManager) loadSocketsConfig() (*SocketConfig, error) {

	if err := c.verifyConfig(
		[]string{HOST, PORT}); err != nil {
		c.logger.Error(fmt.Sprintf("error loading socket configs: %s", err))
		return nil, err
	}

	return &SocketConfig{
		Host: viper.GetString("SOCKET_HOST"),
		Port: viper.GetInt("SOCKET_PORT"),
	}, nil
}

func (c *ConfigManager) getConfigFromEnv() {
	viper.SetEnvPrefix("QC")
	viper.AutomaticEnv()

	// Socket Configurations
	socketConfig, err := c.loadSocketsConfig()
	if err != nil {
		panic("error loading socket configs")
	}

	c.config.SocketConfig = socketConfig
}

func (c *ConfigManager) verifyConfig(keys []string) error {
	for _, key := range keys {
		if !c.varIsSet(key) {
			c.logger.Error(fmt.Sprintf("config key %s is missing", key))
			return errors.ErrConfigMissingValue
		}
	}
	return nil
}

func (c *ConfigManager) varIsSet(key string) bool {
	return viper.IsSet(key)
}
