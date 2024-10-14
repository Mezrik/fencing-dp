package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AllParticipants struct {
	CompetitionId uuid.UUID
}

type AllParticipantsHandler decorator.QueryHandler[AllParticipants, []*Participant]

type allParticipantsHandler struct {
	repo repositories.CompetitorRepo
}

func NewAllParticipantsHandler(repo repositories.CompetitorRepo, logger *logrus.Entry) AllParticipantsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[AllParticipants, []*Participant](allParticipantsHandler{repo}, logger)
}

func (h allParticipantsHandler) Handle(ctx context.Context, query AllParticipants) ([]*Participant, error) {
	participants, err := h.repo.FindAllByCompetitionId(query.CompetitionId)

	if err != nil {
		return nil, err
	}

	return ToParticipantQueryListFromEntities(participants), nil
}
