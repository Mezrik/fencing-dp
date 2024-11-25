package command

import (
	"context"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/common/errors"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateCompetitor struct {
	ID         uuid.UUID  `json:"id" ts_type:"UUID"`
	Surname    string     `json:"surname"`
	Firstname  string     `json:"firstname"`
	ClubID     *uuid.UUID `json:"clubId" ts_type:"UUID"`
	Gender     string     `json:"gender"`
	License    *string    `json:"license"`
	LicenseFie *string    `json:"licenseFie"`
	Birthdate  *time.Time `json:"birthdate"`
}

type UpdateCompetitorHandler decorator.CommandHandler[UpdateCompetitor]

type updateCompetitiorHandler struct {
	repo     repositories.CompetitorRepo
	clubRepo repositories.ClubRepo
}

func NewUpdateCompetitorHandler(repo repositories.CompetitorRepo, clubRepo repositories.ClubRepo, logger *logrus.Entry) UpdateCompetitorHandler {
	if repo == nil || clubRepo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[UpdateCompetitor](updateCompetitiorHandler{repo, clubRepo}, logger)
}

func (h updateCompetitiorHandler) Handle(ctx context.Context, command UpdateCompetitor) (err error) {
	var club *entities.Club

	if command.ClubID != nil {
		club, err = h.clubRepo.FindById(*command.ClubID)
		if err != nil {
			return errors.NewIncorrectInputError("Club not found", "club-not-found")
		}
	}

	competitor, err := h.repo.FindById(command.ID)
	if err != nil {
		return err
	}

	competitor.SetFirstname(command.Firstname)
	competitor.SetSurname(command.Surname)
	competitor.SetGender(entities.GenderEnum(command.Gender))
	competitor.SetClub(club)
	competitor.SetLicense(command.License)
	competitor.SetLicenseFie(command.LicenseFie)
	competitor.SetBirthdate(command.Birthdate)

	competitor.SetUpdatedAt(time.Now())

	if err := h.repo.Create(competitor); err != nil {
		return err
	}

	return nil
}
