package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxNegativeAmount(t *testing.T) {
	tax, err := CalculateTax(-1000.0)
	assert.Error(t, err, "Error should not be nil")
	assert.Equal(t, 0.0, tax, "Tax for negative amount should be 0")
}

func TestCalculateTaxLessThan1000(t *testing.T) {
	tax, err := CalculateTax(500.0)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 5.0, tax, "Tax for 500 should be 5")
}

func TestCalculateTaxBetween1000And20000(t *testing.T) {
	tax, err := CalculateTax(15000.0)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 10.0, tax, "Tax for 15000 should be 10")
}

func TestCalculateTaxGreaterThan20000(t *testing.T) {
	tax, err := CalculateTax(25000.0)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 20.0, tax, "Tax for 25000 should be 20")
}
