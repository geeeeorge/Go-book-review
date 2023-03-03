package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/src/app/model"
)

type Interface interface {
	InsertUser(ctx context.Context, user model.User) error
}
