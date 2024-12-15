package redis

import (
	"context"
	"fmt"

	log "log/slog"

	"github.com/Abhishekkarunakaran/pbin/src/core/constants"
	"github.com/redis/go-redis/v9"
)

func GetConnection() *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", constants.Env.RedisHost, constants.Env.RedisPort),
			Password: constants.Env.RedisPassword,
			DB:       constants.Int(constants.Env.RedisDB),
		},
	)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Error(fmt.Sprintf("Failed to establish redis connection: %s", err.Error()))
		return nil
	}
	log.Info("Redis connected ...")
	return client
}
