package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Health() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}
