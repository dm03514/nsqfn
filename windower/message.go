package windower

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/satori/go.uuid"
	"bytes"
)

type IWindowMessage interface {
	Body() []byte
	// crazy
	// CHANGE TO Key() we are getting the message key, or BufferKey() or something
	GroupByKey(keyTemplate *GroupByKey) *GroupByKey
}

// map of keys to group on
type GroupByKey map[string]interface{}

type WindowMessage struct {
	*nsq.Message
}

func (wm *WindowMessage) Body() []byte {
	return wm.Message.Body
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
// nsq message Body should already be a byte slice, so we just
// need to stream them out and join them somehow
func (wm *WindowMessages) Bytes() []byte {
	bs := [][]byte{}
	for _, m := range wm.Messages {
		bs = append(bs, m.Body())
	}
	return bytes.Join(bs, []byte("\n"))
}

// gets the keys routing key
func NewWindowMessage(m *nsq.Message) *WindowMessage {
	return &WindowMessage{Message: m}
}

func (wm *WindowMessage) GroupByKey(keyTemplate *GroupByKey) *GroupByKey {
	return &GroupByKey{}
}
