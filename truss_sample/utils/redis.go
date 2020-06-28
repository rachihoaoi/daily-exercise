package utils

// docker run -itd --name redis-test -p 6379:6379 redis

import (
	"github.com/go-redis/redis"
	"time"
)

func NewRedisClient(hostPort, password string) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:               hostPort,
		Password:           password,
		DB:                 0,
		MaxRetries:         10,
		IdleTimeout:        time.Second * 30,
		IdleCheckFrequency: time.Second * 10,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return
}
