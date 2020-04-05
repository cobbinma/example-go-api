package models

type Pet interface {
	GetID() int
	GetName() string
	GetAnimal() string
}
