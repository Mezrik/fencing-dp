package repositories

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/util"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/dao"
	"github.com/Mezrik/fencing-dp/internal/competition/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InMemoryCompetitionRepository struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewInMemoryCompetitionRepository(ctx context.Context, db *sqlx.DB) repositories.CompetitionRepository {
	return &InMemoryCompetitionRepository{db: db}
}

func (repo InMemoryCompetitionRepository) Create(competition *entities.Competition) error {
	dao := &dao.CompetitionDao{DB: repo.db}

	competitionModel, err := repo.marshalCompetition(competition)

	if err != nil {
		return err
	}

	return dao.Create(competitionModel)
}

func (repo InMemoryCompetitionRepository) FindById(id uuid.UUID) (*entities.Competition, error) {
	return nil, nil
}

func (repo InMemoryCompetitionRepository) FindAll() ([]*entities.Competition, error) {
	dao := &dao.CompetitionDao{DB: repo.db}

	competitionsModels, err := dao.FindAll()

	if err != nil {
		return nil, err
	}

	competitions := make([]*entities.Competition, 0, len(competitionsModels))

	for _, c := range competitionsModels {
		category := entities.UnmarshalCompetitionCategory(c.Category.ID, c.Category.Name, c.Category.CreatedAt, util.GetTimePtr(c.Category.UpdatedAt))
		weapon := entities.UnmarshalWeapon(c.Weapon.ID, c.Weapon.Name, c.Weapon.CreatedAt, util.GetTimePtr(c.Weapon.UpdatedAt))

		competitions = append(competitions, entities.UnmarshalCompetition(c.ID, c.CreatedAt, util.GetTimePtr(c.UpdatedAt), c.Name, c.OrganizerName, c.FederationName, entities.CompetitionTypeEnum(c.CompetitionType), *category, entities.GenderEnum(c.Gender), *weapon, c.Date))
	}

	return competitions, nil
}

func (repo InMemoryCompetitionRepository) FindCategoryById(id uuid.UUID) (*entities.CompetitionCategory, error) {
	dao := &dao.CompetitionCategoryDao{DB: repo.db}

	competitionCategoryModel, err := dao.FindById(id)

	if err != nil {
		return nil, err
	}

	competitionCategory := entities.UnmarshalCompetitionCategory(
		competitionCategoryModel.ID,
		competitionCategoryModel.Name,
		competitionCategoryModel.CreatedAt,
		util.GetTimePtr(competitionCategoryModel.UpdatedAt),
	)

	return competitionCategory, nil
}

func (repo InMemoryCompetitionRepository) FindAllCategories() ([]*entities.CompetitionCategory, error) {

	return nil, nil
}

func (repo InMemoryCompetitionRepository) FindWeaponById(id uuid.UUID) (*entities.Weapon, error) {
	dao := &dao.WeaponDao{DB: repo.db}

	weaponModel, err := dao.FindById(id)

	if err != nil {
		return nil, err
	}

	weapon := entities.UnmarshalWeapon(
		weaponModel.ID,
		weaponModel.Name,
		weaponModel.CreatedAt,
		util.GetTimePtr(weaponModel.UpdatedAt),
	)

	return weapon, nil
}

func (repo InMemoryCompetitionRepository) FindAllWeapons() ([]*entities.Weapon, error) {
	return nil, nil
}

func (repo InMemoryCompetitionRepository) marshalCompetition(c *entities.Competition) (*models.CompetitionModel, error) {
	competitionModel := &models.CompetitionModel{
		ID:              c.ID,
		Name:            c.Name(),
		OrganizerName:   c.OrganizerName(),
		FederationName:  c.FederationName(),
		CompetitionType: string(c.CompetitionType()),
		Gender:          string(c.Gender()),
		Date:            c.Date(),
		CategoryID:      c.Category().ID,
		WeaponID:        c.Weapon().ID,
		CreatedAt:       c.CreatedAt,
	}

	return competitionModel, nil
}
