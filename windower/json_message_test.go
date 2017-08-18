package windower

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWindowMessages_FileName(t *testing.T) {
	wm := &WindowMessages{}
	fileName := wm.FileName()
	assert.True(t,
		strings.HasSuffix(fileName, ".json"),
		fmt.Sprintf("%s should have suffix .json", fileName),
	)
}


type StubMessage struct {
	body []byte
}
func (sm *StubMessage) GroupByKey(key *GroupByKey) *GroupByKey {
	return &GroupByKey{}
}
func (sm *StubMessage) Body() []byte {
	return sm.body
}

func TestWindowMessages_Bytes_SingleMessage(t *testing.T) {
	wms := &WindowMessages{
		Messages: []IWindowMessage{
			&StubMessage{body: []byte("body1")},
		},
	}
	bs := wms.Bytes()
	assert.Equal(t,
		string(bs),
		"body1",
	)
}

func TestWindowMessages_Bytes_MultipleMessages(t *testing.T) {
	wms := &WindowMessages{
		Messages: []IWindowMessage{
			&StubMessage{body: []byte("body1")},
			&StubMessage{body: []byte("body2")},
		},
	}
	bs := wms.Bytes()
	assert.Equal(t,
		string(bs),
		"body1\nbody2",
	)
}
