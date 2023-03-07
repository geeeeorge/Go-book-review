package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/src/app/dao"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
)

func (c *Client) SelectAllTagsByUserID(ctx context.Context, userID int64) ([]*model.Tag, error) {
	ts := make([]dao.Tag, 0)
	err := c.db.SelectContext(ctx, &ts, "SELECT id, user_id, name FROM tags WHERE user_id= ?", userID)
	if err != nil {
		return nil, errors.Wrap(err, "SelectAllTagsByUserID: failed to select tags")
	}

	tags := make([]*model.Tag, 0, len(ts))
	for _, t := range ts {
		tag := model.Tag{}
		tag.LoadDAO(&t)
		tags = append(tags, &tag)
	}

	return tags, nil
}
