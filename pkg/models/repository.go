package models

import "context"

//go:generate mockgen -package=mock_models -destination=./mock/repository.go -source=repository.go
type Repository interface {
	CreatePet(ctx context.Context, pet *Pet) PetError
	GetPets(ctx context.Context, limit int, page int) ([]*Pet, PetError)
}
