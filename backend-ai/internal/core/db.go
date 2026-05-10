package core

import (
	"log"
	"sync"
	"time"

	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var (
	dbOnce sync.Once
	db     *gorm.DB
)

// GetDB 获取全局 MySQL 连接
func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		dsn := GetConfig().MySQLDSN()
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gormlogger.Default.LogMode(gormlogger.Silent),
		})
		if err != nil {
			log.Fatalf("MySQL 连接失败: %v", err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("获取 sql.DB 失败: %v", err)
		}
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)

		// 自动建表
		db.AutoMigrate(&models.ChatLog{})
	})
	return db
}
