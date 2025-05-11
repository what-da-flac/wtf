package identifiers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUIDv4(t *testing.T) {
	const iterationCount = 1000
	id := NewIdentifier()
	keys := make(map[string]struct{})
	for i := 0; i < iterationCount; i++ {
		uid := id.UUIDv4()
		keys[uid] = struct{}{}
	}
	assert.Len(t, keys, iterationCount)
}
