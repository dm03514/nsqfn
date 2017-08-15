package backends

import (
	"context"
	"github.com/dm03514/nsqfn/windower"
)

type NSQFinner struct {
	fin chan *windower.WindowMessages

	ctx context.Context
}

func (nsqf *NSQFinner) Loop() {
}

func (nsqf *NSQFinner) Pool(size int) {
}
