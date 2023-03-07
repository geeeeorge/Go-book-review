package usecase

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/model"
)

type Interface interface {
	SignUp(ctx context.Context, user *model.User) error
	Login(ctx context.Context, user *model.User) (string, error)
	GetUserIDByUsername(ctx context.Context, username *string) (int64, error)
	GetTagsByUserID(ctx context.Context, userID int64) ([]*model.Tag, error)
	PostTag(ctx context.Context, userID int64, tag *model.Tag) error
	GetTagByID(ctx context.Context, userID int64, id api.TagId) (*model.Tag, error)
	PutTag(ctx context.Context, userID int64, tag *model.Tag) error
	DeleteTagByID(ctx context.Context, userID int64, id api.TagId) error
}
