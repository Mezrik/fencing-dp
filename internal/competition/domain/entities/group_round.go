package entities

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	competitorEntities "github.com/Mezrik/fencing-dp/internal/competitor/domain/entities"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type ShiftCriteria string

const (
	Nation ShiftCriteria = "nation"
	Club   ShiftCriteria = "club"
	Rating ShiftCriteria = "rating"
)

type GroupData struct {
	Group        *Group
	Participants []*competitorEntities.Participant
}

type GroupRound struct {
	common.Entity
	competitionID             uuid.UUID
	number                    int
	participantsStartingCount int

	numberOfGroups       int
	participantsInGroups int

	shiftCriteria []ShiftCriteria

	numberOfAdvancers int
}

func NewGroupRound(competitionID uuid.UUID, number int, participantsStartingCount int, shiftCriteria []ShiftCriteria, advancingPercent int) *GroupRound {

	recommendedSize := CalculateGroupSizes(participantsStartingCount, 3, 6)[0]

	return &GroupRound{
		Entity:                    common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		competitionID:             competitionID,
		number:                    number,
		participantsStartingCount: participantsStartingCount,
		numberOfGroups:            recommendedSize.groupsCount,
		participantsInGroups:      slices.Max(recommendedSize.groupSizes),
		shiftCriteria:             shiftCriteria,
		numberOfAdvancers:         participantsStartingCount * advancingPercent / 100,
	}
}

func UnmarshalGroupRound(ID uuid.UUID, competitionID uuid.UUID, number int, participantsStartingCount int, numberOfGroups int, participantsInGroups int, shiftCriteria []ShiftCriteria, numberOfAdvancers int, CreatedAt time.Time, UpdatedAt *time.Time) *GroupRound {
	return &GroupRound{
		Entity:                    common.Entity{ID, CreatedAt, UpdatedAt},
		competitionID:             competitionID,
		number:                    number,
		participantsStartingCount: participantsStartingCount,
		numberOfGroups:            numberOfGroups,
		participantsInGroups:      participantsInGroups,
		shiftCriteria:             shiftCriteria,
		numberOfAdvancers:         numberOfAdvancers,
	}
}

func (c *GroupRound) CompetitionID() uuid.UUID {
	return c.competitionID
}

func (c *GroupRound) Number() int {
	return c.number
}

func (c *GroupRound) ParticipantsStartingCount() int {
	return c.participantsStartingCount
}

func (c *GroupRound) NumberOfGroups() int {
	return c.numberOfGroups
}

func (c *GroupRound) ParticipantsInGroups() int {
	return c.participantsInGroups
}

func (c *GroupRound) ShiftCriteria() []ShiftCriteria {
	return c.shiftCriteria
}

func (c *GroupRound) NumberOfAdvancers() int {
	return c.numberOfAdvancers
}

func (c GroupRound) CreateGroups(participants []*competitorEntities.Participant) []*GroupData {
	groups := make([]*GroupData, 0, c.numberOfGroups)

	for i := 0; i < c.numberOfGroups; i++ {
		gd := &GroupData{
			Group:        NewCompetitionGroup("Group "+string(i+1), c.competitionID),
			Participants: participants[i*c.participantsInGroups : (i+1)*c.participantsInGroups],
		}

		for _, p := range gd.Participants {
			p.SetGroupID(gd.Group.ID)
		}

		groups = append(groups, &GroupData{
			Group:        NewCompetitionGroup("Group "+string(i+1), c.competitionID),
			Participants: participants[i*c.participantsInGroups : (i+1)*c.participantsInGroups],
		})
	}

	return groups
}
