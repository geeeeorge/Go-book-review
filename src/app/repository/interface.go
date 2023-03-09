package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/model"
)

type Interface interface {
	InsertUser(ctx context.Context, user *model.User) error
	SelectUserByUsername(ctx context.Context, username *string) (*model.User, error)
	SelectAllTags(ctx context.Context, userID int64) ([]*model.Tag, error)
	InsertTag(ctx context.Context, userID int64, tag *model.Tag) error
	SelectTagByID(ctx context.Context, userID int64, id api.TagId) (*model.Tag, error)
	UpdateTag(ctx context.Context, userID int64, tag *model.Tag) error
	DeleteTagByID(ctx context.Context, userID int64, id api.TagId) error
	SelectAllSummariesByBookID(ctx context.Context, userID int64, bookID api.BookId) ([]*model.Summary, error)
	InsertSummary(ctx context.Context, summary *model.Summary) error
	SelectSummaryByID(ctx context.Context, userID int64, id api.SummaryId) (*model.Summary, error)
	UpdateSummary(ctx context.Context, summary *model.Summary) error
	DeleteSummaryByID(ctx context.Context, id api.SummaryId) error
	SelectBookByID(ctx context.Context, userID int64, id api.BookId) (*model.Book, error)
}
