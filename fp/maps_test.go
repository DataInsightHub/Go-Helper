package fp_test

import (
	"fmt"
	"testing"

	"github.com/DataInsightHub/Go-Helper/fp"
	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	slice := fp.Values(m)

	assert.Equal(t, 2, len(slice))
	assert.Equal(t, true, fp.Contains(slice, 1))
	assert.Equal(t, true, fp.Contains(slice, 2))
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	slice := fp.Keys(m)

	assert.Equal(t, 2, len(slice))
	assert.Equal(t, true, fp.Contains(slice, "a"))
	assert.Equal(t, true, fp.Contains(slice, "b"))
}

func TestMapMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	fn := func(value int) string { return fmt.Sprintf("%d", value) }
	result := fp.MapMap(m, fn)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "1", result["a"])
	assert.Equal(t, "2", result["b"])
}
