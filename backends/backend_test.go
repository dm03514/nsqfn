package backends

import (
	"github.com/dm03514/nsqfn/windower"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathInterpolateValidParams(t *testing.T) {
	pt := &PathTemplate{
		Template: "nsqfn/windower/{{ .user_id }}/",
	}
	k := &windower.GroupByKey{
		"user_id": 1,
		"name":    "test",
	}
	assert.Equal(t,
		pt.Path(k),
		"nsqfn/windower/1/",
	)
}
