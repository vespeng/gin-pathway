package initializer

import (
	"gin-pathway/internal/app/config"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

// InitializeRedis 初始化Redis
func InitializeRedis() error {
	// 创建Redis客户端
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Db,
	})

	// 测试Redis连接
	_, err := RedisClient.Ping(RedisClient.Context()).Result()
	if err != nil {
		log.Errorf("redis初始化失败: %v", err)
		return err
	}

	return nil
}
