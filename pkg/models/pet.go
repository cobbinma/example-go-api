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
		pe := newPetError(err, "request is not valid", 0)
		pe.Wrap("could not unmarshal request into Pet")
		return nil, pe
	}
	return &pet, nil
}

func (p *Pet) validate() PetError {
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
