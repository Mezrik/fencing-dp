package mappers

import (
	"github.com/Mezrik/fencing-dp/internal/core/entities"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/models"
)

func ToCompetitionCategoryModel(competitionCategory *entities.CompetitionCategory) *models.CompetitionCategoryModel {
	return &models.CompetitionCategoryModel{
		Model: models.Model{
			ID: competitionCategory.ID,
		},
		Name: competitionCategory.Name,
	}
}

func ToCompetitionCategoryEntity(competitionCategoryModel *models.CompetitionCategoryModel) *entities.CompetitionCategory {
	return &entities.CompetitionCategory{
		Entity: entities.Entity{ID: competitionCategoryModel.ID},
		Name:   competitionCategoryModel.Name,
	}
}
