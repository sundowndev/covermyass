package shred

import (
	"crypto/rand"
	"fmt"
	"io/fs"
	"os"
	"time"
)

// A FileInfo describes a file and is returned by Stat.
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() fs.FileMode  // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() any           // underlying data source (can return nil)
}

type File interface {
	Seek(int64, int) (int64, error)
	Sync() error
	Write([]byte) (int, error)
	Close() error
}

type ShredderOptions struct {
	Zero       bool
	Iterations int
}

type Shredder struct {
	options *ShredderOptions
}

func New(opts *ShredderOptions) *Shredder {
	return &Shredder{opts}
}

func (s *Shredder) Write(pathName string) error {
	// Stat the file for the file length
	fstat, err := os.Stat(pathName)
	if err != nil {
		return fmt.Errorf("file stat failed: %w", err)
	}

	// Open the file
	file, err := os.OpenFile(pathName, os.O_WRONLY, fstat.Mode())
	if err != nil {
		return fmt.Errorf("file opening failed: %w", err)
	}
	defer file.Close()

	err = s.shred(fstat, file)
	if err != nil {
		return fmt.Errorf("shredding failed: %w", err)
	}

	if s.options.Zero {
		if err := os.Truncate(pathName, 0); err != nil {
			return fmt.Errorf("truncate failed: %w", err)
		}
	}

	return nil
}

func (s *Shredder) shred(fstat FileInfo, file File) error {
	fSize := fstat.Size()

	// Avoid shredding if the file is already empty
	if fSize == 0 {
		return nil
	}

	// Write random bytes over the file 3 times
	junkBuf := make([]byte, 1024)
	for i := 0; i < s.options.Iterations; i++ {
		_, err := file.Seek(0, 0)
		if err != nil {
			return err
		}
		for fSize = fstat.Size(); fSize > 1024; fSize -= 1024 {
			// Load a buffer with random data
			_, err = rand.Read(junkBuf)
			if err != nil {
				return err
			}
			// Write random bytes to file
			_, err = file.Write(junkBuf)
			if err != nil {
				return err
			}
		}
		_, err = rand.Read(junkBuf[:fSize])
		if err != nil {
			return err
		}
		_, err = file.Write(junkBuf[:fSize])
		if err != nil {
			return err
		}
		err = file.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}
