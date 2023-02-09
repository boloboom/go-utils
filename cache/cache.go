package cache

import (
	"context"
	"fmt"
	"hyperuo/go-utils/config"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var cache *redis.Client

func Init() {
	conf := config.Instance()
	cacheConfig := conf.GetStringMap("cache")
	if len(cacheConfig) == 0 {
		log.Fatal("缓存配置获取失败！")
	}

	var host = cacheConfig["host"].(string)
	var port = cacheConfig["port"].(string)
	var password = cacheConfig["password"].(string)
	var index = cacheConfig["index"].(float64)

	cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       int(index),
		PoolSize: 100,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cache.Ping(ctx).Result()
	if err != nil {
		log.Fatal("缓存初始化失败！")
	}
}

func Instance() *redis.Client {
	return cache
}
