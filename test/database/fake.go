package database

import (
	"github.com/jmoiron/sqlx"
)

func NewFakeDB() (*sqlx.DB, error) {
	return sqlx.Open("mysql", "mysql:passw0rd@tcp(localhost:3306)/book_review_test?charset=utf8mb4&parseTime=True&loc=Local")
}
