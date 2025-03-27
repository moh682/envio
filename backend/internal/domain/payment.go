package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID     ID        `json:"id"`
	Amount float32   `json:"amount"`
	PaidAt time.Time `json:"paid_at"`
	Method string    `json:"method"`
}

func (p *Payment) Validate() error {
	_, err := uuid.Parse(p.ID.String())
	if err != nil {
		return err
	}

	return nil
}
