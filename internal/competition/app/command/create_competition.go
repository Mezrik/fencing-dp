package command

import (
	"context"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/common/logger"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/entities"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/sirupsen/logrus"
)

type CreateCompetition struct {
	Name           string
	OrganizerName  string
	FederationName string
	Date           time.Time
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

	compt, err := entities.NewCompetition(cmd.Name)

	if err != nil {
		return err
	}

	if err := h.repo.Create(compt); err != nil {
		return err
	}

	return nil
}
