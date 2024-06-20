package mappers

import (
	"github.com/Mezrik/fencing-dp/internal/core/entities"
	"github.com/Mezrik/fencing-dp/internal/infrastructure/models"
)

func ToWeaponModel(weapon *entities.Weapon) *models.WeaponModel {
	return &models.WeaponModel{
		Model: models.Model{ID: weapon.ID, CreatedAt: weapon.CreatedAt, UpdatedAt: weapon.UpdatedAt},
		Name:  weapon.Name,
	}
}

func ToWeaponEntity(weaponModel *models.WeaponModel) *entities.Weapon {
	return &entities.Weapon{
		Entity: entities.Entity{ID: weaponModel.ID},
		Name:   weaponModel.Name,
	}
}
