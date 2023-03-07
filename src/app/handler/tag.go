package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/geeeeorge/Go-book-review/gen/api"
)

func (c *Client) GetTags(ec echo.Context) error {
	ctx := ec.Request().Context()

	username := ec.Get("username").(string)
	uid, err := c.usecase.GetUserIDByUsername(ctx, &username)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	tags, err := c.usecase.GetTagsByUserID(ctx, uid)
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
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetTag(ec echo.Context, tagId api.TagId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) PutTag(ec echo.Context, tagId api.TagId) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DeleteTag(ec echo.Context, tagId api.TagId) error {
	//TODO implement me
	panic("implement me")
}
