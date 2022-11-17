package analysis

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sundowndev/covermyass/v2/lib/check"
	"github.com/sundowndev/covermyass/v2/lib/filter"
	"github.com/sundowndev/covermyass/v2/lib/find"
	"github.com/sundowndev/covermyass/v2/lib/output"
	"os"
	"runtime"
	"sync"
)

type Analyzer struct {
	filter filter.Filter
}

func NewAnalyzer(filterEngine filter.Filter) *Analyzer {
	return &Analyzer{filterEngine}
}

func (a *Analyzer) Analyze() (*Analysis, error) {
	analysis := NewAnalysis()

	output.Printf("Loaded known log files for %s\n", runtime.GOOS)
	output.Printf("Scanning file system...\n\n")

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for _, c := range check.GetAllChecks() {
		wg.Add(1)
		go func(c check.Check) {
			finder := find.New(os.DirFS(""), a.filter)
			results, err := finder.Run(context.TODO(), c.Paths())
			if err != nil {
				logrus.Error(err)
				return
			}

			m.Lock()
			defer m.Unlock()
			for _, info := range results {
				analysis.AddResult(Result{
					Check:    c,
					Path:     info.Path(),
					Size:     info.Size(),
					Mode:     info.Mode(),
					ReadOnly: info.ReadOnly(),
				})
			}

			wg.Done()
		}(c)
	}

	wg.Wait()

	return analysis, nil
}
