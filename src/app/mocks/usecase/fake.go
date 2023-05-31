package usecase

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/geeeeorge/Go-book-review/src/app/usecase"
)

type Fake struct {
	Books *model.Books
}

var _ usecase.Interface = (*Fake)(nil)

func (f Fake) SignUp(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) Login(ctx context.Context, user *model.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (f Fake) GetUserIDByUsername(ctx context.Context, username *string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (f Fake) GetBooks(ctx context.Context, userID int64) (*model.Books, error) {
	return f.Books, nil
}

func (f Fake) PostBook(ctx context.Context, userID int64, book *model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) GetBookByID(ctx context.Context, userID int64, id api.BookId) (*model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (f Fake) DeleteBookByID(ctx context.Context, userID int64, id api.BookId) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) PutBookStatus(ctx context.Context, userID int64, id api.BookId, status string) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) PutBookTags(ctx context.Context, userID int64, book *model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) GetTags(ctx context.Context, userID int64) ([]*model.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (f Fake) PostTag(ctx context.Context, userID int64, tag *model.Tag) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) GetTagByID(ctx context.Context, userID int64, id api.TagId) (*model.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (f Fake) PutTag(ctx context.Context, userID int64, tag *model.Tag) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) DeleteTagByID(ctx context.Context, userID int64, id api.TagId) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) GetSummariesByBookID(ctx context.Context, userID int64, bookID api.BookId) ([]*model.Summary, error) {
	//TODO implement me
	panic("implement me")
}

func (f Fake) PostSummary(ctx context.Context, userID int64, summary *model.Summary) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) GetSummaryByID(ctx context.Context, userID int64, id api.SummaryId) (*model.Summary, error) {
	//TODO implement me
	panic("implement me")
}

func (f Fake) PutSummary(ctx context.Context, userID int64, summary *model.Summary) error {
	//TODO implement me
	panic("implement me")
}

func (f Fake) DeleteSummaryByID(ctx context.Context, userID int64, id api.SummaryId) error {
	//TODO implement me
	panic("implement me")
}

var _ usecase.Interface = (*Fake)(nil)
