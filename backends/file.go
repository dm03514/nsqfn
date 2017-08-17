package backends

import (
	"context"
	"github.com/dm03514/nsqfn/windower"
	"os"
	"path/filepath"
)

type FileSystem struct {
	backendRoot     BackendRoot
	pathTemplate PathTemplate
	persistence chan *windower.WindowMessages
	fin         chan *windower.WindowMessage

	ctx context.Context

}

// For right now we only open files when its time to flush and
// then write to it and close it, so we don't need to keep track of all open
// files and their sizes

// crash recovery

// closing files when they are complete
func (fs *FileSystem) Write(wms *windower.WindowMessages) {
	var f *os.File

	// if file does not exist, initialize it
	if _, err := os.Stat(fs.FullPath(wms)); os.IsNotExist(err) {
		f = fs.file(fs.FullPath(wms))
	}
	defer f.Close()

	// write the messages to it
	f.Write(wms.Bytes())
}

func (fs *FileSystem) file(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func (fs *FileSystem) FullPath(wms *windower.WindowMessages) string {
	return filepath.Join(
		fs.backendRoot.BaseDir(),
		fs.pathTemplate.Path(wms.GroupByKey),
		wms.FileName(),
	)
}
