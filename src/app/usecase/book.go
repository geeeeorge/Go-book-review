package usecase

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/pkg/scraping"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
)

func (c *Client) GetBooks(ctx context.Context, userID int64) (*model.Books, error) {
	books, err := c.repository.SelectAllBookSummaryTags(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "usecase GetBooksBy: failed")
	}

	return books, nil
}

func (c *Client) PostBook(ctx context.Context, userID int64, book *model.Book) error {
	data, err := scraping.AmazonURL(book.AmazonURL)
	if err != nil {
		return errors.Wrap(err, "usecase PostBook: scraping failed")
	}
	book.Title = data.Title
	book.Image = data.Image

	err = c.repository.InsertBook(ctx, userID, book)
	if err != nil {
		return errors.Wrap(err, "usecase PostBook: failed")
	}

	return nil
}

func (c *Client) GetBookByID(ctx context.Context, userID int64, id api.BookId) (*model.Book, error) {
	book, err := c.repository.SelectBookSummaryTagByID(ctx, userID, id)
	if err != nil {
		return nil, errors.Wrap(err, "usecase GetBookByID: failed")
	}

	return book, nil
}

func (c *Client) DeleteBookByID(ctx context.Context, userID int64, id api.BookId) error {
	err := c.repository.DeleteBookByID(ctx, userID, id)
	if err != nil {
		return errors.Wrap(err, "usecase DeleteBookByID: failed")
	}

	return nil
}

func (c *Client) PutBookStatus(ctx context.Context, userID int64, id api.BookId, status string) error {
	err := c.repository.UpdateBookStatus(ctx, userID, id, status)
	if err != nil {
		return errors.Wrap(err, "usecase PutBookStatus: failed")
	}

	return nil
}

func (c *Client) PutBookTags(ctx context.Context, userID int64, book *model.Book) error {
	// book_idがuser_idのものか確認
	b, err := c.repository.SelectBookByID(ctx, userID, book.ID)
	if err != nil {
		return errors.Wrap(err, "usecase PutBookTags: failed")
	}
	if b == nil {
		return errors.Wrap(err, "usecase PutBookTags: There is no combination of user_id and id")
	}

	err = c.repository.UpdateBookTags(ctx, book)
	if err != nil {
		return errors.Wrap(err, "usecase PutBookTags: failed")
	}

	return nil
}
