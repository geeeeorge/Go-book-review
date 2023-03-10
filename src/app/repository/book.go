package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/dao"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
)

func (c *Client) SelectAllBookSummaryTags(ctx context.Context, userID int64) (*model.Books, error) {
	bs := make([]*dao.BookSummaryTag, 0)
	err := c.db.SelectContext(ctx, &bs, `
		SELECT
		    b.id,
		    b.user_id,
		    b.title,
		    b.image,
		    b.amazon_url,
		    b.status,
		    IFNULL(s.id, 0) AS summary_id,
		    IFNULL(s.content, '') AS summary_content,
		    IFNULL(t.id, 0) AS tag_id,
		    IFNULL(t.name, '') AS tag_name
		FROM
		    books b
		LEFT JOIN
		    summaries s
		ON
		    b.id = s.book_id
		LEFT JOIN
		    tag_books tb
		ON
		    b.id = tb.book_id
		INNER JOIN
		    tags t
		ON
		    tb.tag_id = t.id
		WHERE
		    b.user_id=?
	`, userID)
	if err != nil {
		return nil, errors.Wrap(err, "SelectAllBookSummaryTags: failed to select books")
	}

	var books model.Books
	books.LoadDAO(bs)

	return &books, nil
}

func (c *Client) InsertBook(ctx context.Context, userID int64, book *model.Book) error {
	// トランザクションでerrorの際にRollbackできるように
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "InsertBook: failed to begin tx")
	}

	b, tb := book.DAO(userID)

	// booksテーブルにInsert
	_, err = c.db.NamedExecContext(ctx, `
		INSERT INTO
		    books (user_id, title, image, amazon_url, status)
		VALUES
		    (:user_id, :title, :image, :amazon_url, :status)
	`, b)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(err, "InsertBook: failed to rollback")
		}
		return errors.Wrap(err, "InsertBook: failed to insert book")
	}

	// tag_booksテーブルにInsert
	_, err = c.db.NamedExecContext(ctx, `
		INSERT INTO
		    tag_books (book_id, tag_id)
		VALUES
		    (LAST_INSERT_ID(), :tag_id)
	`, tb)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(err, "InsertBook: failed to rollback")
		}
		return errors.Wrap(err, "InsertBook: failed to insert tag_book")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "InsertBook: failed to commit")
	}

	return nil
}

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

func (c *Client) SelectBookSummaryTagByID(ctx context.Context, userID int64, id api.BookId) (*model.Book, error) {
	bs := make([]*dao.BookSummaryTag, 0)
	err := c.db.SelectContext(ctx, &bs, `
		SELECT
		    b.id,
		    b.user_id,
		    b.title,
		    b.image,
		    b.amazon_url,
		    b.status,
		    IFNULL(s.id, 0) AS summary_id,
		    IFNULL(s.content, '') AS summary_content,
		    IFNULL(t.id, 0) AS tag_id,
		    IFNULL(t.name, '') AS tag_name
		FROM
		    books b
		LEFT JOIN
		    summaries s
		ON
		    b.id = s.book_id
		LEFT JOIN
		    tag_books tb
		ON
		    b.id = tb.book_id
		INNER JOIN
		    tags t
		ON
		    tb.tag_id = t.id
		WHERE
		    b.id=? 
		  AND
		    b.user_id=?
	`, id, userID)
	if err != nil {
		return nil, errors.Wrap(err, "SelectBookSummaryTagByID: failed to select book")
	}

	if len(bs) == 0 {
		return nil, errors.Errorf("SelectAllBookSummaryTags: There is no book with book_id = %d", id)
	}

	var book model.Book
	book.LoadDAOFromBookSummaryTag(bs)

	return &book, nil
}

func (c *Client) DeleteBookByID(ctx context.Context, userID int64, id api.BookId) error {
	_, err := c.db.QueryContext(ctx, "DELETE FROM books WHERE id=? AND user_id=?", id, userID)
	if err != nil {
		return errors.Wrap(err, "DeleteBookByID: failed to delete book")
	}

	return nil
}

func (c *Client) UpdateBookStatus(ctx context.Context, userID int64, id api.BookId, status string) error {
	_, err := c.db.QueryContext(ctx, "UPDATE books SET status=? WHERE id=? AND user_id=?", status, id, userID)
	if err != nil {
		return errors.Wrap(err, "UpdateBookStatus: failed to update book")
	}

	return nil
}

func (c *Client) UpdateBookTags(ctx context.Context, book *model.Book) error {
	// トランザクションでerrorの際にRollbackできるように
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "InsertBook: failed to begin tx")
	}

	b, tb := book.DAOForTagBook()

	// tag_booksテーブルからDelete
	_, err = c.db.QueryContext(ctx, "DELETE FROM tag_books WHERE book_id=?", b.ID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(err, "UpdateBookTags: failed to rollback")
		}
		return errors.Wrap(err, "UpdateBookTags: failed to delete tag_book")
	}

	// tag_booksテーブルにInsert
	_, err = c.db.NamedExecContext(ctx, `
		INSERT INTO
		    tag_books (book_id, tag_id)
		VALUES
		    (:book_id, :tag_id)
	`, tb)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.Wrap(err, "UpdateBookTags: failed to rollback")
		}
		return errors.Wrap(err, "UpdateBookTags: failed to insert tag_book")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "UpdateBookTags: failed to commit")
	}

	return nil
}
