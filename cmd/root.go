package cmd

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sundowndev/covermyass/v2/build"
	"github.com/sundowndev/covermyass/v2/lib/analysis"
	"github.com/sundowndev/covermyass/v2/lib/filter"
	"github.com/sundowndev/covermyass/v2/lib/shred"
	"github.com/sundowndev/covermyass/v2/output"
	"os"
	"time"
)

type RootCmdOptions struct {
	List            bool
	ExcludeReadOnly bool
	Write           bool
	Zero            bool
	Iterations      int
	//ExtraPaths []string
	FilterRules []string
}

func NewRootCmd() *cobra.Command {
	opts := &RootCmdOptions{}
	cmd := &cobra.Command{
		Use:   "covermyass",
		Short: "Post-exploitation tool for covering tracks on Linux, Darwin and Windows.",
		Long:  "Covermyass is a post-exploitation tool for pen-testers that finds then erases log files on the current machine. The tool scans the filesystem and look for known log files that can be erased. Files are overwritten multiple times with random data, in order to make it harder for even very expensive hardware probing to recover the data. Running this tool with root privileges is safe and even recommended to avoid access permission errors. This tool does not perform any network call.",
		Example: `
Overwrite log files as well as those found by path /db/*.log
covermyass --write -p /db/*.log

Overwrite log files 5 times with a final overwrite with zeros to hide shredding
covermyass --write -z -n 5
`,
		Version: build.String(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.List {
				opts.Write = false
			} else {
				output.ChangePrinter(output.NewConsolePrinter())
			}

			filterEngine := filter.NewEngine()
			err := filterEngine.AddRule(opts.FilterRules...)
			if err != nil {
				return err
			}

			analyzer := analysis.NewAnalyzer(filterEngine)
			a, err := analyzer.Analyze()
			if err != nil {
				return err
			}

			if opts.List {
				for _, result := range a.Results() {
					if opts.ExcludeReadOnly && result.ReadOnly {
						continue
					}
					fmt.Println(result.Path)
				}
				return nil
			}

			a.Write(os.Stdout)

			if opts.Write {
				output.Printf("\n")
				bar := progressbar.NewOptions64(
					-1,
					progressbar.OptionSetDescription("Shredding files..."),
					progressbar.OptionSetWriter(output.GetPrinter()),
					progressbar.OptionShowBytes(true),
					progressbar.OptionSetWidth(10),
					progressbar.OptionThrottle(65*time.Millisecond),
					progressbar.OptionShowCount(),
					progressbar.OptionOnCompletion(func() {
						_, _ = fmt.Fprint(output.GetPrinter(), "\n")
					}),
					progressbar.OptionSpinnerType(11),
					progressbar.OptionFullWidth(),
					progressbar.OptionSetRenderBlankState(true),
				)

				shredOptions := &shred.ShredderOptions{
					Zero:       opts.Zero,
					Iterations: opts.Iterations,
					Bar:        bar,
				}
				s := shred.New(shredOptions)
				for _, result := range a.Results() {
					logrus.
						WithField("path", result.Path).
						Debug("Shredding file")
					if err := s.Write(result.Path); err != nil {
						return fmt.Errorf("error writing file %s: %s", result.Path, err)
					}
				}
				_ = bar.Finish()
				output.Printf("\nSuccessfully shredded %d files %d times\n", len(a.Results()), opts.Iterations)
			}

			return nil
		},
	}

	cmd.PersistentFlags().BoolVarP(&opts.List, "list", "l", false, "Show files in a simple list format. This will prevent any write operation")
	cmd.PersistentFlags().BoolVar(&opts.Write, "write", false, "Erase found log files. This WILL shred the files!")
	cmd.PersistentFlags().BoolVar(&opts.ExcludeReadOnly, "no-read-only", false, "Exclude read-only files in the list. Must be used with --list")
	cmd.PersistentFlags().BoolVarP(&opts.Zero, "zero", "z", false, "Add a final overwrite with zeros to hide shredding")
	cmd.PersistentFlags().IntVarP(&opts.Iterations, "iterations", "n", 3, "Overwrite N times instead of the default")
	cmd.PersistentFlags().StringSliceVarP(&opts.FilterRules, "filter", "f", []string{}, "File paths to ignore (supports glob patterns)")

	return cmd
}
