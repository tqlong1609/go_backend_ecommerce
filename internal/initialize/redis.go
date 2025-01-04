package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/tqlong1609/go_backend_ecommerce/global"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	rd := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rd.Host, rd.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis connection failed", zap.Error(err))
		return
	}
	global.Logger.Info("Redis connection successful")
	global.Rdb = rdb

	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "key123", "value123", 0).Err()
	if err != nil {
		global.Logger.Error("Redis set error", zap.Error(err))
		panic(err)
	}

	value, err := global.Rdb.Get(ctx, "key123").Result()
	if err != nil {
		global.Logger.Error("Redis get error", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Value:", zap.String("Value redis::", value))

}
