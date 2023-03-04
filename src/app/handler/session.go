package handler

import (
	"github.com/geeeeorge/Go-book-review/gen/api"
	"github.com/geeeeorge/Go-book-review/src/app/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *Client) Signup(ec echo.Context) error {
	ctx := ec.Request().Context()

	req := new(api.User)
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, nil)
	}
	u := model.User{
		Username: *req.Username,
		Password: *req.Password,
	}

	if err := c.usecase.SignUp(ctx, &u); err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	// passwordを消して返却
	u.Password = ""

	return ec.JSON(http.StatusCreated, u)
}

func (c *Client) Login(ec echo.Context) error {
	ctx := ec.Request().Context()

	req := new(api.User)
	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, nil)
	}
	u := &model.User{
		Username: *req.Username,
		Password: *req.Password,
	}

	token, err := c.usecase.Login(ctx, u)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}
	ec.Set("username", u.Username)

	return ec.JSON(http.StatusOK, map[string]string{"token": token})
}
