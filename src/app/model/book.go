package model

import (
	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/dao"
)

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
		if d.SummaryID != 0 {
			_, ok := summaryMap[d.SummaryID]
			if !ok {
				b.Summaries = append(b.Summaries, &Summary{
					ID:      d.SummaryID,
					Content: d.SummaryContent,
				})
				summaryMap[d.SummaryID] = struct{}{}
			}
		}
		if d.TagID != 0 {
			_, ok := tagMap[d.TagID]
			if !ok {
				b.Tags = append(b.Tags, &Tag{
					ID:   d.TagID,
					Name: d.TagName,
				})
				tagMap[d.TagID] = struct{}{}
			}
		}
	}
}

func (b *Book) LoadAPI(a *api.Book) {
	b.ID = a.Id
	b.AmazonURL = a.AmazonUrl
	b.Status = string(a.Status)
	for _, tag := range *a.Tags {
		b.Tags = append(b.Tags, &Tag{
			ID: tag.Id,
		})
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

func (b *Book) DAOForTagBook() (*dao.Book, []*dao.TagBook) {
	book := &dao.Book{
		ID: b.ID,
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

func (b *Book) API() *api.Book {
	image := types.File{}
	image.InitFromBytes(b.Image, b.Title)

	summaries := make([]api.Summary, 0, len(b.Summaries))
	for _, s := range b.Summaries {
		summary := api.Summary{
			Id:      s.ID,
			BookId:  s.BookID,
			Content: s.Content,
		}
		summaries = append(summaries, summary)
	}

	tags := make([]api.Tag, 0, len(b.Tags))
	for _, t := range b.Tags {
		tag := api.Tag{
			Id:   t.ID,
			Name: t.Name,
		}
		tags = append(tags, tag)
	}

	return &api.Book{
		Id:        b.ID,
		Title:     b.Title,
		Image:     &image,
		AmazonUrl: b.AmazonURL,
		Status:    api.BookStatus(b.Status),
		Summaries: &summaries,
		Tags:      &tags,
	}
}

type Books struct {
	Books []*Book
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
		if d.SummaryID != 0 {
			_, ok = summaryMap[d.SummaryID]
			if !ok {
				book.Summaries = append(book.Summaries, &Summary{
					ID:      d.SummaryID,
					BookID:  d.ID,
					Content: d.SummaryContent,
				})
				summaryMap[d.SummaryID] = struct{}{}
			}
		}
		if d.TagID != 0 {
			_, ok = tagMap[d.TagID]
			if !ok {
				book.Tags = append(book.Tags, &Tag{
					ID:   d.TagID,
					Name: d.TagName,
				})
				tagMap[d.TagID] = struct{}{}
			}
		}
	}
	for _, book := range bookMap {
		b.Books = append(b.Books, book)
	}
}

func (b *Books) API() *api.Books {
	res := make([]api.Book, 0, len(b.Books))
	for _, book := range b.Books {
		res = append(res, *book.API())
	}
	return &res
}
