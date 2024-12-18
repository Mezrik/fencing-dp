package entities

import (
	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type ShiftCriteria string

const (
	Nation ShiftCriteria = "nation"
	Club   ShiftCriteria = "club"
	Rating ShiftCriteria = "rating"
)

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
		competitionID:             competitionID,
		number:                    number,
		participantsStartingCount: participantsStartingCount,
		numberOfGroups:            recommendedSize.groupsCount,
		participantsInGroups:      slices.Max(recommendedSize.groupSizes),
		shiftCriteria:             shiftCriteria,
		numberOfAdvancers:         participantsStartingCount * advancingPercent / 100,
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
