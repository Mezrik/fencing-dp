package entities

import (
	"github.com/Mezrik/fencing-dp/internal/common"
)

type CompetitionCategory struct {
	common.Entity
	Name string
}
