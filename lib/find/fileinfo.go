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
	ro   bool
}

func (f *fileInfo) Path() string {
	return f.path
}

func (f *fileInfo) ReadOnly() bool {
	_, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		f.ro = true
	}
	return f.ro
}
