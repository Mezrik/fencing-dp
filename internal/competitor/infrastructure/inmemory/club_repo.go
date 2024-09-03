package inmemory

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/util"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory/database/dao"
	"github.com/Mezrik/fencing-dp/internal/competitor/infrastructure/inmemory/database/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InMemoryClubRepo struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewInMemoryClubRepo(ctx context.Context, db *sqlx.DB) *InMemoryClubRepo {

	return &InMemoryClubRepo{db: db}
}

func (repo InMemoryClubRepo) Create(club *entities.Club) error {
	dao := &dao.ClubDao{DB: repo.db}

	clubModel, err := repo.marshalClub(club)

	if err != nil {
		return err
	}

	return dao.Create(*clubModel)
}

func (repo InMemoryClubRepo) FindById(id uuid.UUID) (*entities.Club, error) {
	dao := &dao.ClubDao{DB: repo.db}

	clubModel, err := dao.FindById(id)

	if err != nil {
		return nil, err
	}

	club := entities.UnmarshalClub(clubModel.ID, clubModel.Name, clubModel.CreatedAt, util.GetTimePtr(clubModel.UpdatedAt))

	return club, nil
}

func (repo InMemoryClubRepo) FindAll() ([]*entities.Club, error) {
	dao := &dao.ClubDao{DB: repo.db}

	clubModels, err := dao.FindAll()

	if err != nil {
		return nil, err
	}

	clubList := make([]*entities.Club, 0, len(clubModels))

	for _, c := range clubModels {
		clubList = append(clubList, entities.UnmarshalClub(c.ID, c.Name, c.CreatedAt, util.GetTimePtr(c.UpdatedAt)))
	}
	return clubList, nil
}

func (repo InMemoryClubRepo) marshalClub(club *entities.Club) (*models.ClubModel, error) {

	return &models.ClubModel{
		ID:   club.ID,
		Name: club.Name(),
	}, nil
}
