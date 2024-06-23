package mappers

import (
	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/models"
)

func ToWeaponModel(weapon *entities.Weapon) *models.WeaponModel {
	return &models.WeaponModel{
		Model: common.Model{ID: weapon.ID, CreatedAt: weapon.CreatedAt, UpdatedAt: weapon.UpdatedAt},
		Name:  weapon.Name,
	}
}

func ToWeaponEntity(weaponModel *models.WeaponModel) *entities.Weapon {
	return &entities.Weapon{
		Entity: common.Entity{ID: weaponModel.ID},
		Name:   weaponModel.Name,
	}
}
