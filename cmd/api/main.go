package main

import (
	"github.com/cobbinma/example-go-api/cmd/api/handler"
	"github.com/cobbinma/example-go-api/config"
	"github.com/cobbinma/example-go-api/models"
	"github.com/cobbinma/example-go-api/repositories/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	p := config.GetPort()

	dbClient, closeDB, err := postgres.NewDBClient()
	if err != nil {
		logrus.Fatalf("could not create database client : %v", err)
	}
	defer func() {
		if err := closeDB(); err != nil {
			logrus.Errorf("could not close database : %v", err)
		}
	}()

	repository := postgres.NewPostgres(dbClient)
	if err := repository.Migrate(); err != nil {
		logrus.Fatalf("could not migrate : %v", err)
	}

	e := getRouter(repository)
	logrus.Infof("Listening for requests on http://localhost:%s/", p)
	e.Logger.Fatal(e.Start(":" + p))
}

func getRouter(repository models.Repository) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/healthz", handler.Health())
	e.GET("/oas", handler.Oas())
	e.POST("/pets", handler.GetPets(repository))
	e.GET("/pets", handler.GetPets(repository))
	e.GET("/pet/:id", handler.GetPet(repository))

	return e
}
