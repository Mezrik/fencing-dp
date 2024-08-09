package models

import (
	"time"

	"github.com/google/uuid"
)

type CompetitionModel struct {
	ID              uuid.UUID `db:"id`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	Name            string    `db:"name"`
	OrganizerName   string    `db:"organizer_name"`
	FederationName  string    `db:"federation_name"`
	Date            time.Time `db:"date"`
	CompetitionType string    `db:"competition_type"`
	Gender          string    `db:"gender"`
	WeaponID        uuid.UUID `db:"weapon_id"`
	Weapon          WeaponModel
	CategoryID      uuid.UUID `db:"category_id"`
	Category        CompetitionCategoryModel
}
