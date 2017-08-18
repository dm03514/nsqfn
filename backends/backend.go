package backends

import "github.com/dm03514/nsqfn/windower"

type Backend interface {
	Write(*windower.WindowMessages) int
}


type BackendRoot struct {
	RootDir string
}

func (br *BackendRoot) BaseDir() string {
	return br.RootDir
}


type PathTemplate struct {
	Template string
}

func (pt *PathTemplate) Path(groupBy windower.GroupByKey) string {
	// interpolate template with the group by key
}