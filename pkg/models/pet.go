package models

import (
	"encoding/json"
	"errors"
)

type Pet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag,omitempty"`
}

func NewPet(id int, name string, tag string) *Pet {
	return &Pet{
		ID:   id,
		Name: name,
		Tag:  tag,
	}
}

func NewPetFromRequest(req []byte) (*Pet, PetError) {
	pet := Pet{}
	if err := json.Unmarshal(req, &pet); err != nil {
		pErr := NewPetError(err, "request is not valid", 0)
		pErr.Wrap("could not unmarshal request into Pet")
		return nil, pErr
	}
	if err := pet.validate(); err != nil {
		pErr := NewPetError(err, "could not validate request", 0)
		pErr.Wrap("could not validate request")
		return nil, pErr
	}
	return &pet, nil
}

func (p *Pet) validate() PetError {
	if err := p.isIDValid(); err != nil {
		pe := NewPetError(err, "ID is not valid", 0)
		pe.Wrap("could not validate ID")
		return pe
	}
	if err := p.isNameValid(); err != nil {
		pe := NewPetError(err, "name is not valid", 0)
		pe.Wrap("could not validate name")
		return pe
	}
	return nil
}

func (p *Pet) isIDValid() error {
	if p.ID < 1 {
		return errors.New("ID must be greater than 0")
	}
	return nil
}

func (p *Pet) isNameValid() error {
	if p.Name == "" {
		return errors.New("name must not be empty")
	}
	return nil
}
