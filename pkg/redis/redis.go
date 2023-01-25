package redis

import (
	"context"
	"dwd-api/pkg/setting"
	"time"

	"github.com/go-redis/redis/v8"
)

//实例化一个对象
func NewRedisClient(rdbConf *setting.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdbConf.Host,
		Password: rdbConf.Password,
		DB:       rdbConf.DbIndex,
		PoolSize: 10,
	})
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
