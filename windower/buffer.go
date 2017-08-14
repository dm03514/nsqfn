package windower

import (
	"context"
	"fmt"
)

type WindowMessages struct {
	Messages   []*WindowMessage
	GroupByKey string
}

type MemoryBuffer struct {
	buffered    map[string][]*WindowMessage
	messages    chan *WindowMessage
	persistence chan *WindowMessages

	ctx context.Context

	maxMessagesPerKey   int
	maxBufferedMessages int
	numBufferedMessages int
}

func (mb *MemoryBuffer) Loop() {
	select {
	case message := <-mb.messages:
		fmt.Println(message)
		mb.Push(message)
	case <-mb.ctx.Done():
		fmt.Println("done")
	}
}

func (mb *MemoryBuffer) FlushAll() {
	for key := range mb.buffered {
		mb.FlushBuffer(key)
	}
}

func (mb *MemoryBuffer) FlushBuffer(key string) {
	windowMessages := &WindowMessages{
		Messages:   mb.buffered[key],
		GroupByKey: key,
	}
	mb.persistence <- windowMessages
	// will this reinitialize ????? and allow us to append to the slice?
	mb.buffered[key] = nil
	delete(mb.buffered, key)
}

func (mb *MemoryBuffer) HaveAllBuffersReachedCapacity() bool {
	return mb.numBufferedMessages >= mb.maxBufferedMessages
}

func (mb *MemoryBuffer) HasBufferReachedCapacity(key string) bool {
	return len(mb.buffered[key]) >= mb.maxMessagesPerKey
}

func (mb *MemoryBuffer) Push(m *WindowMessage) {
	key := m.GroupByKey()
	mb.buffered[key] = append(mb.buffered[key], m)

	mb.numBufferedMessages++
	if mb.HaveAllBuffersReachedCapacity() {
		mb.FlushAll()
	}

	if mb.HasBufferReachedCapacity(key) {
		mb.FlushBuffer(key)
	}
}
