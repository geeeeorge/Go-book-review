package model

import "github.com/geeeeorge/Go-book-review/src/app/dao"

type Book struct {
	ID        int64
	Title     string
	Image     []byte
	AmazonURL string
	Status    string
	Summaries []*Summary
	Tags      []*Tag
}

func (b *Book) LoadDAOFromBook(d *dao.Book) {
	b.ID = d.ID
	b.Title = d.Title
	b.Image = d.Image
	b.AmazonURL = d.AmazonURL
	b.Status = d.Status
}

func (b *Book) LoadDAOFromBookSummaryTag(bst []*dao.BookSummaryTag) {
	d := bst[0]
	b.ID = d.ID
	b.Title = d.Title
	b.Image = d.Image
	b.AmazonURL = d.AmazonURL
	b.Status = d.Status

	summaryMap := map[int64]interface{}{}
	tagMap := map[int64]interface{}{}

	for _, d := range bst {
		_, ok := summaryMap[d.SummaryID]
		if !ok {
			b.Summaries = append(b.Summaries, &Summary{
				ID:      d.SummaryID,
				Content: d.SummaryContent,
			})
		}
		_, ok = tagMap[d.TagID]
		if !ok {
			b.Tags = append(b.Tags, &Tag{
				ID:   d.TagID,
				Name: d.TagName,
			})
		}
	}
}

func (b *Book) DAO(uid int64) (*dao.Book, []*dao.TagBook) {
	book := &dao.Book{
		ID:        b.ID,
		UserID:    uid,
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

type Books struct {
	books []*Book
}

func (b *Books) LoadDAO(bst []*dao.BookSummaryTag) {
	bookMap := map[int64]*Book{}
	summaryMap := map[int64]interface{}{}
	tagMap := map[int64]interface{}{}

	for _, d := range bst {
		book, ok := bookMap[d.ID]
		if !ok {
			book = &Book{
				ID:        d.ID,
				Title:     d.Title,
				Image:     d.Image,
				AmazonURL: d.AmazonURL,
				Status:    d.Status,
				Summaries: []*Summary{},
				Tags:      []*Tag{},
			}
			bookMap[d.ID] = book
		}
		_, ok = summaryMap[d.SummaryID]
		if !ok {
			book.Summaries = append(book.Summaries, &Summary{
				ID:      d.SummaryID,
				Content: d.SummaryContent,
			})
		}
		_, ok = tagMap[d.TagID]
		if !ok {
			book.Tags = append(book.Tags, &Tag{
				ID:   d.TagID,
				Name: d.TagName,
			})
		}
	}
	for _, book := range bookMap {
		b.books = append(b.books, book)
	}
}
