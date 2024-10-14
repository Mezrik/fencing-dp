package command

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AssignParticipant struct {
	CompetitionId uuid.UUID
	ParticipantId uuid.UUID
}

type AssignParticipantHandler decorator.CommandHandler[AssignParticipant]

type assignParticipantHandler struct {
	repo repositories.CompetitorRepo
}

func NewAssignParticipantHandler(repo repositories.CompetitorRepo, logger *logrus.Entry) AssignParticipantHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[AssignParticipant](assignParticipantHandler{repo}, logger)
}

func (h assignParticipantHandler) Handle(ctx context.Context, cmd AssignParticipant) error {
	// _, err := h.competitionService.Queries.GetCompetition.Handle(ctx, GetCompetition{
	// 	ID: cmd.CompetitionId,
	// })

	// if err != nil {
	// 	return errors.NewIncorrectInputError("Competition not found", "competition-not-found")
	// }

	return h.repo.AssignCompetitor(cmd.ParticipantId, cmd.CompetitionId)
}
