package dao

import (
	"angrymiao-go/app/service/main/user/conf"
	"angrymiao-go/punk/gredis"
)

// Setup Initialize the Redis Client instance
func NewRedisClient() (redisClient *gredis.RedisClient, err error) {
	redisPool, err := gredis.NewRedisConn(conf.Conf.Redis)
	if err != nil {
		return
	}
	redisClient = &gredis.RedisClient{RedisPool: redisPool}
	return
}
