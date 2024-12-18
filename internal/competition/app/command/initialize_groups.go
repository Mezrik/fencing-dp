package command

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/common/errors"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	compRepo "github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	"github.com/google/uuid"
)

type InitializeGroups struct {
	CompetitionID uuid.UUID
}

type InitializeGroupsHandler decorator.CommandHandler[InitializeGroups]

type initializeGroupsHandler struct {
	repo            repositories.GroupRepository
	competitionRepo repositories.CompetitionRepository
	competitorRepo  compRepo.CompetitorRepo
}

func NewInitializeGroupsHandler(repo repositories.GroupRepository, competitionRepo repositories.CompetitionRepository, competitorRepo compRepo.CompetitorRepo) InitializeGroupsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[InitializeGroups](initializeGroupsHandler{repo, competitionRepo, competitorRepo}, nil)
}

func (h initializeGroupsHandler) Handle(ctx context.Context, cmd InitializeGroups) error {
	competition, err := h.competitionRepo.FindById(cmd.CompetitionID)

	if err != nil {
		return errors.NewIncorrectInputError("Competition not found", "competition-not-found")
	}

	var round = 1

	allParticipants, err := h.competitorRepo.FindAllByCompetitionId(cmd.CompetitionID)

	groupsData := competition.GroupRounds()[round-1].CreateGroups(allParticipants)

	for _, groupData := range groupsData {
		err = h.repo.Create(groupData.Group)

		if err != nil {
			return err
		}

		for _, participant := range groupData.Participants {
			err = h.competitorRepo.UpdateParticipant(participant)
		}
	}

	if err != nil {
		return err
	}

	return nil
}
