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

type CreateCompetitor struct {
	Surname    string     `json:"surname"`
	Firstname  string     `json:"firstname"`
	ClubID     *uuid.UUID `json:"clubId" ts_type:"UUID"`
	Gender     string     `json:"gender"`
	License    *string    `json:"license"`
	LicenseFie *string    `json:"licenseFie"`
	Birthdate  *time.Time `json:"birthdate"`
}

type CreateCompetitorHandler decorator.CommandHandler[CreateCompetitor]

type createCompetitiorHandler struct {
	repo     repositories.CompetitorRepo
	clubRepo repositories.ClubRepo
}

func NewCreateCompetitorHandler(repo repositories.CompetitorRepo, clubRepo repositories.ClubRepo, logger *logrus.Entry) CreateCompetitorHandler {
	if repo == nil || clubRepo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[CreateCompetitor](createCompetitiorHandler{repo, clubRepo}, logger)
}

func (h createCompetitiorHandler) Handle(ctx context.Context, command CreateCompetitor) (err error) {
	var club *entities.Club

	if command.ClubID != nil {
		club, err = h.clubRepo.FindById(*command.ClubID)
		if err != nil {
			return errors.NewIncorrectInputError("Club not found", "club-not-found")
		}
	}

	competitor, err := entities.NewCompetitor(
		command.Firstname,
		command.Surname,
		entities.GenderEnum(command.Gender),
		club,
		command.License,
		command.LicenseFie,
		command.Birthdate,
	)

	if err != nil {
		return err
	}

	if err := h.repo.Create(competitor); err != nil {
		return err
	}

	return nil
}
