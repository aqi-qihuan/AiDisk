package config

import (
	"fmt"
	"sync"

	"github.com/aqi/AqiCloud-AgentPan-Go/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbOnce sync.Once
	db     *gorm.DB
)

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		cfg := GetConfig()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDatabase)

		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic("数据库连接失败: " + err.Error())
		}
	})
	return db
}

func AutoMigrate() {
	db := GetDB()
	db.AutoMigrate(
		&model.Account{},
		&model.Storage{},
		&model.AccountFile{},
		&model.File{},
		&model.Share{},
		&model.ShareFile{},
		&model.FileChunk{},
		&model.FileType{},
		&model.FileSuffix{},
	)
}
