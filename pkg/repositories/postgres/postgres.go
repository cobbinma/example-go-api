package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cobbinma/example-go-api/pkg/models"
)

type postgres struct {
	dbClient DBClient
}

func NewPostgres(client DBClient) *postgres {
	return &postgres{dbClient: client}
}

func (p *postgres) CreatePet(ctx context.Context, pet models.Pet) models.PetError {
	sql, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("pets").Columns("id", "name", "tag").
		Values(pet.ID, pet.Name, pet.Tag).ToSql()

	_, err = p.dbClient.Exec(sql, args...)
	if err != nil {
		pErr := newPetError(err, "could not store pet", 0)
		pErr.Wrap("could not store pet in database")
		return pErr
	}

	return nil
}
