package database

import (
	"github.com/jmoiron/sqlx"
)

var mysqlDB *sqlx.DB = nil

func NewMysqlDB(dsn string) (*sqlx.DB, error) {
	if mysqlDB != nil {
		return mysqlDB, nil
	}

	mysqlDB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return mysqlDB, nil
}
