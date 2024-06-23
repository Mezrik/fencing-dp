package mappers

import (
	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/models"
)

func ToCompetitionCategoryModel(competitionCategory *entities.CompetitionCategory) *models.CompetitionCategoryModel {
	return &models.CompetitionCategoryModel{
		Model: common.Model{
			ID: competitionCategory.ID,
		},
		Name: competitionCategory.Name,
	}
}

func ToCompetitionCategoryEntity(competitionCategoryModel *models.CompetitionCategoryModel) *entities.CompetitionCategory {
	return &entities.CompetitionCategory{
		Entity: common.Entity{ID: competitionCategoryModel.ID},
		Name:   competitionCategoryModel.Name,
	}
}
