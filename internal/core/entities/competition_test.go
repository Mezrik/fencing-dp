package entities

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewCompetition(t *testing.T) {
	competition := NewCompetition("Competition 1")

	if competition.Name != "Competition 1" {
		t.Error("Expected name to be 'Competition 1', got", competition.Name)
	}

	if competition.ID == (uuid.UUID{}) {
		t.Error("Expected id to be not empty")
	}
}
