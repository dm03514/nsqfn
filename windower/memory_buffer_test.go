package windower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasBufferReachedCapacity(t *testing.T) {
	mb := &MemoryBuffer{
		buffered: map[string][]*WindowMessage{
			"test": {&WindowMessage{}},
		},
		maxMessagesPerKey: 1,
	}
	assert.True(t, mb.HasBufferReachedCapacity("test"), "Buffer has reached capacity")
}

func TestBufferHasNotReachedCapacity(t *testing.T) {
	t.Fail()
}

func TestAllBuffersHaveNotReachedCapacity(t *testing.T) {
	t.Fail()
}

func TestAllBuffersHaveReachedCapacity(t *testing.T) {
	t.Fail()
}

func TestFlushBufferWritesPersistence(t *testing.T) {
	t.Fail()
}
