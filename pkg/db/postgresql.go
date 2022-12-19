package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	defaultHost                  = "127.0.0.1"
	defaultPost                  = 5432
	defaultPassword              = ""
	defaultTimeZone              = "Asia/Shanghai"
	defaultMaxIdleConnections    = 100
	defaultMaxOpenConnections    = 100
	defaultMaxConnectionLifeTime = time.Duration(10) * time.Second
	defaultLogLevel              = 1
)

type PGConfigs struct {
	host                  string
	port                  int
	username              string
	password              string
	database              string
	timeZone              string
	maxIdleConnections    int
	maxOpenConnections    int
	maxConnectionLifeTime time.Duration
}
type Options struct {
	host                  string
	port                  int
	password              string
	timeZone              string
	maxIdleConnections    int
	maxOpenConnections    int
	maxConnectionLifeTime time.Duration
}

type OptionFunc func(options *Options)

func WithHost(host string) OptionFunc {
	return OptionFunc(func(options *Options) {
		options.host = host
	})
}

func WithPort(port int) OptionFunc {
	return OptionFunc(func(options *Options) {
		options.port = port
	})
}

func WithPassword(password string) OptionFunc {
	return OptionFunc(func(options *Options) {
		options.password = password
	})
}

func WithTimeZone(timezone string) OptionFunc {
	return OptionFunc(func(options *Options) {
		options.timeZone = timezone
	})
}

func WithMaxIdleConnections(maxIdleConnections int) OptionFunc {
	return OptionFunc(func(options *Options) {
		options.maxIdleConnections = maxIdleConnections
	})
}

func WithMaxOpenConnections(maxOpenConnections int) OptionFunc {
	return OptionFunc(func(options *Options) {
		options.maxOpenConnections = maxOpenConnections
	})
}

func WithMaxConnectionLifeTime(maxConnectionLifeTime time.Duration) OptionFunc {
	return OptionFunc(func(options *Options) {
		options.maxConnectionLifeTime = maxConnectionLifeTime
	})
}

func NewPostgresql(username string, database string, options ...OptionFunc) (*gorm.DB, error) {
	ops := Options{
		host:                  defaultHost,
		port:                  defaultPost,
		password:              defaultPassword,
		timeZone:              defaultTimeZone,
		maxConnectionLifeTime: defaultMaxConnectionLifeTime,
		maxOpenConnections:    defaultMaxOpenConnections,
		maxIdleConnections:    defaultMaxIdleConnections,
	}

	for _, o := range options {
		o(&ops)
	}

	configs := PGConfigs{
		username:              username,
		database:              database,
		host:                  ops.host,
		port:                  ops.port,
		password:              ops.password,
		timeZone:              ops.timeZone,
		maxIdleConnections:    ops.maxIdleConnections,
		maxOpenConnections:    ops.maxOpenConnections,
		maxConnectionLifeTime: ops.maxConnectionLifeTime,
	}
	dns := fmt.Sprintf(`host=%s user=%s dbname=%s port=%d sslmode=disable TimeZone=%s password=%s`,
		configs.host,
		configs.username,
		configs.database,
		configs.port,
		configs.timeZone,
		configs.password,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}
	// 设定数据库最大连接数
	sqlDB.SetMaxOpenConns(configs.maxOpenConnections)
	// 设定数据库最长连接时长
	sqlDB.SetConnMaxLifetime(configs.maxConnectionLifeTime)
	// 设定数据库最大空闲数
	sqlDB.SetMaxIdleConns(configs.maxIdleConnections)

	return db, nil

}
