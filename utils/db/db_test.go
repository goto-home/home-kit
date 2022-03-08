package db

import "testing"

func TestNewGormConn(t *testing.T) {
	cfg := &DatabaseCfg{
		Username:     "root",
		Password:     "root",
		Url:          "127.0.0.1:3306",
		Port:         "3306",
		DatabaseName: "mysql",
	}
	_, f := NewGormConn(cfg)
	defer f()
}
