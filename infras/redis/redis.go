package redis

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"hkn-be/config"
	"time"
)

type redisServer struct {
	*config.RedisServer
}

type RedisServerInterface interface {
	Connect(ctx context.Context) (*redis.Client, error)
}

func NewRedisServer(server *config.RedisServer) RedisServerInterface {
	return &redisServer{
		server,
	}
}

func (r redisServer) Connect(ctx context.Context) (*redis.Client, error) {
	timeout := time.Duration(r.Timeout) * time.Second
	rdb := redis.NewClient(
		&redis.Options{
			Addr:        r.Host,
			Password:    r.Password,
			DialTimeout: timeout,
		},
	)

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error("cannot connect to redis")
		return nil, err
	}
	log.Infof("success connect to redis %s", rdb)
	return rdb, nil
}
