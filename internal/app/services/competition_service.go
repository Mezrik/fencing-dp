package services

import (
	"github.com/Mezrik/fencing-dp/internal/app/common"
	"github.com/Mezrik/fencing-dp/internal/app/interfaces"
	"github.com/Mezrik/fencing-dp/internal/app/mappers"
	"github.com/Mezrik/fencing-dp/internal/core/repositories"
)

type CompetitionService struct {
	competitionRepository repositories.CompetitionRepository
}

func NewCompetitionService(competitionRepository repositories.CompetitionRepository) interfaces.CompetitionService {
	return &CompetitionService{competitionRepository: competitionRepository}
}

func (s *CompetitionService) GetAllCompetitions() ([]*common.CompetitionResult, error) {
	storedCompetitions, err := s.competitionRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var list []*common.CompetitionResult
	for _, storedCompetition := range storedCompetitions {
		list = append(list, mappers.ToCompetitionResultFromEntity(storedCompetition))
	}

	return list, nil
}
