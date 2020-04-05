package models

import "errors"

type Pet interface {
	GetID() int
	GetName() string
	GetTag() string
}

type pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func NewPet(req []byte) (Pet, PetError) {
	return nil, nil
}

func (p *pet) validate() PetError {
	if err := p.isIDValid(); err != nil {
		return
	}
}

func (p *pet) isIDValid() error {
	if p.ID < 1 {
		return errors.New("ID must be greater than 0")
	}
	return nil
}

func (p *pet) isNameValid() error {
	if p.Name == "" {
		return errors.New("name must not be empty")
	}
	return nil
}
