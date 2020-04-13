package handler

import "github.com/cobbinma/example-go-api/models"

type handler struct {
	repository models.Repository
}

func NewHandler(repo models.Repository) *handler {
	return &handler{repository: repo}
}
