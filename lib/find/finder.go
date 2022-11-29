package find

import (
	"context"
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
	Run(context.Context, []string) ([]FileInfo, error)
}

type finder struct {
	fs     fs.FS
	filter filter.Filter
}

func New(fsys fs.FS, filterEngine filter.Filter) Finder {
	return &finder{
		fs:     fsys,
		filter: filterEngine,
	}
}

func (f *finder) Run(ctx context.Context, paths []string) ([]FileInfo, error) {
	results := make([]FileInfo, 0)

	for _, pattern := range paths {
		if len(pattern) == 0 {
			logrus.Warn("pattern skipped because it has length of 0")
			continue
		}

		if !doublestar.ValidatePathPattern(pattern) {
			return results, fmt.Errorf("pattern %s is not valid", pattern)
		}

		if f.filter.Match(pattern) {
			logrus.WithField("pattern", pattern).Debug("pattern ignored by filter")
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
			results = append(results, &fileInfo{
				FileInfo: info,
				path:     fmt.Sprintf("%s%s", string(os.PathSeparator), filepath.FromSlash(path)),
			})
			return nil
		})
		if err != nil {
			logrus.WithField("pattern", filepath.ToSlash(formattedPattern)).Error(err)
		}
	}
	return f.removeDuplicates(results), nil
}

func (f *finder) removeDuplicates(results []FileInfo) []FileInfo {
	resultsMap := make(map[string]FileInfo, 0)
	for _, res := range results {
		resultsMap[res.Path()] = res
	}
	resultsSlice := make([]FileInfo, 0)
	for _, file := range resultsMap {
		resultsSlice = append(resultsSlice, file)
	}
	return resultsSlice
}
