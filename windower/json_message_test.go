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
