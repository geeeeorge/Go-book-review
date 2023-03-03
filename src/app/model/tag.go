package model

import "github.com/geeeeorge/Go-book-review/src/app/dao"

type Tag struct {
	ID     int64
	UserID int64
	Name   string
}

func (t *Tag) LoadDAO(d *dao.Tag) {
	t.ID = d.ID
	t.UserID = d.UserID
	t.Name = d.Name
}

func (t *Tag) DAO() *dao.Tag {
	ret := &dao.Tag{
		ID:     t.ID,
		UserID: t.UserID,
		Name:   t.Name,
	}
	return ret
}
