package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceToMap(t *testing.T) {
	type Contact struct {
		Name string
		Age  int
	}
	input := []Contact{
		{
			Name: "John",
			Age:  25,
		},
		{
			Name: "Susan",
			Age:  25,
		},
		{
			Name: "Mary",
			Age:  31,
		},
	}
	got := SliceToMap(input, func(v Contact) int {
		return v.Age
	})
	t.Log(got)
	expected := map[int]interface{}{
		25: []Contact{
			{
				Name: "John",
				Age:  25,
			},
			{
				Name: "Susan",
				Age:  25,
			},
		},
		31: []Contact{
			{
				Name: "Mary",
				Age:  31,
			},
		},
	}
	require.Len(t, got, len(expected))
	for k, v := range expected {
		val, ok := got[k]
		assert.True(t, ok)
		assert.ElementsMatch(t, v, val)
	}
}
