package models

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type CompetitionModel struct {
	common.Model
	Name            string
	Category        CompetitionCategoryModel
	CategoryID      uuid.UUID
	Weapon          WeaponModel
	WeaponID        uuid.UUID
	OrganizerName   string
	FederationName  string
	Date            time.Time
	CompetitionType entities.CompetitionTypeEnum
	Gender          entities.GenderEnum
}
