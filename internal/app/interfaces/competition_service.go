package interfaces

import "github.com/Mezrik/fencing-dp/internal/app/common"

type CompetitionService interface {
	GetAllCompetitions() ([]*common.CompetitionResult, error)
}
