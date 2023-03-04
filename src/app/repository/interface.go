package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/src/app/model"
)

type Interface interface {
	InsertUser(ctx context.Context, user *model.User) error
	SelectUserByUsername(ctx context.Context, username *string) (*model.User, error)
}
