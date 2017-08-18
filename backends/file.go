package backends

import (
	"context"
	"fmt"
	"github.com/dm03514/nsqfn/windower"
	"os"
	"path/filepath"
)

type FileSystem struct {
	BackendRoot  BackendRoot
	PathTemplate PathTemplate
	Persistence  chan *windower.WindowMessages
	Fin          chan windower.IWindowMessage

	ctx context.Context
}

// For right now we only open files when its time to flush and
// then write to it and close it, so we don't need to keep track of all open
// files and their sizes

func (fs *FileSystem) Loop() {
	select {
	case windowMessages := <-fs.Persistence:
		fmt.Println(windowMessages)
		fs.Write(windowMessages)
	case <-fs.ctx.Done():
		fmt.Println("done")
	}
}

// closing files when they are complete
func (fs *FileSystem) Write(wms *windower.WindowMessages) {
	var f *os.File

	// if file does not exist, initialize it
	f = fs.file(fs.FullPath(wms))
	defer f.Close()

	// write the messages to it
	f.Write(wms.Bytes())

	for _, m := range wms.Messages {
		fs.Fin <- m
	}
}

func (fs *FileSystem) file(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func (fs *FileSystem) FullPath(wms *windower.WindowMessages) string {
	fullPath := filepath.Join(
		fs.BackendRoot.BaseDir(),
		fs.PathTemplate.Path(wms.GroupByKey),
		wms.FileName(),
	)
	fmt.Println(fullPath)
	return fullPath
}
