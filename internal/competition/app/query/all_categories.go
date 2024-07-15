package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/sirupsen/logrus"
)

type AllCategories struct{}

type AllCategoriesHandler decorator.QueryHandler[AllCategories, []*CompetitionCategory]

type allCategoriesHandler struct {
	repo repositories.CompetitionRepository
}

func NewAllCategoriesHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) AllCategoriesHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[AllCategories, []*CompetitionCategory](allCategoriesHandler{repo}, logger)
}

func (h allCategoriesHandler) Handle(ctx context.Context, _ AllCategories) ([]*CompetitionCategory, error) {
	categories, err := h.repo.FindAllCategories()

	if err != nil {
		return nil, err
	}

	return ToCategoriesQueryListFromEntities(categories), nil
}
