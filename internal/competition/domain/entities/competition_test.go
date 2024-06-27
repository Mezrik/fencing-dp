package entities

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewCompetition(t *testing.T) {
	weapon, _ := NewWeapon("Weapon 1")
	category, _ := NewCompetitionCategory("Category 1")
	competition, _ := NewCompetition("Competition 1", "Org 1", "Fed 1", National, *category, Male, *weapon, time.Now())

	if competition.Name() != "Competition 1" {
		t.Error("Expected name to be 'Competition 1', got")
	}

	if competition.ID == (uuid.UUID{}) {
		t.Error("Expected id to be not empty")
	}
}
