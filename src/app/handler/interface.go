package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/geeeeorge/Go-book-review/gen/api"
)

type Interface interface {
	GetApiHealthz(ctx echo.Context) error
	GetBooks(ctx echo.Context, params api.GetBooksParams) error
	PostBooks(ctx echo.Context) error
	DeleteBook(ctx echo.Context, bookId api.BookId) error
	GetBook(ctx echo.Context, bookId api.BookId) error
	PutBook(ctx echo.Context, bookId api.BookId, params api.PutBookParams) error
	GetSummaries(ctx echo.Context, params api.GetSummariesParams) error
	PostSummaries(ctx echo.Context) error
	DeleteSummary(ctx echo.Context, summaryId api.SummaryId) error
	GetSummary(ctx echo.Context, summaryId api.SummaryId) error
	PutSummary(ctx echo.Context, summaryId api.SummaryId) error
	GetTags(ctx echo.Context, params api.GetTagsParams) error
	PostTags(ctx echo.Context) error
	DeleteTag(ctx echo.Context, tagId api.TagId) error
	GetTag(ctx echo.Context, tagId api.TagId) error
	PutTag(ctx echo.Context, tagId api.TagId) error
	Login(ctx echo.Context) error
	Signup(ctx echo.Context) error
}

func (c *Client) GetApiHealthz(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetBooks(ctx echo.Context, params api.GetBooksParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DeleteBook(ctx echo.Context, bookId api.BookId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetBook(ctx echo.Context, bookId api.BookId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) PutBook(ctx echo.Context, bookId api.BookId, params api.PutBookParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetSummaries(ctx echo.Context, params api.GetSummariesParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DeleteSummary(ctx echo.Context, summaryId api.SummaryId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetSummary(ctx echo.Context, summaryId api.SummaryId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) PutSummary(ctx echo.Context, summaryId api.SummaryId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetTags(ctx echo.Context, params api.GetTagsParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DeleteTag(ctx echo.Context, tagId api.TagId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetTag(ctx echo.Context, tagId api.TagId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) PutTag(ctx echo.Context, tagId api.TagId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Login(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Signup(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) PostBooks(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) PostSummaries(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) PostTags(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
