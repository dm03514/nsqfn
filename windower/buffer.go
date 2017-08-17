package windower

import (
	"context"
	"fmt"
)

type MemoryBuffer struct {
	buffered    map[*GroupByKey][]IWindowMessage
	messages    chan *WindowMessage
	persistence chan *WindowMessages
	keyTemplate *GroupByKey

	ctx context.Context

	maxMessagesPerKey   int
	maxBufferedMessages int
	numBufferedMessages int
}

func (mb *MemoryBuffer) Init() {
	mb.buffered = make(map[*GroupByKey][]IWindowMessage)
}

// add timeout to flush all
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

func (mb *MemoryBuffer) FlushBuffer(key *GroupByKey) {
	windowMessages := &WindowMessages{
		Messages:   mb.buffered[key],
		GroupByKey: key,
	}
	mb.persistence <- windowMessages
	mb.buffered[key] = nil
	delete(mb.buffered, key)
}

func (mb *MemoryBuffer) HaveAllBuffersReachedCapacity() bool {
	return mb.numBufferedMessages >= mb.maxBufferedMessages
}

func (mb *MemoryBuffer) HasBufferReachedCapacity(key *GroupByKey) bool {
	return len(mb.buffered[key]) >= mb.maxMessagesPerKey
}

func (mb *MemoryBuffer) Push(wm IWindowMessage) {
	key := wm.GroupByKey(mb.keyTemplate)
	mb.buffered[key] = append(mb.buffered[key], wm)

	mb.numBufferedMessages++
	if mb.HaveAllBuffersReachedCapacity() {
		mb.FlushAll()
	}

	if mb.HasBufferReachedCapacity(key) {
		mb.FlushBuffer(key)
	}
}
