package command

import (
	"context"
	"time"

	"github.com/Mezrik/fencing-dp/internal/app/decorator"
	"github.com/Mezrik/fencing-dp/internal/core/entities"
	"github.com/Mezrik/fencing-dp/internal/core/repositories"
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

func NewCreateTrainingHandler(repo repositories.CompetitionRepository) CreateCompetitionHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[CreateCompetition](createCompetitionHandler{repo})
}

func (h createCompetitionHandler) Handle(ctx context.Context, cmd CreateCompetition) error {
	compt, err := entities.NewCompetition(cmd.Name)

	if err != nil {
		return err
	}

	if err := h.repo.Create(compt); err != nil {
		return err
	}

	return nil
}
