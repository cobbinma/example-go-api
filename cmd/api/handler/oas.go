package handler

import "github.com/labstack/echo/v4"

func (h *handler) Oas(c echo.Context) error {
	return c.File("./api/v1.yaml")
}
