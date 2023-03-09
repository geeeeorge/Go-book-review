package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/dao"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
)

func (c *Client) SelectAllSummariesByBookID(ctx context.Context, userID int64, bookID api.BookId) ([]*model.Summary, error) {
	ss := make([]dao.Summary, 0)
	err := c.db.SelectContext(ctx, &ss, `
		SELECT
		    s.id,
		    s.book_id,
		    s.content
		FROM
		    summaries s
		INNER JOIN
		    books b
		ON
		    s.book_id = b.id
		WHERE
		    s.bookID=?
		  AND
		    b.user_id=?
	`, bookID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "SelectAllSummariesByBookID: failed to select summaries")
	}

	summaries := make([]*model.Summary, 0, len(ss))
	for _, s := range ss {
		summary := model.Summary{}
		summary.LoadDAO(&s)
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (c *Client) InsertSummary(ctx context.Context, summary *model.Summary) error {
	_, err := c.db.NamedExecContext(ctx, "INSERT INTO summaries (id, book_id, content) VALUES (:id, :book_id, :content)", summary.DAO())
	if err != nil {
		return errors.Wrap(err, "InsertSummary: failed to insert summary")
	}

	return nil
}

func (c *Client) SelectSummaryByID(ctx context.Context, userID int64, id api.SummaryId) (*model.Summary, error) {
	var s dao.Summary
	err := c.db.GetContext(ctx, &s, `
		SELECT
		    s.id,
		    s.book_id,
		    s.content
		FROM
		    summaries s
		INNER JOIN
		    books b
		ON
		    s.book_id = b.id
		WHERE
		    s.id=?
		  AND
		    b.user_id=?
	`, id, userID)
	if err != nil {
		return nil, errors.Wrap(err, "SelectSummaryByID: failed to select summary")
	}

	var summary model.Summary
	summary.LoadDAO(&s)

	return &summary, nil
}

func (c *Client) UpdateSummary(ctx context.Context, summary *model.Summary) error {
	_, err := c.db.NamedExecContext(ctx, "UPDATE summaries SET content=:content WHERE id=:id", summary.DAO())
	if err != nil {
		return errors.Wrap(err, "UpdateSummary: failed to delete summary")
	}

	return nil
}

func (c *Client) DeleteSummaryByID(ctx context.Context, id api.SummaryId) error {
	_, err := c.db.QueryContext(ctx, "DELETE FROM summaries WHERE id=?", id)
	if err != nil {
		return errors.Wrap(err, "DeleteSummaryByID: failed to delete summary")
	}

	return nil
}
