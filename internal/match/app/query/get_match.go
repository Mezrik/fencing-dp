package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/match/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetMatch struct {
	ID uuid.UUID
}

type GetMatchHandler decorator.QueryHandler[GetMatch, *MatchDetail]

type getMatchHandler struct {
	repo repositories.MatchRepository
}

func NewGetMatchHandler(repo repositories.MatchRepository, logger *logrus.Entry) GetMatchHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[GetMatch, *MatchDetail](getMatchHandler{repo}, logger)
}

func (h getMatchHandler) Handle(ctx context.Context, query GetMatch) (*MatchDetail, error) {
	match, err := h.repo.FindById(query.ID)
	if err != nil {
		return nil, err
	}

	return ToMatchDetailQueryFromEntity(match), nil
}
