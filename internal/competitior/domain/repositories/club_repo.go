package repositories

import "github.com/Mezrik/fencing-dp/internal/competitior/domain/entities"

type ClubRepo interface {
	Create(competitor *entities.Club) error
	FindAll() ([]*entities.Club, error)
}
