package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// redis

// 定义一个全局变量
//var redisdb *redis.Client

var redisdb *redis.Client

var RedisEnabled = true

func InitRedisClient() (err error) {

	redisStr := RedisStr

	if redisStr == "" {
		RedisEnabled = false
		return nil
	}

	opt, err := redis.ParseURL(redisStr)

	if err != nil {
		panic(err)
	}
	redisdb = redis.NewClient(opt)
	//redisdb = redis.NewClient(&redis.Options{
	//	Network:     "tcp",
	//	Addr:        "10.0.0.1:12138",
	//	DB:          1,
	//	DialTimeout: 3 * time.Second,
	//	ReadTimeout: 6 * time.Second,
	//	MaxRetries:  2,
	//	Password:    "123456",
	//})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := redisdb.Ping(ctx).Result()
	fmt.Println(result)

	err = redisdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	return err
}

func Get(key string) string {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	value := redisdb.Get(ctx, key)

	return value.String()
}

func Set(key string, value string) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	redisdb.Set(ctx, key, value, time.Minute*100)
}
