package mappers

import (
	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/models"
)

func ToCompetitionModel(competition *entities.Competition) *models.CompetitionModel {
	return &models.CompetitionModel{
		Model:      common.Model{ID: competition.ID, CreatedAt: competition.CreatedAt, UpdatedAt: competition.UpdatedAt},
		Name:       competition.Name,
		CategoryID: competition.Category.ID,
		WeaponID:   competition.Weapon.ID,
	}
}

func ToCompetitionEntity(model *models.CompetitionModel) *entities.Competition {
	return &entities.Competition{
		Entity:   common.Entity{ID: model.ID, CreatedAt: model.CreatedAt, UpdatedAt: model.UpdatedAt},
		Name:     model.Name,
		Category: *ToCompetitionCategoryEntity(&model.Category),
		Weapon:   *ToWeaponEntity(&model.Weapon),
	}
}
