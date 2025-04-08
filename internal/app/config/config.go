package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App      AppConfig      `yaml:"app" mapstructure:"app"`
	Database DatabaseConfig `yaml:"database" mapstructure:"database"`
	Redis    RedisConfig    `yaml:"redis" mapstructure:"redis"`
	Log      LogConfig      `yaml:"log" mapstructure:"log"`
}

type AppConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver" mapstructure:"driver"`
	Source string `yaml:"source" mapstructure:"source"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr" mapstructure:"addr"`
	Password string `yaml:"password" mapstructure:"password"`
	Db       int    `yaml:"db" mapstructure:"db"`
}

type LogConfig struct {
	Format       string `yaml:"format" mapstructure:"format"`
	Level        string `yaml:"level" mapstructure:"level"`
	ReportCaller bool   `yaml:"reportCaller" mapstructure:"reportCaller"`
}

var Conf *Config

// LoadConfig 加载配置文件
func LoadConfig() error {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("加载 .env 文件失败: %v", err)
	}

	// 读取环境变量 APP_ENV
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // 默认环境为 dev
	}

	// 设置配置文件路径和名称
	viper.AddConfigPath("./configs")
	viper.SetConfigName(fmt.Sprintf("config_%s", env))
	viper.SetConfigType("yaml")

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 将配置文件内容解析到 Conf 变量中
	Conf = &Config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	return nil
}
