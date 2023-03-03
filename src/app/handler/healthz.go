package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *Client) GetApiHealthz(ec echo.Context) error {
	return ec.JSON(http.StatusOK, "Hello")
}
