package find

import (
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/sirupsen/logrus"
	"github.com/sundowndev/covermyass/v2/lib/filter"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Finder interface {
	Run() error
	Results() []FileInfo
}

type finder struct {
	fs      fs.FS
	filter  filter.Filter
	paths   []string
	results []FileInfo
}

func New(fsys fs.FS, filterEngine filter.Filter, paths []string) Finder {
	return &finder{
		fs:      fsys,
		filter:  filterEngine,
		paths:   paths,
		results: make([]FileInfo, 0),
	}
}

func (f *finder) Run() error {
	// Voluntary reset the results slice
	f.results = make([]FileInfo, 0)

	for _, pattern := range f.paths {
		if len(pattern) == 0 {
			logrus.Warn("pattern skipped because it has lengh of 0")
			continue
		}

		var formattedPattern string
		if strings.Split(pattern, "")[0] == string(os.PathSeparator) {
			formattedPattern = strings.Join(strings.Split(pattern, "")[1:], "")
		}

		// TODO(sundowndev): run this in a goroutine?
		err := doublestar.GlobWalk(f.fs, filepath.ToSlash(formattedPattern), func(path string, d fs.DirEntry) error {
			info, err := d.Info()
			if err != nil {
				return err
			}
			f.results = append(f.results, &fileInfo{info, fmt.Sprintf("%s%s", string(os.PathSeparator), filepath.FromSlash(path))})
			return nil
		})
		if err != nil {
			logrus.WithField("pattern", filepath.ToSlash(formattedPattern)).Error(err)
		}
	}
	return nil
}

func (f *finder) Results() []FileInfo {
	// Remove duplicates
	resultsMap := make(map[string]FileInfo, 0)
	for _, res := range f.results {
		resultsMap[res.Path()] = res
	}
	resultsSlice := make([]FileInfo, 0)
	for _, file := range resultsMap {
		resultsSlice = append(resultsSlice, file)
	}

	return resultsSlice
}
