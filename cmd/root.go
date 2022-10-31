package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sundowndev/covermyass/v2/build"
	"github.com/sundowndev/covermyass/v2/lib/analysis"
	"github.com/sundowndev/covermyass/v2/lib/filter"
	"github.com/sundowndev/covermyass/v2/lib/find"
	"github.com/sundowndev/covermyass/v2/lib/output"
	"github.com/sundowndev/covermyass/v2/lib/services"
	"log"
	"os"
	"runtime"
)

type RootCmdOptions struct {
	List  bool
	Write bool
	//ExtraPaths []string
	//FilterRules []string
}

func NewRootCmd() *cobra.Command {
	opts := &RootCmdOptions{}
	cmd := &cobra.Command{
		Use:   "covermyass",
		Short: "Post-exploitation tool for covering tracks on Linux, Darwin and Windows.",
		Long:  "Covermyass is a post-exploitation tool for pen-testers that finds then erases log files on the current machine. The tool scans the filesystem and look for known log files that can be erased. Running this tool with root privileges is safe and even recommended to avoid access permission errors. This tool does not perform any network call.",
		Example: "covermyass --write -p /db/*.log\n" +
			"covermyass --list -p /db/**/*.log",
		Version: build.String(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.List {
				opts.Write = false
			} else {
				output.ChangePrinter(output.NewConsolePrinter())
			}

			a := analysis.New()

			output.Printf("Loading known log files for %s\n", runtime.GOOS)
			services.Init()
			for _, service := range services.Services() {
				patterns, ok := service.Paths()[runtime.GOOS]
				if !ok {
					continue
				}
				a.AddPatterns(patterns...)
			}

			filterEngine := filter.NewEngine()
			finder := find.New(os.DirFS(""), filterEngine, a.Patterns())

			output.Printf("Searching for log files...\n\n")
			err := finder.Run()
			if err != nil {
				log.Fatal(err)
			}

			for _, info := range finder.Results() {
				a.AddResult(analysis.Result{
					Path: info.Path(),
					Size: info.Size(),
					Mode: info.Mode(),
				})
			}

			if opts.List {
				for _, result := range a.Results() {
					fmt.Println(result.Path)
				}
				return nil
			}

			a.Write(os.Stdout)

			return nil
		},
	}

	cmd.PersistentFlags().BoolVarP(&opts.List, "list", "l", false, "Show files in a simple list format. This will prevent any write operation.")
	cmd.PersistentFlags().BoolVar(&opts.Write, "write", false, "Erase found log files. This WILL truncate the files!")

	return cmd
}
