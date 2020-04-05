package handler

import (
	"github.com/cobbinma/example-go-api/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func (h *handler) CreatePet(c echo.Context) error {
	ctx := c.Request().Context()

	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		pErr := models.NewPetError(err, "could not read request", 0)
		pErr.Wrap("error reading request body")
		logrus.Info(pErr)
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse(pErr))
	}

	pet, pErr := models.NewPetFromRequest(reqBody)
	if pErr != nil {
		pErr.Wrap("error creating pet from request")
		logrus.Info(pErr)
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse(pErr))
	}

	pErr = h.repository.CreatePet(ctx, pet)
	if pErr != nil {
		pErr.Wrap("error storing pet")
		logrus.Error(pErr)
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse(pErr))
	}

	return c.NoContent(http.StatusCreated)
}
