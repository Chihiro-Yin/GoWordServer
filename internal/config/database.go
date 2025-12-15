package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

// 数据库类型枚举（可通过环境变量覆盖）
const (
	DBTypeMySQL  = "mysql"
	DBTypeSQLite = "sqlite"
	// 默认使用 MySQL，可改为 SQLite
	defaultDBType = DBTypeMySQL
)

// InitDB 初始化数据库（自动判断 MySQL/SQLite）
func InitDB() {
	// 1. 获取数据库类型（优先读取环境变量，无则用默认）
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = defaultDBType
	}

	// 2. 根据类型初始化不同数据库
	var (
		gormDB *gorm.DB
		err    error
	)
	switch dbType {
	case DBTypeMySQL:
		gormDB, err = initMySQL()
	case DBTypeSQLite:
		gormDB, err = initSQLite()
	default:
		panic(fmt.Sprintf("不支持的数据库类型：%s（仅支持 mysql/sqlite）", dbType))
	}

	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败：%v", err))
	}

	// 3. 配置连接池（通用配置，MySQL/SQLite 都适用）
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(fmt.Sprintf("获取底层 sql.DB 失败：%v", err))
	}

	// MySQL 连接池配置（SQLite 建议调整为 1，因为 SQLite 是单文件数据库）
	if dbType == DBTypeMySQL {
		sqlDB.SetConnMaxLifetime(time.Minute * 3) // 连接最大存活时间
		sqlDB.SetConnMaxIdleTime(time.Minute * 1) // 连接最大空闲时间
		sqlDB.SetMaxIdleConns(10)                 // 最大空闲连接数
		sqlDB.SetMaxOpenConns(100)                // 最大打开连接数
	} else if dbType == DBTypeSQLite {
		// SQLite 特殊配置：最多 1 个连接（避免文件锁冲突）
		sqlDB.SetMaxOpenConns(1)
		sqlDB.SetMaxIdleConns(1)
		sqlDB.SetConnMaxLifetime(0) // 无超时（SQLite 无需）
	}

	// 4. 验证连接
	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Sprintf("数据库连接验证失败：%v", err))
	}

	fmt.Printf("✅ %s 数据库连接成功\n", dbType)
	DB = gormDB
}

// initMySQL 初始化 MySQL 数据库（原有逻辑）
func initMySQL() (*gorm.DB, error) {
	const (
		url      = "127.0.0.1"
		dbName   = "WordDB"
		username = "WordUser"
		password = "WordPass123123"
	)

	// 构建 MySQL DSN

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, url, dbName)

	// 打开 MySQL 连接
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// initSQLite 初始化 SQLite 数据库
func initSQLite() (*gorm.DB, error) {
	// 1. 创建数据目录（不存在则创建）
	workDir, _ := os.Getwd()
	dataDir := filepath.Join(workDir, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("创建数据目录失败：%v", err)
	}

	// 2. SQLite 数据库文件路径（data/sqlite.db）
	dbPath := filepath.Join(dataDir, "sqlite.db")
	fmt.Printf("📂 SQLite 数据库文件路径：%s\n", dbPath)

	dsn := fmt.Sprintf("file:%s?cache=shared", dbPath)
	gormDB, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "sqlite", // 指定纯 Go 驱动
		DSN:        dsn,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // SQLite 推荐
	})
	if err != nil {
		return nil, fmt.Errorf("连接 SQLite 失败：%v", err)
	}

	return gormDB, nil
}
