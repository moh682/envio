package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (u *ID) String() string {
	return uuid.UUID(*u).String()
}

func (u *ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uuid.UUID(*u).String())), nil
}

func (u *ID) UnmarshalJSON(b []byte) error {
	id, err := uuid.Parse(string(b[:]))
	if err != nil {
		return err
	}
	*u = ID(id)
	return nil
}
