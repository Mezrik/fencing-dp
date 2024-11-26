package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetCompetitor struct {
	ID uuid.UUID
}

type GetCompetitorHandler decorator.QueryHandler[GetCompetitor, *Competitor]

type getCompetitorHandler struct {
	repo     repositories.CompetitorRepo
	clubRepo repositories.ClubRepo
}

func NewGetCompetitorHandler(repo repositories.CompetitorRepo, clubRepo repositories.ClubRepo, logger *logrus.Entry) GetCompetitorHandler {
	if repo == nil || clubRepo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[GetCompetitor, *Competitor](getCompetitorHandler{repo, clubRepo}, logger)
}

func (h getCompetitorHandler) Handle(ctx context.Context, payload GetCompetitor) (*Competitor, error) {
	competitor, err := h.repo.FindById(payload.ID)

	if err != nil {
		return nil, err
	}

	return ToCompetitorQueryFromEntity(competitor), nil
}
