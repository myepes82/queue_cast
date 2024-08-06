package persistence

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"queuecast/pkg/config"

	"go.uber.org/zap"
)

type RedisPersistence struct {
	session *redis.Client
	logger  *zap.Logger
}

func NewRedisPersistence(config *config.DatabaseConfig, logger *zap.Logger) *RedisPersistence {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		DB:   0,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		logger.Fatal("Failed to connect to redis", zap.Error(err))
	}
	logger.Info("Connected to redis")

	return &RedisPersistence{
		session: rdb,
		logger:  logger,
	}
}

func (db *RedisPersistence) Close() {
	if err := db.session.Close(); err != nil {
		db.logger.Error("Failed to close redis", zap.Error(err))
	}
}
