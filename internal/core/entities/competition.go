package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Competition struct {
	Entity
	Name     string
	Category CompetitionCategory
	Weapon   Weapon
}

func (c *Competition) validate() error {
	if c.Name == "" {
		return errors.New("name must not be empty")
	}

	return nil
}

func NewCompetition(name string) *Competition {
	return &Competition{
		Entity: Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:   name,
	}
}

func (c *Competition) UpdateName(name string) error {
	c.Name = name
	c.UpdatedAt = time.Now()

	return c.validate()
}

func (c *Competition) UpdateCategory(category CompetitionCategory) error {
	c.Category = category
	c.UpdatedAt = time.Now()

	return c.validate()
}
