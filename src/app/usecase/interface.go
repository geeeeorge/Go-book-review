package usecase

import (
	"context"
	"github.com/geeeeorge/Go-book-review/src/app/model"
)

type Interface interface {
	SignUp(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *model.User) (string, error)
}
