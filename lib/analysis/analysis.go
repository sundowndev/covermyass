package analysis

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Summary struct {
	TotalFiles   int
	TotalRWFiles int
	TotalROFiles int
}

type Result struct {
	Path string
	Size int64
	Mode os.FileMode
}

type Analysis struct {
	Date     time.Time
	summary  Summary
	patterns []string
	results  []Result
}

func New() *Analysis {
	return &Analysis{
		Date:     time.Now(),
		summary:  Summary{},
		patterns: []string{},
		results:  []Result{},
	}
}

func (a *Analysis) AddPatterns(patterns ...string) {
	a.patterns = append(a.patterns, patterns...)
}

func (a *Analysis) Patterns() []string {
	return a.patterns
}

func (a *Analysis) AddResult(result Result) {
	a.results = append(a.results, result)
	a.summary.TotalFiles += 1
}

func (a *Analysis) Results() []Result {
	return a.results
}

func (a *Analysis) Write(w io.Writer) {
	if len(a.results) > 0 {
		_, _ = fmt.Fprintf(w, "Found the following files\n")

		for _, res := range a.results {
			_, _ = fmt.Fprintf(w, "%s (%s)\n", res.Path, byteCountSI(res.Size))
		}
		_, _ = fmt.Fprintf(w, "\n")
	}

	_, _ = fmt.Fprintf(
		w,
		"Summary\nFound %d files (%d RW, %d RO) in %s\n",
		a.summary.TotalFiles,
		a.summary.TotalRWFiles,
		a.summary.TotalROFiles,
		time.Since(a.Date).Round(time.Millisecond).String(),
	)
}
