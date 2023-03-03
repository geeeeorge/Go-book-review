package repository

import (
	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Interface {
	return &Client{db: db}
}
