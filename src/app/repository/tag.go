package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/gen/api"
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

func (c *Client) InsertTag(ctx context.Context, userID int64, tag *model.Tag) error {
	_, err := c.db.NamedExecContext(ctx, "INSERT INTO tags (id, user_id, name) VALUES (:id, :user_id, :name)", tag.DAO(userID))
	if err != nil {
		return errors.Wrap(err, "InsertTag: failed to insert tag")
	}

	return nil
}

func (c *Client) SelectTagByID(ctx context.Context, userID int64, id api.TagId) (*model.Tag, error) {
	var t dao.Tag
	err := c.db.GetContext(ctx, &t, "SELECT id, user_id, name FROM tags WHERE id= ? AND user_id=?", id, userID)
	if err != nil {
		return nil, errors.Wrap(err, "SelectTagByID: failed to select tag")
	}

	var tag model.Tag
	tag.LoadDAO(&t)

	return &tag, nil
}

func (c *Client) UpdateTag(ctx context.Context, userID int64, tag *model.Tag) error {
	_, err := c.db.NamedExecContext(ctx, "UPDATE tags SET name=:name WHERE id=:id AND user_id=:user_id", tag.DAO(userID))
	if err != nil {
		return errors.Wrap(err, "UpdateTag: failed to delete tag")
	}

	return nil
}

func (c *Client) DeleteTagByID(ctx context.Context, userID int64, id api.TagId) error {
	_, err := c.db.QueryContext(ctx, "DELETE FROM tags WHERE id= ? AND user_id=?", id, userID)
	if err != nil {
		return errors.Wrap(err, "DeleteTagByID: failed to delete tag")
	}

	return nil
}
