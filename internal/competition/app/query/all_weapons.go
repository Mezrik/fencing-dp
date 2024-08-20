package query

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/sirupsen/logrus"
)

type AllWeapons struct{}

type AllWeaponsHandler decorator.QueryHandler[AllWeapons, []*Weapon]

type allWeaponsHandler struct {
	repo repositories.CompetitionRepository
}

func NewAllWeaponsHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) AllWeaponsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators[AllWeapons, []*Weapon](allWeaponsHandler{repo}, logger)
}

func (h allWeaponsHandler) Handle(ctx context.Context, _ AllWeapons) ([]*Weapon, error) {
	weapons, err := h.repo.FindAllWeapons()

	if err != nil {
		return nil, err
	}

	return ToWeaponQueryListFromEntities(weapons), nil
}
