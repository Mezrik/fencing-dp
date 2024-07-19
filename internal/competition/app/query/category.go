package query

import (
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/google/uuid"
)

type CompetitionCategory struct {
	ID   uuid.UUID `json:"id" ts_type:"UUID"`
	Name string    `json:"name"`
}

func ToCategoryQueryFromEntity(category *entities.CompetitionCategory) *CompetitionCategory {
	return &CompetitionCategory{
		ID:   category.ID,
		Name: category.Name(),
	}
}
func ToCategoriesQueryListFromEntities(categories []*entities.CompetitionCategory) []*CompetitionCategory {
	categoriesList := make([]*CompetitionCategory, 0, len(categories))
	for _, category := range categories {
		categoriesList = append(categoriesList, ToCategoryQueryFromEntity(category))
	}
	return categoriesList
}
