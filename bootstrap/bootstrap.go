package bootstrap

import (
	"context"
	"fmt"
	"general-service/library/resource"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化 resource 的函数
func ResourceInit(ctx context.Context) {

	// 初始化 redis client
	initRedis(ctx)

	// 初始化 mysql client
	initMysql(ctx)

}

// initRedis 初始化 Redis Client
// https://github.com/redis/go-redis
func initRedis(_ context.Context) {

	// 初始化出来一个 redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", resource.Conf.Redis.Ip, resource.Conf.Redis.Port),
		Password: resource.Conf.Redis.Password, // no password set
		DB:       resource.Conf.Redis.DB,       // use default DB
	})

	// client 初始化为单例，放到 resource 里去
	resource.RedisClient = rdb
}

func initMysql(_ context.Context) {

	// 初始化 mysql 客户端
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		resource.Conf.Mysql.User, resource.Conf.Mysql.Password, resource.Conf.Mysql.Ip,
		resource.Conf.Mysql.Port, resource.Conf.Mysql.DB)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql init fail, err: %v", err)
	}
	// 设置连接池相关的信息
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("mysql init, set connect poll info fail, err: %v", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(resource.Conf.Mysql.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(resource.Conf.Mysql.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 将初始化的mysql client, 放到 resource 里去
	resource.MysqlClient = db
}
