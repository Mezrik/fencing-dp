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
	CompetitionID              uuid.UUID               `json:"competition_id" ts_type:"UUID"`
	ExpectedParticipants       int                     `json:"expected_participants"`
	DeploymentType             entities.DeploymentType `json:"deployment_type"`
	GroupHits                  int                     `json:"group_hits"`
	EliminationHits            int                     `json:"elimination_hits"`
	QualificationBasedOnRounds int                     `json:"qualification_based_on_rounds"`
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

	competition.SetParameters(parameters, len(competition.GroupRounds()))

	// Save the updated competition
	err = h.repo.Update(competition)
	if err != nil {
		return err
	}

	return nil
}
