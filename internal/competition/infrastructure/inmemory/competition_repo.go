package repositories

import (
	"context"
	"database/sql"

	models "github.com/Mezrik/fencing-dp/internal/common/database/generated"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type InMemoryCompetitionRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewInMemoryCompetitionRepository(ctx context.Context, db *sql.DB) repositories.CompetitionRepository {
	return &InMemoryCompetitionRepository{db: db}
}

func (repo InMemoryCompetitionRepository) Create(competition *entities.Competition) error {
	dbCompetition, err := repo.marshalCompetition(competition)

	if err != nil {
		return err
	}

	return dbCompetition.Insert(repo.ctx, repo.db, boil.Infer())
}

func (repo InMemoryCompetitionRepository) FindById(id uuid.UUID) (*entities.Competition, error) {
	dbCompetition, err := models.FindCompetition(repo.ctx, repo.db, id)

	if err != nil {
		return nil, err
	}

	return repo.unmarshalCompetition(dbCompetition), nil
}

func (repo InMemoryCompetitionRepository) FindAll() ([]*entities.Competition, error) {
	dbCompetitions, err := models.Competitions().All(repo.ctx, repo.db)

	if err != nil {
		return nil, err
	}

	competitions := make([]*entities.Competition, len(dbCompetitions))
	for i, dbCompetition := range dbCompetitions {
		competitions[i] = repo.unmarshalCompetition(dbCompetition)
	}

	return competitions, nil
}

func (repo InMemoryCompetitionRepository) FindCategoryById(id uuid.UUID) (*entities.CompetitionCategory, error) {
	category, err := models.FindCompetitionCategory(repo.ctx, repo.db, id)

	if err != nil {
		return nil, err
	}

	return entities.UnmarshalCompetitionCategory(
		uuid.UUID(category.ID),
		category.Name, category.CreatedAt,
		category.UpdatedAt.Time), nil
}

func (repo InMemoryCompetitionRepository) FindAllCategories() ([]*entities.CompetitionCategory, error) {
	dbCategories, err := models.CompetitionCategories().All(repo.ctx, repo.db)

	if err != nil {
		return nil, err
	}

	categories := make([]*entities.CompetitionCategory, len(dbCategories))
	for i, dbCategory := range dbCategories {
		categories[i] = entities.UnmarshalCompetitionCategory(
			uuid.UUID(dbCategory.ID),
			dbCategory.Name,
			dbCategory.CreatedAt,
			dbCategory.UpdatedAt.Time)
	}

	return categories, nil
}

func (repo InMemoryCompetitionRepository) FindWeaponById(id uuid.UUID) (*entities.Weapon, error) {
	weapon, err := models.FindWeapon(repo.ctx, repo.db, id)

	if err != nil {
		return nil, err
	}

	return entities.UnmarshalWeapon(
		uuid.UUID(weapon.ID),
		weapon.Name,
		weapon.CreatedAt,
		weapon.UpdatedAt.Time), nil
}

func (repo InMemoryCompetitionRepository) FindAllWeapons() ([]*entities.Weapon, error) {
	dbWeapons, err := models.Weapons().All(repo.ctx, repo.db)

	if err != nil {
		return nil, err
	}

	weapons := make([]*entities.Weapon, len(dbWeapons))
	for i, dbWeapon := range dbWeapons {
		weapons[i] = entities.UnmarshalWeapon(
			uuid.UUID(dbWeapon.ID),
			dbWeapon.Name,
			dbWeapon.CreatedAt,
			dbWeapon.UpdatedAt.Time)
	}

	return weapons, nil
}

func (repo InMemoryCompetitionRepository) unmarshalCompetition(m *models.Competition) *entities.Competition {
	competitionCategory := entities.UnmarshalCompetitionCategory(uuid.UUID(m.R.Category.ID), m.R.Category.Name, m.R.Category.CreatedAt, m.R.Category.UpdatedAt.Time)

	weapon := entities.UnmarshalWeapon(uuid.UUID(m.R.Weapon.ID), m.R.Weapon.Name, m.R.Weapon.CreatedAt, m.R.Weapon.UpdatedAt.Time)

	return entities.UnmarshalCompetition(uuid.UUID(m.ID), m.CreatedAt, m.UpdatedAt.Time, m.Name, m.OrganizerName, m.FederationName, entities.CompetitionTypeEnum(m.CompetitionType), *competitionCategory, entities.GenderEnum(m.Gender), *weapon, m.Date)
}

func (repo InMemoryCompetitionRepository) marshalCompetition(c *entities.Competition) (*models.Competition, error) {
	competitionModel := &models.Competition{
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
