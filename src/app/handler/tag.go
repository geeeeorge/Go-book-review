package handler

import (
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/geeeeorge/Go-book-review/gen/api"
)

func (c *Client) GetTags(ec echo.Context) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	tags, err := c.usecase.GetTags(ctx, uid)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	res := make(api.Tags, 0, len(tags))
	for _, t := range tags {
		res = append(res, *t.API())
	}
	return ec.JSON(http.StatusOK, res)
}

func (c *Client) PostTags(ec echo.Context) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	var req api.Tag
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, nil)
	}

	var tag model.Tag
	tag.LoadAPI(&req)
	err := c.usecase.PostTag(ctx, uid, &tag)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}

func (c *Client) GetTag(ec echo.Context, tagId api.TagId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	tag, err := c.usecase.GetTagByID(ctx, uid, tagId)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusOK, tag.API())
}

func (c *Client) PutTag(ec echo.Context, tagId api.TagId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	var req api.Tag
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, nil)
	}
	req.Id = tagId

	var tag model.Tag
	tag.LoadAPI(&req)
	err := c.usecase.PutTag(ctx, uid, &tag)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}

func (c *Client) DeleteTag(ec echo.Context, tagId api.TagId) error {
	ctx := ec.Request().Context()
	uid, _ := ec.Get("user_id").(int64)

	if err := c.usecase.DeleteTagByID(ctx, uid, tagId); err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ec.JSON(http.StatusNoContent, "成功")
}
