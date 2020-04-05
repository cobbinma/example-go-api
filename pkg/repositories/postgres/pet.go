package postgres

type pet struct {
	Identifier int    `db:"identifier"`
	Name       string `db:"name"`
	Animal     string `db:"animal"`
}

func (p *pet) GetID() int {
	return p.Identifier
}

func (p *pet) GetName() string {
	return p.Name
}

func (p *pet) GetAnimal() string {
	return p.Animal
}
