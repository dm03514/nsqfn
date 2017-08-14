package windower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasBufferReachedCapacity(t *testing.T) {
	mb := &MemoryBuffer{
		buffered: map[string]WindowMessages{
			"test": {&WindowMessage{}},
		},
		maxMessagesPerKey: 1,
	}
	assert.True(t, mb.HasBufferReachedCapacity("test"), "Buffer has reached capacity")
}
