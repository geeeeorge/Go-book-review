package model

import (
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/dao"
)

type Summary struct {
	ID      int64
	BookID  int64
	Content string
}

func (s *Summary) LoadDAO(a *dao.Summary) {
	s.ID = a.ID
	s.BookID = a.BookID
	s.Content = a.Content
}

func (s *Summary) LoadAPI(d *api.Summary) {
	s.ID = d.Id
	s.BookID = d.BookId
	s.Content = d.Content
}

func (s *Summary) DAO() *dao.Summary {
	return &dao.Summary{
		ID:      s.ID,
		BookID:  s.BookID,
		Content: s.Content,
	}
}

func (s *Summary) API() *api.Summary {
	return &api.Summary{
		Id:      s.ID,
		BookId:  s.BookID,
		Content: s.Content,
	}
}
