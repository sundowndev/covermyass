package analysis

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sundowndev/covermyass/v2/lib/filter"
	"github.com/sundowndev/covermyass/v2/lib/find"
	"github.com/sundowndev/covermyass/v2/lib/output"
	"github.com/sundowndev/covermyass/v2/lib/services"
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

	output.Printf("Loading known log files for %s\n", runtime.GOOS)
	services.Init()
	output.Printf("Scanning file system...\n\n")

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for _, service := range services.Services() {
		wg.Add(1)
		go func(svc services.Service) {
			finder := find.New(os.DirFS(""), a.filter, svc.Paths())
			if err := finder.Run(context.TODO()); err != nil {
				logrus.Error(err)
				return
			}

			m.Lock()
			defer m.Unlock()
			for _, info := range finder.Results() {
				analysis.AddResult(Result{
					Service:  svc.Name(),
					Path:     info.Path(),
					Size:     info.Size(),
					Mode:     info.Mode(),
					ReadOnly: info.ReadOnly(),
				})
			}

			wg.Done()
		}(service)
	}

	wg.Wait()

	return analysis, nil
}
