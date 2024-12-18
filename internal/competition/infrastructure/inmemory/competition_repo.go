package repositories

import (
	"context"
	"strings"

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
	dao := &dao.CompetitionDao{DB: repo.db}

	c, err := dao.FindById(id)

	if err != nil {
		return nil, err
	}

	category := entities.UnmarshalCompetitionCategory(c.Category.ID, c.Category.Name, c.Category.CreatedAt, util.GetTimePtr(c.Category.UpdatedAt))
	weapon := entities.UnmarshalWeapon(c.Weapon.ID, c.Weapon.Name, c.Weapon.CreatedAt, util.GetTimePtr(c.Weapon.UpdatedAt))

	competition := entities.UnmarshalCompetition(
		c.ID,
		c.CreatedAt,
		util.GetTimePtr(c.UpdatedAt),
		c.Name,
		c.OrganizerName,
		c.FederationName,
		entities.CompetitionTypeEnum(c.CompetitionType),
		*category,
		entities.GenderEnum(c.Gender),
		*weapon,
		c.Date,
		nil,
		[]*entities.GroupRound{},
	)

	return competition, nil
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

		competitions = append(competitions, entities.UnmarshalCompetition(
			c.ID,
			c.CreatedAt,
			util.GetTimePtr(c.UpdatedAt),
			c.Name,
			c.OrganizerName,
			c.FederationName,
			entities.CompetitionTypeEnum(c.CompetitionType),
			*category,
			entities.GenderEnum(c.Gender),
			*weapon,
			c.Date,
			nil,
			[]*entities.GroupRound{},
		))
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
	dao := &dao.CompetitionCategoryDao{DB: repo.db}

	competitionCategoriesModels, err := dao.FindAll()

	if err != nil {
		return nil, err
	}

	competitionCategories := make([]*entities.CompetitionCategory, 0, len(competitionCategoriesModels))

	for _, c := range competitionCategoriesModels {
		competitionCategories = append(competitionCategories, entities.UnmarshalCompetitionCategory(c.ID, c.Name, c.CreatedAt, util.GetTimePtr(c.UpdatedAt)))
	}

	return competitionCategories, nil
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
	dao := &dao.WeaponDao{DB: repo.db}

	weaponsModels, err := dao.FindAll()

	if err != nil {
		return nil, err
	}

	weapons := make([]*entities.Weapon, 0, len(weaponsModels))

	for _, w := range weaponsModels {
		weapons = append(weapons, entities.UnmarshalWeapon(w.ID, w.Name, w.CreatedAt, util.GetTimePtr(w.UpdatedAt)))
	}

	return weapons, nil
}

func (repo InMemoryCompetitionRepository) FindAllGroups(competitionId uuid.UUID) ([]*entities.Group, error) {
	dao := &dao.CompetitionGroupDao{DB: repo.db}

	competitionGroupModels, err := dao.FindAll(competitionId)

	if err != nil {
		return nil, err
	}

	competitionGroups := make([]*entities.Group, 0, len(competitionGroupModels))

	for _, c := range competitionGroupModels {
		competitionGroups = append(competitionGroups, entities.UnmarshalCompetitionGroup(c.ID, c.Name, util.GetInt64Ptr(c.PisteNumber), c.CreatedAt, util.GetTimePtr(c.UpdatedAt), competitionId))
	}

	return competitionGroups, nil
}

func (repo InMemoryCompetitionRepository) FindGroupById(id uuid.UUID) (*entities.Group, error) {
	dao := &dao.CompetitionGroupDao{DB: repo.db}

	competitionGroupModel, err := dao.FindById(id)

	if err != nil {
		return nil, err
	}

	competitionGroup := entities.UnmarshalCompetitionGroup(
		competitionGroupModel.ID,
		competitionGroupModel.Name,
		util.GetInt64Ptr(competitionGroupModel.PisteNumber),
		competitionGroupModel.CreatedAt,
		util.GetTimePtr(competitionGroupModel.UpdatedAt),
		competitionGroupModel.CompetitionId,
	)

	return competitionGroup, nil
}

func (repo InMemoryCompetitionRepository) Update(competition *entities.Competition) error {
	compDao := &dao.CompetitionDao{DB: repo.db}
	competitionParametersDao := &dao.CompetitionParametersDao{DB: repo.db}
	competitionGroupRoundDao := &dao.CompetitionGroupRoundDao{DB: repo.db}

	competitionModel, err := repo.marshalCompetition(competition)

	if err != nil {
		return err
	}

	parameters, err := competitionParametersDao.FindById(competitionModel.Parameters.ID)

	if err != nil {
		return err
	}

	if parameters == nil {
		err = competitionParametersDao.Create(&competitionModel.Parameters)
	} else {
		err = competitionParametersDao.Update(&competitionModel.Parameters)
	}

	if err != nil {
		return err
	}

	for _, round := range competition.GroupRounds() {
		roundModel := models.CompetitionGroupRoundModel{
			ID:                       round.ID,
			CreatedAt:                round.CreatedAt,
			Number:                   round.Number(),
			CompetitionID:            round.CompetitionID(),
			ParticipantStartingCount: round.ParticipantsStartingCount(),
			NumberOfGroups:           round.NumberOfGroups(),
			ParticipantsInGroups:     round.ParticipantsInGroups(),
			ShiftCriteria:            strings.Join(shiftCriteriaToStrings(round.ShiftCriteria()), ","),
			NumberOfAdvancers:        round.NumberOfAdvancers(),
		}

		err = competitionGroupRoundDao.Create(&roundModel)
		if err != nil {
			return err
		}
	}

	return compDao.Update(competitionModel)
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

func shiftCriteriaToStrings(criteria []entities.ShiftCriteria) []string {
	strs := make([]string, len(criteria))
	for i, c := range criteria {
		strs[i] = string(c)
	}
	return strs
}
