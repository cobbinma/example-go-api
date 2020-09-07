package handler

import (
	"github.com/labstack/echo/v4"
)

func Oas() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.File("./api/v1.yaml")
	}
}
