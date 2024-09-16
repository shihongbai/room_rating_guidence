package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"room_rating_guidence/common/config"
	"time"
)

var DB *gorm.DB

// SetupDBLink 数据库初始化
func SetupDBLink() error {
	var err error
	var dbConfig = config.Config.Db

	// 连接MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Charset,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	// 检查数据库是否存在
	var result int64
	DB.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbConfig.Db).Scan(&result)
	if result == 0 {
		// 数据库不存在，创建数据库
		createDBSQL := fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8 COLLATE utf8_general_ci", dbConfig.Db)
		if err := DB.Exec(createDBSQL).Error; err != nil {
			panic(fmt.Sprintf("failed to create database: %v", err))
		}

		fmt.Println("Database created:", dbConfig.Db)
		// 强制等待一段时间，以确保数据库创建完毕
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Database already exists:", dbConfig.Db)
	}

	// 重新连接到 MySQL 服务器以确保数据库列表刷新
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	// 连接到数据库
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)
	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)

	return nil
}
