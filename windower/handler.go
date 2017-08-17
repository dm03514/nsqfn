package windower

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type WindowHandler struct {
	// keys
	// key types
	// flush interval
	// backend
	messages chan *WindowMessage
	// backendRoot BackendRoot
	// pathTemplate PathTemplate
}

func (wh *WindowHandler) HandleMessage(m *nsq.Message) error {
	// get the key for this message
	// get time period
	// build groupby opts from top level and pass into WindowHandler
	wm := NewWindowMessage(m)

	// validate that the message has the expected keys or fail here
	// VALIDATE THE WINDOW MESSAGE BEFORE GOING ON
	wh.messages <- wm

	fmt.Println(wm)
	return nil
}
