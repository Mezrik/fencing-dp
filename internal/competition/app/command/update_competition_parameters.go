package command

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/common/logger"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateCompetitionParameters struct {
	CompetitionID              uuid.UUID               `json:"competitionId" ts_type:"UUID"`
	ExpectedParticipants       int                     `json:"expectedParticipants"`
	DeploymentType             entities.DeploymentType `json:"deploymentType"`
	GroupHits                  int                     `json:"groupHits"`
	EliminationHits            int                     `json:"eliminationHits"`
	QualificationBasedOnRounds int                     `json:"qualificationBasedOnRounds"`
	RoundsCount                int                     `json:"roundsCount"`
}

type UpdateCompetitionParametersHandler decorator.CommandHandler[UpdateCompetitionParameters]

type updateCompetitionParametersHandler struct {
	repo repositories.CompetitionRepository
}

func NewUpdateCompetitionParametersHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) UpdateCompetitionParametersHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[UpdateCompetitionParameters](updateCompetitionParametersHandler{repo}, logger)
}

func (h updateCompetitionParametersHandler) Handle(ctx context.Context, cmd UpdateCompetitionParameters) (err error) {
	defer func() {
		logger.LogCommandExecution("UpdateCompetitionParameters", cmd, err)
	}()

	// Retrieve the existing competition
	competition, err := h.repo.FindById(cmd.CompetitionID)
	if err != nil {
		return err
	}

	// Update competition parameters
	parameters := entities.NewCompetitionParameters(
		cmd.ExpectedParticipants,
		cmd.DeploymentType,
		cmd.GroupHits,
		cmd.EliminationHits,
		cmd.QualificationBasedOnRounds,
	)

	competition.SetParameters(parameters, cmd.RoundsCount)

	// Save the updated competition
	err = h.repo.Update(competition)
	if err != nil {
		return err
	}

	return nil
}
