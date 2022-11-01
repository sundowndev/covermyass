package find

import (
	"io/fs"
	"os"
)

type FileInfo interface {
	fs.FileInfo
	Path() string
	ReadOnly() bool
}

type fileInfo struct {
	fs.FileInfo
	path string
}

func (f *fileInfo) Path() string {
	return f.path
}

func (f *fileInfo) ReadOnly() bool {
	_, err := os.OpenFile(f.path, os.O_RDWR, 0666)
	return err != nil
}
