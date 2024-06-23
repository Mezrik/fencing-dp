package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
)

type AllCompetitions struct{}

type AllCompetitionsHandler decorator.QueryHandler[AllCompetitions, []*Competition]

type allCompetitionsHandler struct {
	repo repositories.CompetitionRepository
}

func NewAllCompetitionsHandler(repo repositories.CompetitionRepository) AllCompetitionsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[AllCompetitions, []*Competition](allCompetitionsHandler{repo})
}

func (h allCompetitionsHandler) Handle(ctx context.Context, _ AllCompetitions) ([]*Competition, error) {
	competitions, err := h.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return ToCompetitionQueryListFromEntities(competitions), nil
}
