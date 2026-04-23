// Package red
package red

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"server/internal/consts"
	"server/internal/db/pg"
	"server/internal/err/panics"
	"server/internal/utils/logger"
	"time"
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

func GetSession(sessionToken string) *pg.SessionData {
	val, err := Client.Get(Ctx, sessionToken).Result()
	// if key is not found redis return redis.Nil
	if err == redis.Nil {
		logger.Logger.Debug("Session Token not found: ", "sessionToken", sessionToken)
		return nil
	}
	if err != nil {
		panics.PanicRedis("GetSession getting from redis", err)
	}
	var session pg.SessionData
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		panics.PanicRedis("Get Session Unmarshalling", err)
	}
	return &session
}
