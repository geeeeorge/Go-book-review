package repository

import (
	"context"
	"github.com/geeeeorge/Go-book-review/src/app/dao"
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

func (c *Client) SelectUserByUsername(ctx context.Context, username *string) (*model.User, error) {
	var d dao.User
	err := c.db.Get(&d, "SELECT id, username FROM users WHERE username = ?", username)
	if err != nil {
		return nil, errors.Wrap(err, "repository SelectUserByUsername: failed to select user")
	}

	var u model.User
	u.LoadDAO(&d)

	return &u, nil
}
