package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/match/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AllMatches struct {
	GroupID uuid.UUID
}

type AllMatchesHandler decorator.QueryHandler[AllMatches, []*Match]

type allMatchesHandler struct {
	repo repositories.MatchRepository
}

func NewAllMatchesHandler(repo repositories.MatchRepository, logger *logrus.Entry) AllMatchesHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[AllMatches, []*Match](allMatchesHandler{repo}, logger)
}

func (h allMatchesHandler) Handle(ctx context.Context, query AllMatches) ([]*Match, error) {
	matches, err := h.repo.FindAll(query.GroupID)
	if err != nil {
		return nil, err
	}

	return ToMatchQueryListFromEntities(matches), nil
}
