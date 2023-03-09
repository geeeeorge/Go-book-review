package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/dao"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
)

func (c *Client) SelectBookByID(ctx context.Context, userID int64, id api.BookId) (*model.Book, error) {
	var b dao.Book
	err := c.db.GetContext(ctx, &b, `
		SELECT
		    id,
		    user_id,
		    title,
		    image,
		    amazon_url,
		    status
		FROM
		    books
		WHERE
		    id=?
		  AND
		    user_id=?
	`, id, userID)
	if err != nil {
		return nil, errors.Wrap(err, "SelectBookByID: failed to select book")
	}

	var book model.Book
	book.LoadDAOFromBook(&b)

	return &book, nil
}
