package command

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/common/decorator"
	"github.com/Mezrik/fencing-dp/internal/common/errors"
	"github.com/Mezrik/fencing-dp/internal/competition/domain/repositories"
	"github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	compRepo "github.com/Mezrik/fencing-dp/internal/competitor/domain/repositories"
	matchEntities "github.com/Mezrik/fencing-dp/internal/match/domain/entities"
	matchRepo "github.com/Mezrik/fencing-dp/internal/match/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type InitializeGroups struct {
	CompetitionID uuid.UUID
}

type InitializeGroupsHandler decorator.CommandHandler[InitializeGroups]

type initializeGroupsHandler struct {
	repo            repositories.GroupRepository
	competitionRepo repositories.CompetitionRepository
	competitorRepo  compRepo.CompetitorRepo
	matchRepo       matchRepo.MatchRepository
}

func NewInitializeGroupsHandler(repo repositories.GroupRepository, competitionRepo repositories.CompetitionRepository, competitorRepo compRepo.CompetitorRepo, matchRepo matchRepo.MatchRepository, logger *logrus.Entry) InitializeGroupsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[InitializeGroups](initializeGroupsHandler{repo, competitionRepo, competitorRepo, matchRepo}, logger)
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

		matches := createMatches(groupData.Participants, groupData.Group.ID)

		for _, match := range matches {
			err = h.matchRepo.Create(&match)

			if err != nil {
				return err
			}
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func createMatches(participants []*entities.Participant, groupId uuid.UUID) []matchEntities.Match {
	var matches []matchEntities.Match

	for i := 0; i < len(participants); i++ {
		for j := i + 1; j < len(participants); j++ {
			if participants[i].Competitor().ID == participants[j].Competitor().ID {
				continue
			}

			match := matchEntities.NewMatch(groupId, participants[i].Competitor().ID, participants[j].Competitor().ID)

			matches = append(matches, *match)
		}
	}

	return matches
}
