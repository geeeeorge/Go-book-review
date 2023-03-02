package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB = nil

func NewMysqlDB(dsn string) (*gorm.DB, error) {
	if mysqlDB != nil {
		return mysqlDB, nil
	}

	mysqlDB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return mysqlDB, nil
}
