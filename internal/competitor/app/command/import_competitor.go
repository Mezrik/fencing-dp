package command

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/common/errors"
	"github.com/Mezrik/fencing-dp/internal/common/util"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ImportCompetitor struct {
	File multipart.File
}

type ImportCompetitorHandler decorator.CommandHandler[ImportCompetitor]

type importCompetitiorHandler struct {
	repo     repositories.CompetitorRepo
	clubRepo repositories.ClubRepo
}

func NewImportCompetitorHandler(repo repositories.CompetitorRepo, clubRepo repositories.ClubRepo, logger *logrus.Entry) ImportCompetitorHandler {
	if repo == nil || clubRepo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[ImportCompetitor](importCompetitiorHandler{repo, clubRepo}, logger)
}

func (h importCompetitiorHandler) Handle(ctx context.Context, command ImportCompetitor) (err error) {
	o := util.NewCsvScanner(command.File)

	competitors := []*entities.Competitor{}

	for o.Scan() {
		var club *entities.Club

		if len(o.Text("ClubID")) > 0 {
			club, err = h.clubRepo.FindById(uuid.MustParse(o.Text("ClubID")))
			if err != nil {
				return errors.NewIncorrectInputError("Club not found", "club-not-found")
			}
		}

		license := o.Text("License")
		licenseFie := o.Text("LicenseFie")

		var birthdate time.Time

		if len(o.Text("Birthdate")) > 0 {
			birthdate, err = time.Parse(time.RFC3339, o.Text("Birthdate"))

			if err != nil {
				return err
			}
		}

		var gender string

		if len(o.Text("Gender")) > 0 {
			gender = o.Text("Gender")
		} else {
			gender = "mixed"
		}

		competitor, err := entities.NewCompetitor(
			o.Text("Firstname"),
			o.Text("Surname"),
			entities.GenderEnum(gender),
			club,
			&license,
			&licenseFie,
			&birthdate,
		)

		if err != nil {
			return err
		}

		competitors = append(competitors, competitor)
	}

	err = h.repo.BatchCreate(competitors)

	if err != nil {
		return err
	}

	return nil
}
