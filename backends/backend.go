package backends

import (
	"github.com/dm03514/nsqfn/windower"
	"text/template"
	"bytes"
)

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

func (pt *PathTemplate) Path(groupBy *windower.GroupByKey) string {
	// interpolate template with the group by key
	var result bytes.Buffer
	t, err := template.New("path").Parse(pt.Template)
	if err != nil {
		panic(err)
	}
	t.Execute(&result, groupBy)
	if err != nil {
		panic(err)
	}
	return result.String()
}