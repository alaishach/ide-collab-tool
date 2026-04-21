// Package red
package red

import (
	"context"
	"github.com/redis/go-redis/v9"
	"server/internal/consts"
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
	println("Redis connected successfully at", addr)
}
