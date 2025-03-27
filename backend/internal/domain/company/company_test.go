package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/moh682/envio/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

type ExpenseTester struct {
}

func TestNewCompany(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		id       uuid.UUID
		cvr      int
		err      error
	}{
		{testName: "should pass when all data is valid", name: "Car Cleaning", cvr: 12345678, id: uuid.MustParse("c6c42fdc-ab89-4066-9bfa-659c6ab7bca5"), err: nil},
		{testName: "should fail when cvr number is not 8 characters", name: "Wrap in town", id: uuid.MustParse("c6c42fdc-ab89-4066-9bfa-659c6ab7bca5"), cvr: 2918264, err: domain.ErrInvalidCVR},
		{testName: "should fail when company name is too short", name: "ab", cvr: 12345678, id: uuid.MustParse("c6c42fdc-ab89-4066-9bfa-659c6ab7bca5"), err: domain.ErrNameTooShort},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			c := &domain.Company{
				Name: test.name,
				Cvr:  test.cvr,
				ID:   test.id,
			}
			err := c.Validate()
			if test.err != nil {
				assert.ErrorIs(t, err, test.err)
				assert.NotNil(t, c)
				return
			}
			assert.True(t, c.ID.String() != "")
			assert.Equal(t, c.Name, test.name)
			assert.Equal(t, c.Cvr, test.cvr)
			assert.Nil(t, err)
		})

	}

}
