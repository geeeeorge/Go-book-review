package repository

import "gorm.io/gorm"

type Client struct {
	db *gorm.DB
}

func New(db *gorm.DB) Interface {
	return &Client{db: db}
}
