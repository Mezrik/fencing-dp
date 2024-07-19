package command

import (
	"context"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/common/errors"
	"github.com/Mezrik/fencing-dp/internal/common/logger"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateCompetition struct {
	Name            string                       `json:"name"`
	OrganizerName   string                       `json:"organizerName"`
	FederationName  string                       `json:"federationName"`
	Date            time.Time                    `json:"date"`
	CompetitionType entities.CompetitionTypeEnum `json:"competitionType"`
	CategoryID      uuid.UUID                    `json:"categoryId" ts_type:"UUID"`
	Gender          entities.GenderEnum          `json:"gender"`
	WeaponID        uuid.UUID                    `json:"weaponId" ts_type:"UUID"`
}

type CreateCompetitionHandler decorator.CommandHandler[CreateCompetition]

type createCompetitionHandler struct {
	repo repositories.CompetitionRepository
}

func NewCreateCompetitionHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) CreateCompetitionHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[CreateCompetition](createCompetitionHandler{repo}, logger)
}

func (h createCompetitionHandler) Handle(ctx context.Context, cmd CreateCompetition) (err error) {
	defer func() {
		logger.LogCommandExecution("CreateCompetition", cmd, err)
	}()

	if cmd.Name == "" {
		return errors.NewIncorrectInputError("Name must not be empty", "empty-name")
	}

	category, err := h.repo.FindCategoryById(cmd.CategoryID)
	if err != nil {
		return errors.NewIncorrectInputError("Category not found", "category-not-found")
	}

	weapon, err := h.repo.FindWeaponById(cmd.WeaponID)
	if err != nil {
		return errors.NewIncorrectInputError("Weapon not found", "weapon-not-found")
	}

	compt, err := entities.NewCompetition(
		cmd.Name,
		cmd.OrganizerName,
		cmd.FederationName,
		cmd.CompetitionType,
		*category,
		cmd.Gender,
		*weapon,
		cmd.Date,
	)

	if err != nil {
		return err
	}

	if err := h.repo.Create(compt); err != nil {
		return err
	}

	return nil
}
