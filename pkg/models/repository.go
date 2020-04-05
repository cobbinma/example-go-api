package models

import "context"

type Repository interface {
	CreatePet(ctx context.Context, pet Pet) PetError
}
