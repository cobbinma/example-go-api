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

func (p *postgres) CreatePet(ctx context.Context, pet *models.Pet) models.PetError {
	sql, args, _ := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("pets").Columns("id", "name", "tag").
		Values(pet.ID, pet.Name, pet.Tag).ToSql()

	_, err := p.dbClient.Exec(sql, args...)
	if err != nil {
		pErr := newPetError(err, "could not store pet", 0)
		pErr.Wrap("could not store pet in database")
		return pErr
	}

	return nil
}

func (p *postgres) GetPets(ctx context.Context, limit int, page int) ([]*models.Pet, models.PetError) {
	sql, args, _ := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id", "name", "tag").From("pets").
		Limit(uint64(limit)).Offset(uint64((page) * (limit))).ToSql()

	dbPets, err := p.dbClient.GetPets(sql, args...)
	if err != nil {
		pErr := newPetError(err, "could get pets", 0)
		pErr.Wrap("could not get pets from database")
		return nil, pErr
	}

	return dbPets.toPets(), nil
}

func (p *postgres) GetPet(ctx context.Context, id int) (*models.Pet, models.PetError) {
	sql, args, _ := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id", "name", "tag").From("pets").
		Where(sq.Eq{"id": id}).ToSql()

	dbPet, err := p.dbClient.GetPet(sql, args...)
	if err != nil {
		pErr := newPetError(err, "could get pet", 0)
		pErr.Wrap("could not get pet from database")
		return nil, pErr
	}

	return dbPet.toPet(), nil
}
