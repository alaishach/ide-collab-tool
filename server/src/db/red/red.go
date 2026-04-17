// Package red
package red

import (
	"context"
	"github.com/redis/go-redis/v9"
	"server/src/utils"
)

var Client *redis.Client
var Ctx context.Context

func init() {
	ctx := context.Background()
	addr := utils.REDIS_HOST + ":" + utils.REDIS_PORT
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: utils.REDIS_PWD,
		DB:       0,
	})
	Ctx = ctx
	Client = client
	println("Redis connected successfully at", addr)
}
