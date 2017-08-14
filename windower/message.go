package windower

import "github.com/nsqio/go-nsq"

type WindowMessage struct {
	*nsq.Message
}

// gets the keys routing key
func NewWindowMessage(m *nsq.Message) *WindowMessage {
	return &WindowMessage{Message: m}
}

func (wm *WindowMessage) GroupByKey() string {
	return ""
}
