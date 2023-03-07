package usecase

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
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

func (c *Client) PostTag(ctx context.Context, userID int64, tag *model.Tag) error {
	err := c.repository.InsertTag(ctx, userID, tag)
	if err != nil {
		return errors.Wrap(err, "usecase PostTag: failed")
	}

	return nil
}

func (c *Client) GetTagByID(ctx context.Context, userID int64, id api.TagId) (*model.Tag, error) {
	tag, err := c.repository.SelectTagByID(ctx, userID, id)
	if err != nil {
		return nil, errors.Wrap(err, "usecase GetTagByID: failed")
	}

	return tag, nil
}

func (c *Client) PutTag(ctx context.Context, userID int64, tag *model.Tag) error {
	err := c.repository.UpdateTag(ctx, userID, tag)
	if err != nil {
		return errors.Wrap(err, "usecase PutTag: failed")
	}

	return nil
}

func (c *Client) DeleteTagByID(ctx context.Context, userID int64, id api.TagId) error {
	err := c.repository.DeleteTagByID(ctx, userID, id)
	if err != nil {
		return errors.Wrap(err, "usecase DeleteTag: failed")
	}

	return nil
}
