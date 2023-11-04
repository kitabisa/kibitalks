package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kitabisa/kibitalk/config"
	"time"
)

type ClientRedis struct {
	Client *redis.Client
}

var ClientInstance ICache

type ICache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Ping(ctx context.Context) error
}

func InitCache() {
	ClientInstance = NewCacheClient()
	// Ping the Redis server to check the connection.
	err := ClientInstance.Ping(context.Background())
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}

}

func NewCacheClient() ICache {
	c := config.AppCfg
	// Create a new Redis client.
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Cache.Host, c.Cache.Port), // Replace with your Redis server address.
		Password: "",                                               // No password by default.
		DB:       0,                                                // Default DB.
	})

	return &ClientRedis{
		Client: client,
	}
}

func (c *ClientRedis) Get(ctx context.Context, key string) ([]byte, error) {
	return c.Client.Get(ctx, key).Bytes()
}

func (c *ClientRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	byteRes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Client.Set(ctx, key, byteRes, expiration).Err()
}

func (c *ClientRedis) Ping(ctx context.Context) error {
	pong, err := c.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	fmt.Println("Redis Ping Response:", pong)
	return nil
}
