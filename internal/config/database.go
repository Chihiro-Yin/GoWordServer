package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	const url = "127.0.0.1"
	const dbName = "33yzp"
	const username = "root"
	const password = "123456"

	// 构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, url, dbName)

	// 打开GORM连接
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %v", err))
	}

	// 关键：获取底层 *sql.DB 才能设置连接池参数
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get sql.DB: %v", err))
	}

	// 配置连接池（必须通过 sqlDB 配置）
	sqlDB.SetConnMaxLifetime(time.Minute * 3) // 连接最大存活时间（小于MySQL wait_timeout）
	sqlDB.SetConnMaxIdleTime(time.Minute * 1) // 连接最大空闲时间
	sqlDB.SetMaxIdleConns(10)                 // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)                // 最大打开连接数

	// 可选：验证连接有效性
	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	fmt.Println("Database connected successfully")
	DB = gormDB
}
