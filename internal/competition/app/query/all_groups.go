package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AllGroups struct {
	CompetitionID uuid.UUID
}

type AllGroupsHandler decorator.QueryHandler[AllGroups, []*Group]

type allGroupsHandler struct {
	repo repositories.CompetitionRepository
}

func NewAllGroupsHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) AllGroupsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[AllGroups, []*Group](allGroupsHandler{repo}, logger)
}

func (h allGroupsHandler) Handle(ctx context.Context, query AllGroups) ([]*Group, error) {
	groups, err := h.repo.FindAllGroups(query.CompetitionID)

	if err != nil {
		return nil, err
	}

	return ToGroupQueryListFromEntities(groups), nil
}
