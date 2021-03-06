package dao

import (
	"angrymiao-go/punk/conf/paladin"
	"angrymiao-go/punk/gredis"
)

// Setup Initialize the Redis Client instance
func NewRedisClient() (redisClient *gredis.RedisClient, err error) {
	var (
		cfg gredis.RedisConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	redisPool, err := gredis.NewRedisConn(cfg)
	if err != nil {
		return
	}
	redisClient = &gredis.RedisClient{RedisPool: redisPool}
	return
}