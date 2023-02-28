package usecase

import "github.com/geeeeorge/Go-book-review/src/app/repository"

type Client struct {
	repository repository.Interface
}

func New(repository repository.Interface) Interface {
	return &Client{repository: repository}
}
