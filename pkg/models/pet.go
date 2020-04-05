package models

import (
	"encoding/json"
	"errors"
)

type Pet interface {
	GetID() int
	GetName() string
	GetTag() string
}

type pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag,omitempty"`
}

func NewPet(req []byte) (Pet, PetError) {
	pet := pet{}
	if err := json.Unmarshal(req, &pet); err != nil {
		pe := newPetError(err, "request is not valid", 0)
		pe.Wrap("could not unmarshal request into pet")
		return nil, pe
	}
	return &pet, nil
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

func (p *pet) validate() PetError {
	if err := p.isIDValid(); err != nil {
		pe := newPetError(err, "ID is not valid", 0)
		pe.Wrap("could not validate ID")
		return pe
	}
	if err := p.isNameValid(); err != nil {
		pe := newPetError(err, "name is not valid", 0)
		pe.Wrap("could not validate name")
		return pe
	}
	return nil
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
