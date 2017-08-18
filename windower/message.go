package windower

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/satori/go.uuid"
)

type IWindowMessage interface {
	Body() []byte
	// crazy
	// CHANGE TO Key() we are getting the message key, or BufferKey() or something
	GroupByKey(keyTemplate GroupByKey) *GroupByKey
}

// map of keys to group on
type GroupByKey map[string]interface{}

type WindowMessage struct {
	*nsq.Message
}

func (wm *WindowMessage) Body() []byte {
	return wm.Message.Body
}

// TODO do maps automatically pass by reference??
// Can't copy this here, some payloads are huge
func (wm *WindowMessage) JSON() map[string]interface{} {
	var data map[string]interface{}

	if err := json.Unmarshal(wm.Body(), &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	return data
}

// Given a template, will populate that template with
// the WindowMessages values and return a pointer to the template
// TODO Need to have additional checks, and return an error, if the
// message doesn't completely fulfill the template it is invalid
func (wm *WindowMessage) GroupByKey(keyTemplate GroupByKey) *GroupByKey {
	j := wm.JSON()

	for k, v := range keyTemplate {
		// TODO what effects does this type conversion have? utf-8? what about float?
		keyTemplate[k] = j[k].(string)
		fmt.Println(k)
		fmt.Println(v)
	}
	return &keyTemplate
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
