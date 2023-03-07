package usecase

import (
	"context"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
)

func (c *Client) GetTagsByUserID(ctx context.Context, userID int64) ([]*model.Tag, error) {
	tags, err := c.repository.SelectAllTagsByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "usecase GetTagsByUserID: failed")
	}

	return tags, nil
}
