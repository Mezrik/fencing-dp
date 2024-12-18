// internal/competition/app/command/update_competition.go

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

type UpdateCompetitionCommand struct {
	ID              uuid.UUID                    `json:"id" ts_type:"UUID"`
	Name            string                       `json:"name"`
	OrganizerName   string                       `json:"organizerName"`
	FederationName  string                       `json:"federationName"`
	Date            time.Time                    `json:"date"`
	CompetitionType entities.CompetitionTypeEnum `json:"competitionType"`
	CategoryID      uuid.UUID                    `json:"categoryId" ts_type:"UUID"`
	Gender          entities.GenderEnum          `json:"gender"`
	WeaponID        uuid.UUID                    `json:"weaponId" ts_type:"UUID"`
}

type UpdateCompetitionHandler decorator.CommandHandler[UpdateCompetitionCommand]

func NewUpdateCompetitionHandler(repo repositories.CompetitionRepository, logger *logrus.Entry) UpdateCompetitionHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[UpdateCompetitionCommand](updateCompetitionHandler{repo}, logger)
}

type updateCompetitionHandler struct {
	repo repositories.CompetitionRepository
}

func (h updateCompetitionHandler) Handle(ctx context.Context, cmd UpdateCompetitionCommand) (err error) {
	defer func() {
		logger.LogCommandExecution("UpdateCompetition", cmd, err)
	}()

	// Retrieve the existing competition
	competition, err := h.repo.FindById(cmd.ID)
	if err != nil {
		return err
	}

	weapon, err := h.repo.FindWeaponById(cmd.WeaponID)
	if err != nil {
		return errors.NewIncorrectInputError("Weapon not found", "weapon-not-found")
	}

	category, err := h.repo.FindCategoryById(cmd.CategoryID)
	if err != nil {
		return errors.NewIncorrectInputError("Category not found", "category-not-found")
	}

	// Update competition fields
	competition.SetName(cmd.Name)
	competition.SetOrganizerName(cmd.OrganizerName)
	competition.SetFederationName(cmd.FederationName)
	competition.SetDate(cmd.Date)
	competition.SetCompetitionType(cmd.CompetitionType)
	competition.SetCategory(*category)
	competition.SetGender(cmd.Gender)
	competition.SetWeapon(*weapon)

	// Save the updated competition
	err = h.repo.Update(competition)
	if err != nil {
		return err
	}

	return nil
}
