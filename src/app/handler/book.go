package handler

import (
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/geeeeorge/Go-book-review/gen/api"
)

func (c *Client) GetBooks(ec echo.Context) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	books, err := c.usecase.GetBooks(ctx, uid)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusOK, books.API())
}

func (c *Client) PostBooks(ec echo.Context) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	var req api.Book
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, nil)
	}

	var book model.Book
	book.LoadAPI(&req)

	err := c.usecase.PostBook(ctx, uid, &book)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}

func (c *Client) GetBook(ec echo.Context, bookId api.BookId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	book, err := c.usecase.GetBookByID(ctx, uid, bookId)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusOK, book.API())
}

func (c *Client) DeleteBook(ec echo.Context, bookId api.BookId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	if err := c.usecase.DeleteBookByID(ctx, uid, bookId); err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}

func (c *Client) PutBookStatus(ec echo.Context, bookId api.BookId, params api.PutBookStatusParams) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	err := c.usecase.PutBookStatus(ctx, uid, bookId, string(params.Status))
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}

func (c *Client) PutBookTags(ec echo.Context, bookId api.BookId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	var req api.Book
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, nil)
	}
	req.Id = bookId

	var book model.Book
	book.LoadAPI(&req)
	err := c.usecase.PutBookTags(ctx, uid, &book)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}
