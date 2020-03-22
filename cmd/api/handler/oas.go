package handler

import "github.com/labstack/echo/v4"

func (h *handler) Oas(c echo.Context) error {
	return c.File("./files/oas/v1.yaml")
}
