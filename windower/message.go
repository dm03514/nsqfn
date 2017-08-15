package windower

import "github.com/nsqio/go-nsq"

type WindowMessage struct {
	*nsq.Message
}

type WindowMessages struct {
	Messages   []*WindowMessage
	GroupByKey string
}

func (wm *WindowMessages) Path() string {
	return wm.GroupByKey
}

// serializes the window messages and returns byte array
func (wm *WindowMessages) Bytes() []byte {
	return []byte{}
}

// gets the keys routing key
func NewWindowMessage(m *nsq.Message) *WindowMessage {
	return &WindowMessage{Message: m}
}

func (wm *WindowMessage) GroupByKey() string {
	return ""
}
