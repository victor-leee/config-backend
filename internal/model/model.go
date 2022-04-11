package model

import (
	"fmt"
	"github.com/victor-leee/config-backend/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func Init(cfg *config.Config) error {
	mysqlCfg := cfg.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", mysqlCfg.Username, mysqlCfg.Password,
		mysqlCfg.IP, mysqlCfg.Port, mysqlCfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	gormDB = db

	return nil
}

func GetDatabase() *gorm.DB {
	return gormDB
}
