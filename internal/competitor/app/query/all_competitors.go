package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/sirupsen/logrus"
)

type AllCompetitors struct{}

type AllCompetitorsHandler decorator.QueryHandler[AllCompetitors, []*Competitor]

type allCompetitorHandler struct {
	repo     repositories.CompetitorRepo
	clubRepo repositories.ClubRepo
}

func NewAllCompetitionsHandler(repo repositories.CompetitorRepo, clubRepo repositories.ClubRepo, logger *logrus.Entry) AllCompetitorsHandler {
	if repo == nil || clubRepo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[AllCompetitors, []*Competitor](allCompetitorHandler{repo, clubRepo}, logger)
}

func (h allCompetitorHandler) Handle(ctx context.Context, _ AllCompetitors) ([]*Competitor, error) {
	competitors, err := h.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return ToCompetitorQueryListFromEntities(competitors), nil
}
