package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/pkg/errors"
)

func (c *Client) InsertUser(ctx context.Context, rec *model.User) error {
	_, err := c.db.NamedExec("INSERT INTO users (username) VALUES (:username);", rec.DAO())
	if err != nil {
		return errors.Wrap(err, "InsertUser: failed to insert user")
	}

	return nil
}
