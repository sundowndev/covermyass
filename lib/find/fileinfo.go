package find

import "io/fs"

type FileInfo interface {
	fs.FileInfo
	Path() string
}

type fileInfo struct {
	fs.FileInfo
	path string
}

func (f *fileInfo) Path() string {
	return f.path
}
