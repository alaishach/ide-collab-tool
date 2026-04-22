// Package red
package red

import (
	"context"
	"encoding/json"
	"server/internal/consts"
	"server/internal/db/pg"
	"server/internal/err/panics"
	"server/internal/utils/logger"
	"time"

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
	Client.Options().MaxConcurrentDials = 20
	logger.Logger.Info("Redis connected successfully at: " + addr)
}

func AddSession(session pg.SessionData) {
	data, e := json.Marshal(session)
	panics.PanicRedis("AddSession sessionData won't Marshal", e)
	err := Client.Set(Ctx, session.SessionToken.String(), data, time.Minute*10).Err()
	panics.PanicRedis("AddSession", err)
}

func GetSession(sessionToken string) *int {
	id, err := Client.Get(Ctx, sessionToken).Result()
	println("!!!!! GetSession: ", id)
	println("!!!!! GetSession: ", err)
	if err != nil {
		println("!!!!! GetSession error: ", err.Error())
		return nil
	}
	i := 0
	return &i
}
