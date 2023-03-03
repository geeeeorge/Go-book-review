package model

import "github.com/geeeeorge/Go-book-review/src/app/dao"

type Book struct {
	ID        int
	UserID    int
	Title     string
	Image     []byte
	AmazonURL string
	Status    string
	Summaries []*Summary
	Tags      []*Tag
}

func (b *Book) LoadDAO(bss []*dao.BookSummary, bts []*dao.BookTag) {
	d := bss[0]
	if d.ID != 0 {
		b.ID = d.ID
	}
	if d.UserID != 0 {
		b.UserID = d.UserID
	}
	if d.Title != "" {
		b.Title = d.Title
	}
	if d.Image != nil {
		b.Image = d.Image
	}
	if d.AmazonURL != "" {
		b.AmazonURL = d.AmazonURL
	}
	if d.Status != "" {
		b.Status = d.Status
	}

	summaries := make([]*Summary, 0, len(bss))
	for _, bs := range bss {
		s := Summary{}
		if bs.SummaryID != 0 {
			s.ID = bs.SummaryID
		}
		if bs.SummaryContent != "" {
			s.Content = bs.SummaryContent
		}
		summaries = append(summaries, &s)
	}

	tags := make([]*Tag, 0, len(bts))
	for _, bt := range bts {
		t := Tag{}
		if bt.TagID != 0 {
			t.ID = bt.TagID
		}
		if bt.TagName != "" {
			t.Name = bt.TagName
		}
		tags = append(tags, &t)
	}
}

func (b *Book) DAO(uid int) (*dao.Book, []*dao.TagBook) {
	book := &dao.Book{
		ID:        b.ID,
		UserID:    b.UserID,
		Title:     b.Title,
		Image:     b.Image,
		AmazonURL: b.AmazonURL,
		Status:    b.Status,
	}
	tagBooks := make([]*dao.TagBook, 0, len(b.Tags))
	for _, t := range b.Tags {
		tagBook := dao.TagBook{
			BookID: b.ID,
			TagID:  t.ID,
		}
		tagBooks = append(tagBooks, &tagBook)
	}
	return book, tagBooks
}
