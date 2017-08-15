package backends

import (
	"context"
	"github.com/dm03514/nsqfn/windower"
	"os"
	"path/filepath"
)

type FileSystem struct {
	baseDir     string
	persistence chan *windower.WindowMessages
	fin         chan *windower.WindowMessages

	ctx context.Context
}

// For right now we only open files when its time to flush and
// then write to it and close it, so we don't need to keep track of all open
// files and their sizes

// crash recovery

// closing files when they are complete
func (fs *FileSystem) Write(wms *windower.WindowMessages) int {
	var f *os.File

	// if file does not exist, initialize it
	if _, err := os.Stat(fs.Path(wms)); os.IsNotExist(err) {
		f = fs.file(fs.Path(wms))
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

func (fs *FileSystem) Path(wms *windower.WindowMessages) string {
	return filepath.Join(fs.baseDir, wms.Path())
}
