package model

import (
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/dao"
)

type Tag struct {
	ID   int64
	Name string
}

func (t *Tag) LoadDAO(d *dao.Tag) {
	t.ID = d.ID
	t.Name = d.Name
}

func (t *Tag) LoadAPI(a *api.Tag) {
	t.ID = a.Id
	t.Name = a.Name
}

func (t *Tag) DAO(uid int64) *dao.Tag {
	return &dao.Tag{
		ID:     t.ID,
		UserID: uid,
		Name:   t.Name,
	}
}

func (t *Tag) API() *api.Tag {
	return &api.Tag{
		Id:   t.ID,
		Name: t.Name,
	}
}
