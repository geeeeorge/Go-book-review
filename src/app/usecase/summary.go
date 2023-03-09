package usecase

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
)

func (c *Client) GetSummariesByBookID(ctx context.Context, userID int64, bookID int64) ([]*model.Summary, error) {
	summaries, err := c.repository.SelectAllSummariesByBookID(ctx, userID, bookID)
	if err != nil {
		return nil, errors.Wrap(err, "usecase GetTagsByUserID: failed")
	}

	return summaries, nil
}

func (c *Client) PostSummary(ctx context.Context, userID int64, summary *model.Summary) error {
	// book_idがuser_idのものか確認
	b, err := c.repository.SelectBookByID(ctx, userID, summary.BookID)
	if err != nil {
		return errors.Wrap(err, "usecase PostSummary: failed")
	}
	if b == nil {
		return errors.Wrap(err, "usecase PostSummary: There is no combination of user_id and book_id")
	}

	err = c.repository.InsertSummary(ctx, summary)
	if err != nil {
		return errors.Wrap(err, "usecase PostSummary: failed")
	}

	return nil
}

func (c *Client) GetSummaryByID(ctx context.Context, userID int64, id api.SummaryId) (*model.Summary, error) {
	summaries, err := c.repository.SelectSummaryByID(ctx, userID, id)
	if err != nil {
		return nil, errors.Wrap(err, "usecase GetSummaryByID: failed")
	}

	return summaries, nil
}

func (c *Client) PutSummary(ctx context.Context, userID int64, summary *model.Summary) error {
	// idがuser_idのものか確認
	b, err := c.repository.SelectSummaryByID(ctx, userID, summary.ID)
	if err != nil {
		return errors.Wrap(err, "usecase PutSummary: failed")
	}
	if b == nil {
		return errors.Wrap(err, "usecase PutSummary: There is no combination of user_id and id")
	}

	err = c.repository.UpdateSummary(ctx, summary)
	if err != nil {
		return errors.Wrap(err, "usecase Put Summary: failed")
	}

	return nil
}

func (c *Client) DeleteSummaryByID(ctx context.Context, userID int64, id api.SummaryId) error {
	// idがuser_idのものか確認
	b, err := c.repository.SelectSummaryByID(ctx, userID, id)
	if err != nil {
		return errors.Wrap(err, "usecase DeleteSummary: failed")
	}
	if b == nil {
		return errors.Wrap(err, "usecase DeleteSummary: There is no combination of user_id and id")
	}

	err = c.repository.DeleteSummaryByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "usecase DeleteSummaryByID: failed")
	}

	return nil
}
