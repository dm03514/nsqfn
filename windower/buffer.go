package windower

import (
	"context"
	"fmt"
)

type WindowMessages []*WindowMessage

type MemoryBuffer struct {
	buffered map[string]WindowMessages
	messages chan *WindowMessage
	ctx      context.Context

	maxMessagesPerKey   int
	maxTotalEvents      int
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

func (mb *MemoryBuffer) FlushAll()              {}
func (mb *MemoryBuffer) FlushBuffer(key string) {}
func (mb *MemoryBuffer) HaveAllBuffersReachedCapacity() bool {

	return true
}

func (mb *MemoryBuffer) HasBufferReachedCapacity(key string) bool {

	return true
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
