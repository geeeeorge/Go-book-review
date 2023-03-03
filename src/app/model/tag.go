package model

import "github.com/geeeeorge/Go-book-review/src/app/dao"

type Tag struct {
	ID     int
	UserID int
	Name   string
}

func (t *Tag) LoadDAO(d *dao.Tag) {
	if d.ID != 0 {
		t.ID = d.ID
	}
	if d.UserID != 0 {
		t.UserID = d.UserID
	}
	if d.Name != "" {
		t.Name = d.Name
	}
}

func (t *Tag) DAO() *dao.Tag {
	ret := &dao.Tag{
		ID:     t.ID,
		UserID: t.UserID,
		Name:   t.Name,
	}
	return ret
}
