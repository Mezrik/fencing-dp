package command

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AssignParticipants struct {
	CompetitionId  uuid.UUID
	ParticipantIds []uuid.UUID
}

type AssignParticipantsHandler decorator.CommandHandler[AssignParticipants]

type assignParticipantsHandler struct {
	repo repositories.CompetitorRepo
}

func NewAssignParticipantsHandler(repo repositories.CompetitorRepo, logger *logrus.Entry) AssignParticipantsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[AssignParticipants](assignParticipantsHandler{repo}, logger)
}

func (h assignParticipantsHandler) Handle(ctx context.Context, cmd AssignParticipants) error {
	// _, err := h.competitionService.Queries.GetCompetition.Handle(ctx, GetCompetition{
	// 	ID: cmd.CompetitionId,
	// })

	// if err != nil {
	// 	return errors.NewIncorrectInputError("Competition not found", "competition-not-found")
	// }

	return h.repo.AssignCompetitors(cmd.ParticipantIds, cmd.CompetitionId)
}
