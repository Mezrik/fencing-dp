package entities

import (
	"errors"
	"time"

	"github.com/Mezrik/fencing-dp/internal/common"
	"github.com/google/uuid"
)

type GenderEnum string

const (
	Male   GenderEnum = "male"
	Female GenderEnum = "female"
	Mixed  GenderEnum = "mixed"
	Unset  GenderEnum = "unset"
)

func (g GenderEnum) TSName() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	case Mixed:
		return "mixed"
	case Unset:
		return "unset"
	default:
		return "???"
	}
}

type Competitor struct {
	common.Entity
	firstname   string
	surname     string
	gender      GenderEnum
	club        *Club
	license     *string
	license_fie *string
	birthdate   *time.Time
}

func NewCompetitor(firstname string, surname string, gender GenderEnum, club *Club, license *string, license_fie *string, birthdate *time.Time) (*Competitor, error) {
	if firstname == "" {
		return nil, errors.New("firstname must not be empty")
	}
	if surname == "" {
		return nil, errors.New("surname must not be empty")
	}

	return &Competitor{
		Entity:      common.Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: nil},
		firstname:   firstname,
		surname:     surname,
		gender:      gender,
		club:        club,
		license:     license,
		license_fie: license_fie,
		birthdate:   birthdate,
	}, nil
}

func UnmarshalCompetitor(id uuid.UUID, firstname string, surname string, gender GenderEnum, club *Club, license *string, license_fie *string, birthdate *time.Time, createdAt time.Time, updatedAt *time.Time) *Competitor {
	return &Competitor{
		Entity:      common.Entity{ID: id, CreatedAt: createdAt, UpdatedAt: updatedAt},
		firstname:   firstname,
		surname:     surname,
		gender:      gender,
		club:        club,
		license:     license,
		license_fie: license_fie,
		birthdate:   birthdate,
	}
}

func (c Competitor) FirstName() string {
	return c.firstname
}

func (c Competitor) Surname() string {
	return c.surname
}

func (c Competitor) Gender() GenderEnum {
	return c.gender
}

func (c Competitor) Club() *Club {
	return c.club
}

func (c Competitor) License() *string {
	return c.license
}

func (c Competitor) LicenseFie() *string {
	return c.license_fie
}

func (c Competitor) Birthdate() *time.Time {
	return c.birthdate
}

func (c Competitor) HasMissingInfo() bool {
	return c.license == nil || c.birthdate == nil || c.club == nil || c.gender == Unset
}

func (c *Competitor) SetFirstname(s string) {
	c.firstname = s
}

func (c *Competitor) SetSurname(s string) {
	c.surname = s
}

func (c *Competitor) SetGender(g GenderEnum) {
	c.gender = g
}

func (c *Competitor) SetClub(club *Club) {
	c.club = club
}

func (c *Competitor) SetLicense(license *string) {
	c.license = license
}

func (c *Competitor) SetLicenseFie(license_fie *string) {
	c.license_fie = license_fie
}

func (c *Competitor) SetBirthdate(birthdate *time.Time) {
	c.birthdate = birthdate
}

func (c *Competitor) SetUpdatedAt(updatedAt time.Time) {
	c.UpdatedAt = &updatedAt
}
