package postgres

import (
	"context"
	"github.com/cobbinma/example-go-api/pkg/models"
)

type postgres struct {
	dbClient DBClient
}

func NewPostgres(client DBClient) *postgres {
	return &postgres{dbClient: client}
}

func (p *postgres) CreatePet(ctx context.Context, pet models.Pet) models.PetError {

}
