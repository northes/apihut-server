package redis

import (
	"fmt"

	"github.com/apihutco/server/config"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() (err error) {
	cfg := config.Conf.Redis
	if !cfg.Enable {
		return nil
	}
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err = client.Ping().Result()
	return err
}
