package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetGroup struct {
	ID uuid.UUID
}

type GetGroupHandler decorator.QueryHandler[GetGroup, *Group]

type getGroupHandler struct {
	repo repositories.CompetitionRepository
}

func NewGetGroupHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) GetGroupHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[GetGroup, *Group](getGroupHandler{repo}, logger)
}

func (h getGroupHandler) Handle(ctx context.Context, query GetGroup) (*Group, error) {
	group, err := h.repo.FindGroupById(query.ID)

	if err != nil {
		return nil, err
	}

	return ToGroupQueryFromEntity(group), nil
}
