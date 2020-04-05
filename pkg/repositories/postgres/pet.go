package postgres

import "github.com/cobbinma/example-go-api/pkg/models"

type dbPets []dbPet

func (ds dbPets) toPets() []*models.Pet {
	pets := make([]*models.Pet, len(ds))
	for i := 0; i < len(ds); i++ {
		pets[i] = ds[i].toPet()
	}
	return pets
}

type dbPet struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Tag  string `db:"tag"`
}

func (dbp *dbPet) toPet() *models.Pet {
	return models.NewPet(dbp.ID, dbp.Name, dbp.Tag)
}
