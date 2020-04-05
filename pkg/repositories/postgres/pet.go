package postgres

import "github.com/cobbinma/example-go-api/pkg/models"

type dbPet struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Tag  string `db:"tag"`
}

func (dbp *dbPet) toPet() *models.Pet {
	return models.NewPet(dbp.ID, dbp.Name, dbp.Tag)
}
