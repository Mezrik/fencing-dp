package entities

import (
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type DeploymentType string

const (
	Deployment DeploymentType = "deployment"
)

type CompetitionParameters struct {
	common.Entity
	expectedParticipants       int
	deploymentType             DeploymentType
	groupHits                  int
	eliminationHits            int
	qualificationBasedOnRounds int
}

func (g DeploymentType) TSName() string {
	switch g {
	case Deployment:
		return "deployment"
	default:
		return "???"
	}
}

func NewCompetitionParameters(expectedParticipants int, deploymentType DeploymentType, groupHits int, eliminationHits int, qualificationBasedOnRounds int) *CompetitionParameters {
	return &CompetitionParameters{
		Entity:                     common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		expectedParticipants:       expectedParticipants,
		deploymentType:             deploymentType,
		groupHits:                  groupHits,
		eliminationHits:            eliminationHits,
		qualificationBasedOnRounds: qualificationBasedOnRounds,
	}
}

func UnmarshalCompetitionParameters(id uuid.UUID, expectedParticipants int, deploymentType DeploymentType, groupHits int, eliminationHits int, qualificationBasedOnRounds int, createdAt time.Time, updatedAt *time.Time) *CompetitionParameters {
	return &CompetitionParameters{
		Entity:                     common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		expectedParticipants:       expectedParticipants,
		deploymentType:             deploymentType,
		groupHits:                  groupHits,
		eliminationHits:            eliminationHits,
		qualificationBasedOnRounds: qualificationBasedOnRounds,
	}
}

func (c *CompetitionParameters) ExpectedParticipants() int {
	return c.expectedParticipants
}

func (c *CompetitionParameters) DeploymentType() DeploymentType {
	return c.deploymentType
}

func (c *CompetitionParameters) GroupHits() int {
	return c.groupHits
}

func (c *CompetitionParameters) EliminationHits() int {
	return c.eliminationHits
}

func (c *CompetitionParameters) QualificationBasedOnRounds() int {
	return c.qualificationBasedOnRounds
}
