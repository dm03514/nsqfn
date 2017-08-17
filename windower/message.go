package windower

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/satori/go.uuid"
)

type IWindowMessage interface {
	// crazy
	// CHANGE TO Key() we are getting the message key, or BufferKey() or something
	GroupByKey(keyTemplate *GroupByKey) *GroupByKey
}

// map of keys to group on
type GroupByKey map[string]interface{}

type WindowMessage struct {
	*nsq.Message
}

// JSON window messages, need an interface
// WindowMessage need an interface JSONWindowMessage
type WindowMessages struct {
	Messages   []IWindowMessage
	fileName   string
	GroupByKey *GroupByKey
}

func (wm *WindowMessages) FileName() string {
	if wm.fileName != "" {
		return wm.fileName
	}
	u := uuid.NewV4()
	return fmt.Sprintf("%s.json", u.String())
}

// serializes the window messages and returns byte array
// we need to not copy this to memory.   Each of the underlying
// nsq message Body should already be a byte array, so we just
// need to stream them out and join them somehow
func (wm *WindowMessages) Bytes() []byte {
	return []byte{}
}

// gets the keys routing key
func NewWindowMessage(m *nsq.Message) *WindowMessage {
	return &WindowMessage{Message: m}
}

func (wm *WindowMessage) GroupByKey(keyTemplate *GroupByKey) *GroupByKey {
	return &GroupByKey{}
}
