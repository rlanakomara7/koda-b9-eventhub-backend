package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Host     string
	Port     string
	Username string
	Password string
	Client   *redis.Client
}

func NewRedisClient(host, port, username, password string) *RedisClient {
	return &RedisClient{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (r *RedisClient) Connect() error {
	addr := fmt.Sprintf("%s:%s", r.Host, r.Port)

	r.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: r.Username,
		Password: r.Password,
		DB:       0,
	})

	ctx := context.Background()

	if err := r.Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect redis:%w", err)
	}
	fmt.Println("Redis connected successfully")
	return nil
}
