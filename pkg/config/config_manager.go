package config

import (
	"queuecast/pkg/errors"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ConfigManager struct {
	logger *zap.Logger
	config *ApplicationConfig
}

const (
	//Config vars
	PORT              = "QC_WS_PORT"
	TIMEOUT           = "QC_WS_TIMEOUT"
	READ_BUFFER_SIZE  = "QC_WS_READ_BUFFER_SIZE"
	WRITE_BUFFER_SIZE = "QC_WS_WRITE_BUFFER_SIZE"
	WS_TOPICS         = "QC_WS_TOPICS"
)

func NewConfigManager(logger *zap.Logger) *ConfigManager {
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

	c.getConfigFromEnv()
}

func (c *ConfigManager) GetSocketConfig() *SocketConfig {
	return c.config.SocketConfig
}

func (c *ConfigManager) GetServerConfig() *ServerConfig {
	return c.config.ServerConfig
}

func (c *ConfigManager) loadServerConfig() error {
	if err := c.verifyConfig(
		[]string{PORT}); err != nil {
		c.logger.Error("error loading server configs", zap.Error(err))
		return err
	}

	c.config.ServerConfig = &ServerConfig{
		Port: viper.GetInt(PORT),
	}
	return nil
}

func (c *ConfigManager) loadSocketsConfig() error {

	if err := c.verifyConfig(
		[]string{
			TIMEOUT,
			READ_BUFFER_SIZE,
			WRITE_BUFFER_SIZE,
			WS_TOPICS,
		}); err != nil {
		c.logger.Error("error loading socket configs", zap.Error(err))
		return err
	}

	c.config.SocketConfig = &SocketConfig{
		TimeOut:     viper.GetDuration(TIMEOUT),
		RBufferSize: viper.GetInt(READ_BUFFER_SIZE),
		WBufferSize: viper.GetInt(WRITE_BUFFER_SIZE),
		Topics:      viper.GetStringSlice(WS_TOPICS),
	}

	return nil
}

func (c *ConfigManager) getConfigFromEnv() {

	viper.SetEnvPrefix("QC")
	viper.AutomaticEnv()

	if err := c.loadServerConfig(); err != nil {
		c.logger.Fatal("error loading server configs", zap.Error(err))
	}

	if err := c.loadSocketsConfig(); err != nil {
		c.logger.Fatal("error loading socket configs", zap.Error(err))
	}

}

func (c *ConfigManager) verifyConfig(keys []string) error {
	for _, key := range keys {
		if !c.varIsSet(key) {
			c.logger.Error("config key is missing", zap.String("key", key))
			return errors.ErrConfigMissingValue
		}
	}
	return nil
}

func (c *ConfigManager) varIsSet(key string) bool {
	return viper.IsSet(key)
}
