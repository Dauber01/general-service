package bootstrap

import (
	"context"
	"general-service/library/resource"

	"github.com/redis/go-redis/v9"
)

// 初始化 resource 的函数
func ResourceInit(ctx context.Context) {

	// 初始化 redis client
	initRedis(ctx)

}

// initRedis 初始化 Redis Client
// https://github.com/redis/go-redis
func initRedis(_ context.Context) {

	// 初始化出来一个 redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password", // no password set
		DB:       0,          // use default DB
	})

	// client 初始化为单例，放到 resource 里去
	resource.RedisClient = rdb
}
