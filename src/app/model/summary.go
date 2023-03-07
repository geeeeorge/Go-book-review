package model

import "github.com/geeeeorge/Go-book-review/src/app/dao"

type Summary struct {
	ID      int64
	BookID  int64
	Content string
}

func (s *Summary) LoadDAO(d *dao.Summary) {
	s.ID = d.ID
	s.BookID = d.BookID
	s.Content = d.Content
}

func (s *Summary) DAO() *dao.Summary {
	return &dao.Summary{
		ID:      s.ID,
		BookID:  s.BookID,
		Content: s.Content,
	}
}
