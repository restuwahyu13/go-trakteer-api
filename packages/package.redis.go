package packages

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type redisClient struct {
	goRedis *redis.Client
}

func Redis(db int) *redisClient {
	redisOptions := &redis.Options{}
	redisOptions.Addr = fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT"))
	redisOptions.Password = viper.GetString("REDIS_PASSWORD")
	redisOptions.DB = db
	redisOptions.MaxRetries = 15
	redisOptions.PoolFIFO = true
	redisOptions.IdleTimeout = time.Duration(10 * time.Minute)
	redisOptions.DialTimeout = time.Duration(10 * time.Second)
	redisOptions.PoolTimeout = time.Duration(5 * time.Second)
	redisOptions.IdleCheckFrequency = time.Duration(5 * time.Minute)
	redisOptions.ReadTimeout = time.Duration(10 * time.Second)
	redisOptions.WriteTimeout = time.Duration(5 * time.Second)

	if viper.GetString("GO_ENV") == "development" {
		redisOptions.PoolSize = 5
	} else {
		redisOptions.PoolSize = 20
	}

	return &redisClient{goRedis: redis.NewClient(redisOptions)}
}

func (h *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (*redis.StatusCmd, error) {
	res := h.goRedis.SetEX(ctx, key, value, expiration)

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

func (h *redisClient) Get(ctx context.Context, key string) (*redis.StringCmd, error) {
	res := h.goRedis.Get(ctx, key)

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

func (h *redisClient) MSet(ctx context.Context, values map[string]interface{}, expiration time.Duration) (*redis.StatusCmd, error) {
	res := h.goRedis.MSet(ctx, values)

	for i, _ := range values {
		h.goRedis.Expire(ctx, i, expiration)
	}

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

func (h *redisClient) MGet(ctx context.Context, keys ...string) (*redis.SliceCmd, error) {
	res := h.goRedis.MGet(ctx, strings.Join(keys, ","))

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

func (h *redisClient) Hset(ctx context.Context, field string, values map[string]interface{}, expiration time.Duration) (*redis.IntCmd, error) {
	res := h.goRedis.HSet(ctx, field, values)
	h.goRedis.Expire(ctx, field, expiration)

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

func (h *redisClient) HGet(ctx context.Context, key, field string) (*redis.StringCmd, error) {
	res := h.goRedis.HGet(ctx, key, field)

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

func (h *redisClient) HGetAll(ctx context.Context, key string) (*redis.StringStringMapCmd, error) {
	res := h.goRedis.HGetAll(ctx, key)

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

func (h *redisClient) HMset(ctx context.Context, field string, values map[string]interface{}, expiration time.Duration) (*redis.BoolCmd, error) {
	res := h.goRedis.HMSet(ctx, field, values)
	h.goRedis.Expire(ctx, field, expiration)

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}

func (h *redisClient) HMGet(ctx context.Context, key string, fields ...string) (*redis.SliceCmd, error) {
	res := h.goRedis.HMGet(ctx, key, strings.Join(fields, ","))

	if res.Err() != nil {
		defer logrus.Errorf("Redis Error: %v", res.Err())
		return nil, res.Err()
	}

	return res, nil
}
