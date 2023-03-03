package model

import "github.com/geeeeorge/Go-book-review/src/app/dao"

type Summary struct {
	ID      int
	BookID  int
	Content string
}

func (s *Summary) LoadDAO(d *dao.Summary) {
	if d.ID != 0 {
		s.ID = d.ID
	}
	if d.BookID != 0 {
		s.BookID = d.BookID
	}
	if d.Content != "" {
		s.Content = d.Content
	}
}

func (s *Summary) DAO() *dao.Summary {
	ret := &dao.Summary{
		ID:      s.ID,
		BookID:  s.BookID,
		Content: s.Content,
	}
	return ret
}
