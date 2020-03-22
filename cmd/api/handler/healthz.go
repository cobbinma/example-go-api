package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) Health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
