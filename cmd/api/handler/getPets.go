package handler

import (
	"github.com/cobbinma/example-go-api/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *handler) GetPets(c echo.Context) error {
	ctx := c.Request().Context()

	limit := getLimit(c)

	page := 0

	pets, pErr := h.repository.GetPets(ctx, limit, page)
	if pErr != nil {
		pErr.Wrap("error getting pets")
		logrus.Error(pErr)
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse(pErr))
	}

	return c.JSON(http.StatusOK, pets)
}

func getLimit(c echo.Context) int {
	l := c.QueryParam("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		return 100
	}

	if limit > 100 || limit < 0 {
		return 100
	}

	return limit
}
