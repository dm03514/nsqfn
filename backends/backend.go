package backends

import "github.com/dm03514/nsqfn/windower"

type Backend interface {
	Write(*windower.WindowMessages) int
}


type BackendRoot string

func (br *BackendRoot) BaseDir() string {
	return ""
}


type PathTemplate string

func (pt *PathTemplate) Path(groupBy windower.GroupByKey) string {
	// interpolate template with the
}