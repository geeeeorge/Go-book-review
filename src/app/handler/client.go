package handler

import "github.com/geeeeorge/Go-book-review/src/app/usecase"

type Client struct {
	usecase usecase.Interface
}

func New(usecase usecase.Interface) Interface {
	return &Client{usecase: usecase}
}
