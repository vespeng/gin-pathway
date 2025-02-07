package initializers

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"vesgo/config"
)

/*
@Author : Vespeng
@Time   : 2025/1/20
@Desc   : redis初始化
*/

var RedisClient *redis.Client

// InitializeRedis 初始化Redis
func InitializeRedis() error {
	// 创建Redis客户端
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Db,
	})

	// 测试数据库连接
	_, err := RedisClient.Ping(RedisClient.Context()).Result()
	if err != nil {
		return fmt.Errorf("redis初始化失败: %v", err)
	}

	return nil
}
