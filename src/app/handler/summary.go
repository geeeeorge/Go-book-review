package handler

import (
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/geeeeorge/Go-book-review/gen/api"
)

func (c *Client) GetSummaries(ec echo.Context, params api.GetSummariesParams) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	summaries, err := c.usecase.GetSummariesByBookID(ctx, uid, params.BookId)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	res := make(api.Summaries, 0, len(summaries))
	for _, s := range summaries {
		res = append(res, *s.API())
	}
	return ec.JSON(http.StatusOK, res)
}

func (c *Client) PostSummaries(ec echo.Context) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	var req api.Summary
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, nil)
	}

	var summary model.Summary
	summary.LoadAPI(&req)
	err := c.usecase.PostSummary(ctx, uid, &summary)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}

func (c *Client) GetSummary(ec echo.Context, summaryId api.SummaryId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	summary, err := c.usecase.GetSummaryByID(ctx, uid, summaryId)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusOK, summary.API())
}

func (c *Client) PutSummary(ec echo.Context, summaryId api.SummaryId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	var req api.Summary
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, nil)
	}
	req.Id = summaryId

	var summary model.Summary
	summary.LoadAPI(&req)
	err := c.usecase.PutSummary(ctx, uid, &summary)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}

func (c *Client) DeleteSummary(ec echo.Context, summaryId api.SummaryId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	if err := c.usecase.DeleteSummaryByID(ctx, uid, summaryId); err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}
