package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetCompetition struct {
	ID uuid.UUID
}

type GetCompetitionHandler decorator.QueryHandler[GetCompetition, *Competition]

type getCompetitionHandler struct {
	repo repositories.CompetitionRepository
}

func NewGetCompetitionHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) GetCompetitionHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[GetCompetition, *Competition](getCompetitionHandler{repo}, logger)
}

func (h getCompetitionHandler) Handle(ctx context.Context, query GetCompetition) (*Competition, error) {
	competition, err := h.repo.FindById(query.ID)

	if err != nil {
		return nil, err
	}

	return ToCompetitionQueryFromEntity(competition), nil
}
