package backends

import "github.com/dm03514/nsqfn/windower"

type Backend interface {
	Write(*windower.WindowMessages) int
}
