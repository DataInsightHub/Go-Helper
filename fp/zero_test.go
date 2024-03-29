package fp_test

import (
	"testing"

	"github.com/DataInsightsHub/Go-Helper/fp"
	"github.com/stretchr/testify/assert"
)

func TestIsZeroValue(t *testing.T) {
	t.Run("Should return true if value is zero", func(t *testing.T) {
		var value string

		result := fp.IsZeroValue(value)

		assert.True(t, result)
	})

	t.Run("Should return false if value is not zero", func(t *testing.T) {
		value := "test"

		result := fp.IsZeroValue(value)

		assert.False(t, result)
	})

	t.Run("Should return true if value is nil", func(t *testing.T) {
		var value *string

		result := fp.IsZeroValue(value)

		assert.True(t, result)
	})

	t.Run("Should return true if is zero struct", func(t *testing.T) {
		type test struct{}

		var value test

		result := fp.IsZeroValue(value)

		assert.True(t, result)
	})

	t.Run("Should return false if is not zero struct", func(t *testing.T) {

		type nested struct {
			m map[string]string
		}

		type test struct {
			nested nested
		}

		value := test{
			nested: nested{
				m: map[string]string{
					"test": "test",
				},
			}}

		result := fp.IsZeroValue(value)

		assert.False(t, result)
	})

	t.Run("Should return true on nil error", func(t *testing.T) {
		var value error

		result := fp.IsZeroValue(value)

		assert.True(t, result)
	})
}

func TestZeroValue(t *testing.T) {

	t.Run("Should return zero value of type", func(t *testing.T) {

		result := fp.ZeroValueOf[string]()

		assert.Equal(t, "", result)
	})
}
