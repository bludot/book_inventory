package db

import (
	"fmt"
	"github.com/bludot/tempmee/inventory/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDatabase(cfg config.DBConfig) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DB{DB: db}
}
