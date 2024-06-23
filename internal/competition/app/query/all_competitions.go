package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/sirupsen/logrus"
)

type AllCompetitions struct{}

type AllCompetitionsHandler decorator.QueryHandler[AllCompetitions, []*Competition]

type allCompetitionsHandler struct {
	repo repositories.CompetitionRepository
}

func NewAllCompetitionsHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) AllCompetitionsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[AllCompetitions, []*Competition](allCompetitionsHandler{repo}, logger)
}

func (h allCompetitionsHandler) Handle(ctx context.Context, _ AllCompetitions) ([]*Competition, error) {
	competitions, err := h.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return ToCompetitionQueryListFromEntities(competitions), nil
}
