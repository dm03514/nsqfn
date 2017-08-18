package windower

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
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

func (sm *StubMessage) GroupByKey(key GroupByKey) *GroupByKey {
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

func TestWindowMessage_GroupByKey_FillCopyTemplate(t *testing.T) {
	keyTemplate := GroupByKey{
		"key1": nil,
		"key2": nil,
	}
	m := &WindowMessage{
		Message: &nsq.Message{
			Body: []byte("{\"key1\":\"value1\",\"key2\":\"value2\"}"),
		},
	}
	key := m.GroupByKey(keyTemplate)
	jsonKey, err := json.Marshal(key)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t,
		string(jsonKey),
		"{\"key1\":\"value1\",\"key2\":\"value2\"}",
	)
}
