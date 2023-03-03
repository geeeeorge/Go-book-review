package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Client) GetApiHealthz(ec echo.Context) error {
	return ec.JSON(http.StatusOK, "Hello")
}
