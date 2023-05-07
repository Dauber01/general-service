package resource

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Config struct {
	Server *ServerConfig `toml:"server"`
	Redis  *RedisConfig  `toml:"redis"`
	Mysql  *MysqlConfig  `toml:"mysql"`
}
type ServerConfig struct {
	Addr         string `toml:"addr"`
	ReadTimeout  int    `toml:"readTimeout"`
	WriteTimeout int    `toml:"writeTimeout"`
}
type RedisConfig struct {
	Ip       string `toml:"ip"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type MysqlConfig struct {
	Ip           string `toml:"ip"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	DB           string `toml:"db"`
	MaxIdleConns int    `toml:"MaxIdleConns"`
	MaxOpenConns int    `toml:"MaxOpenConns"`
}

var RedisClient *redis.Client

var Conf *Config

var MysqlClient *gorm.DB
