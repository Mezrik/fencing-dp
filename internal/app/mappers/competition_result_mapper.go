package mappers

import (
	"github.com/Mezrik/fencing-dp/internal/app/common"
	"github.com/Mezrik/fencing-dp/internal/core/entities"
)

func ToCompetitionResultFromEntity(productResult *entities.Competition) *common.CompetitionResult {
	return &common.CompetitionResult{
		Result: common.Result{Id: productResult.ID, CreatedAt: productResult.CreatedAt, UpdatedAt: productResult.UpdatedAt},
		Name:   productResult.Name,
	}
}
