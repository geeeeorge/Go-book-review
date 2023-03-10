package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/geeeeorge/Go-book-review/gen/api"
)

type Interface interface {
	GetApiHealthz(ec echo.Context) error
	GetBooks(ec echo.Context) error
	PostBooks(ec echo.Context) error
	DeleteBook(ec echo.Context, bookId api.BookId) error
	GetBook(ec echo.Context, bookId api.BookId) error
	PutBookStatus(ec echo.Context, bookId api.BookId, params api.PutBookStatusParams) error
	PutBookTags(ec echo.Context, bookId api.BookId) error
	GetSummaries(ec echo.Context, params api.GetSummariesParams) error
	PostSummaries(ec echo.Context) error
	DeleteSummary(ec echo.Context, summaryId api.SummaryId) error
	GetSummary(ec echo.Context, summaryId api.SummaryId) error
	PutSummary(ec echo.Context, summaryId api.SummaryId) error
	GetTags(ec echo.Context) error
	PostTags(ec echo.Context) error
	DeleteTag(ec echo.Context, tagId api.TagId) error
	GetTag(ec echo.Context, tagId api.TagId) error
	PutTag(ec echo.Context, tagId api.TagId) error
	Login(ec echo.Context) error
	Signup(ec echo.Context) error
}
