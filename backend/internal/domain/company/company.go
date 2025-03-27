package domain

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidCVR         = errors.New("cvr is not valid")
	ErrNameTooShort       = errors.New("name is too short, must be at least 3 characters")
	ErrInvalidPhoneNumber = errors.New("phone number is not valid")
)

type Company struct {
	ID        uuid.UUID
	Name      string
	Cvr       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCompany(id uuid.UUID, name string, cvr int, createdAt, updatedAt time.Time) (*Company, error) {

	comp := &Company{
		ID:        id,
		Name:      name,
		Cvr:       cvr,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	err := comp.validate()
	if err != nil {
		return nil, err
	}

	return comp, nil
}

func (c *Company) validate() error {
	//  Validate CVR
	strCvr := strconv.FormatInt(int64(c.Cvr), 10)
	hasNonDigit, err := regexp.MatchString(`\D`, strCvr)
	if err != nil {
		return err
	}
	switch {
	case len(strCvr) != 8:
		return ErrInvalidCVR
	case hasNonDigit:
		return ErrInvalidCVR
	}

	// Validate Name
	switch {
	case len(c.Name) < 3:
		return ErrNameTooShort
	}

	return nil

}
