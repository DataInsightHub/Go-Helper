package fp_test

import (
	"testing"

	"github.com/DataInsightsHub/Go-Helper/fp"
	"github.com/stretchr/testify/assert"
)

func TestDereferenceValue(t *testing.T) {
	t.Run("Should return the value of a pointer", func(t *testing.T) {
		value := "test"
		ptr := &value

		result := fp.DereferencePointer(ptr)

		assert.Equal(t, value, result)
	})

	t.Run("Should return null value of ptr", func(t *testing.T) {
		var ptr *string

		result := fp.DereferencePointer(ptr)

		assert.Equal(t, result, "")
	})
}

func TestReferenceValue(t *testing.T) {
	t.Run("Should return a pointer to the value", func(t *testing.T) {
		value := "test"

		result := fp.ReferenceValue(value)

		assert.Equal(t, value, *result)
	})
}
