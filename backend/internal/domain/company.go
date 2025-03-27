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
	ID        ID        `json:"id"`
	Name      string    `json:"name"`
	Cvr       int       `json:"cvr"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Company) Validate() error {

	// Validate ID
	_, err := uuid.Parse(c.ID.String())
	if err != nil {
		return err
	}

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
