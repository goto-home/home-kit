package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UtilsDbConn 目前持久化框架只考虑gorm实现
type UtilsDbConn struct {
	cfg *DatabaseCfg
	conn *gorm.DB
}

func NewGormConn(cfg *DatabaseCfg) (*UtilsDbConn, func()) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Url, cfg.DatabaseName)
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	// test ping
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := sqlDb.Ping(); err != nil {
		panic(fmt.Sprintf("mysql ping fail,%v", err))
	}
	cleanUp := func() {
		// no handler err
		_ = sqlDb.Close()
	}
	return &UtilsDbConn{
		conn: db,
	}, cleanUp
}
