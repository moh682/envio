package domain_test

import (
	"testing"

	"github.com/moh682/envio/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewPaymentMethod(t *testing.T) {
	tests := []struct {
		testName  string
		arg       int
		intResult int
		strResult string
		err       error
	}{
		{testName: `"UNKNOWN" should have an equal in text`, arg: domain.UNKNOWN.Index(), intResult: 0, strResult: "UNKNOWN", err: nil},
		{testName: "'CASH' should have an equal in text", arg: domain.CASH.Index(), intResult: 1, strResult: "CASH", err: nil},
		{testName: "'MOBILE_PAY' should have an equal in text", arg: domain.MOBILEPAY.Index(), intResult: 2, strResult: "MOBILEPAY", err: nil},
		{testName: "'BANK_TRANSFER' should have an equal in text", arg: domain.BANK_TRANSFER.Index(), intResult: 3, strResult: "BANK_TRANSFER", err: nil},
		{testName: "should break when an negative index is provided", arg: -1, intResult: 0, strResult: "UNKNOWN", err: domain.ErrInvalidPaymentMethod},
		{testName: "should break when an an overflow index is provided", arg: 5, intResult: 0, strResult: "UNKNOWN", err: domain.ErrInvalidPaymentMethod},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			method, err := domain.NewPaymentMethod(test.arg)
			if err != nil {
				assert.ErrorIs(t, err, test.err)
				assert.Equal(t, domain.UNKNOWN, method)
				return
			}

			assert.Equal(t, test.intResult, method.Index())
			assert.Equal(t, test.strResult, method.String())

		})

	}
}
