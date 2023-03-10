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
	GetBooks(ctx context.Context, userID int64) (*model.Books, error)
	PostBook(ctx context.Context, userID int64, book *model.Book) error
	GetBookByID(ctx context.Context, userID int64, id api.BookId) (*model.Book, error)
	DeleteBookByID(ctx context.Context, userID int64, id api.BookId) error
	PutBookStatus(ctx context.Context, userID int64, id api.BookId, status string) error
	PutBookTags(ctx context.Context, userID int64, book *model.Book) error
	GetTags(ctx context.Context, userID int64) ([]*model.Tag, error)
	PostTag(ctx context.Context, userID int64, tag *model.Tag) error
	GetTagByID(ctx context.Context, userID int64, id api.TagId) (*model.Tag, error)
	PutTag(ctx context.Context, userID int64, tag *model.Tag) error
	DeleteTagByID(ctx context.Context, userID int64, id api.TagId) error
	GetSummariesByBookID(ctx context.Context, userID int64, bookID api.BookId) ([]*model.Summary, error)
	PostSummary(ctx context.Context, userID int64, summary *model.Summary) error
	GetSummaryByID(ctx context.Context, userID int64, id api.SummaryId) (*model.Summary, error)
	PutSummary(ctx context.Context, userID int64, summary *model.Summary) error
	DeleteSummaryByID(ctx context.Context, userID int64, id api.SummaryId) error
}
