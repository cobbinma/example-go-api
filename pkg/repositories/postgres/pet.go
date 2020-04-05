package postgres

type pet struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Tag  string `db:"tag,omitempty"`
}

func (p *pet) GetID() int {
	return p.ID
}

func (p *pet) GetName() string {
	return p.Name
}

func (p *pet) GetTag() string {
	return p.Tag
}
