// Package red
package red

import (
	"context"
	"server/internal/consts"
	"server/internal/utils/logger"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx context.Context

func init() {
	ctx := context.Background()
	addr := consts.REDIS_HOST + ":" + consts.REDIS_PORT
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: consts.REDIS_PWD,
		DB:       0,
	})
	Ctx = ctx
	Client = client
	logger.Logger.Info("Redis connected successfully at: " + addr)
}
