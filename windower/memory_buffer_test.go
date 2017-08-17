package windower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasBufferReachedCapacity(t *testing.T) {
	key := &GroupByKey{}
	mb := &MemoryBuffer{
		buffered: map[*GroupByKey][]IWindowMessage{
			key: {&WindowMessage{}},
		},
		maxMessagesPerKey: 1,
	}
	assert.True(t,
		mb.HasBufferReachedCapacity(key),
		"Buffer should have reached capacity",
	)
}

func TestBufferHasNotReachedCapacity(t *testing.T) {
	key := &GroupByKey{}
	mb := &MemoryBuffer{
		buffered: map[*GroupByKey][]IWindowMessage{
			key: {&WindowMessage{}},
		},
		maxMessagesPerKey: 2,
	}
	assert.False(t,
		mb.HasBufferReachedCapacity(key),
		"Buffer shouldn't have reached capacity",
	)
}

func TestBufferHaveAllBuffersReachedCapacity(t *testing.T) {
	key := &GroupByKey{}

	mb := &MemoryBuffer{
		buffered: map[*GroupByKey][]IWindowMessage{
			key: {&WindowMessage{}},
		},
		maxBufferedMessages: 2,
	}
	assert.False(t,
		mb.HaveAllBuffersReachedCapacity(),
		"Buffers shouldn't have reached capacity",
	)
}

type StubWindowMessage struct {
	keyTemplate *GroupByKey
	key         *GroupByKey
}

func (swm *StubWindowMessage) GroupByKey(keyTemplate *GroupByKey) *GroupByKey {
	swm.keyTemplate = keyTemplate
	return swm.key
}

func TestBufferPushNoBuffersFilled(t *testing.T) {
	stubKey := &GroupByKey{
		"stub": nil,
	}
	keyTemplate := &GroupByKey{}
	swm := &StubWindowMessage{
		keyTemplate: keyTemplate,
		key:         stubKey,
	}
	mb := &MemoryBuffer{
		maxBufferedMessages: 2,
		maxMessagesPerKey:   2,
		keyTemplate:         keyTemplate,
	}
	mb.Init()
	mb.Push(swm)

	assert.Equal(t,
		len(mb.buffered),
		1,
		"should have a single buffered key",
	)

	assert.Equal(t,
		len(mb.buffered[stubKey]),
		1,
		"should have a single buffered message",
	)
}

func TestMemoryBuffer_FlushBufferDeletesKey(t *testing.T) {
	stubKey := &GroupByKey{
		"stub": nil,
	}

	persistence := make(chan *WindowMessages)
	wm := &WindowMessage{}

	mb := &MemoryBuffer{
		maxBufferedMessages: 2,
		maxMessagesPerKey:   2,
		persistence:         persistence,
		buffered: map[*GroupByKey][]IWindowMessage{
			stubKey: {wm},
		},
	}

	go mb.FlushBuffer(stubKey)

	<-persistence

	assert.Equal(t,
		len(mb.buffered),
		0,
		"should have no buffered keys",
	)

}
