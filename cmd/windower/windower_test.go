package main

import (
	"testing"
	"github.com/nsqio/go-nsq"
	"github.com/dm03514/nsqfn/backends"
	"runtime"
)

// Test the full integration of the whole windower life cycle starting at the
// WindowMessage channel
// TODO create interface for nsq.Message, so we can test the whole
func TestWindowerFullMessageLifeCycle(t *testing.T) {
	messagesChan := make(chan *WindowMessage)
	persistenceChan := make(chan *WindowMessages)
	finChan := make(chan *WindowMessage)

	handler := &WindowHandler{
		messages: messagesChan,
	}

	buffer := &MemoryBuffer{
		messages: messagesChan,
		persistence: persistenceChan,
		maxBufferedMessages: 1,
	}
	buffer.Init()

	fs := &backends.FileSystem{
		BackendRoot: &backends.BackendRoot{
			RootDir: "/tmp",
		},
		PathTemplate: &backends.PathTemplate{
			Template: "nsqfn/windower/{{ .user_id }}/",
		},
		Persistence: persistenceChan,
		Fin: finChan,
	}

	m := &nsq.Message{
		Body: []byte("body1"),
	}

	go fs.Loop()
	go buffer.Loop()
	go handler.HandleMessage(m)

	t.Fatalf("%s", <-finChan)
}
