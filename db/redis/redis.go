package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisLimiter *redis_rate.Limiter
var ctx = context.Background()

func Init() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password:   os.Getenv("REDIS_PASSWORD"),
		DB:         int(db),
		PoolSize:   runtime.NumCPU(),
		MaxRetries: 1,
	})

	_, connectErr := client.Ping(ctx).Result()

	if connectErr != nil {
		log.Panic("Connect Redis failed, err: ", connectErr)
	}

	fmt.Println("Connect Redis success")

	flushErr := client.FlushDB(ctx).Err()

	if flushErr != nil {
		fmt.Println("Error flushing database:", flushErr)
		return
	}

	limiter := redis_rate.NewLimiter(client)

	RedisClient = client
	RedisLimiter = limiter
}

func GetCache(key string) string {
	return RedisClient.Get(ctx, key).Val()
}

func SetCache(key string, value interface{}, expiration time.Duration) bool {
	isSetSuccesfully, err := RedisClient.SetNX(ctx, key, value, expiration).Result()

	if err != nil {
		panic(err)
	}

	return isSetSuccesfully
}
