package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sundowndev/covermyass/v2/build"
	"github.com/sundowndev/covermyass/v2/cmd"
	"log"
	"runtime"

	"github.com/sundowndev/covermyass/v2/logs"
)

func main() {
	logs.Init()
	logrus.WithFields(logrus.Fields{
		"is_release": fmt.Sprintf("%t", build.IsRelease()),
		"version":    build.Name(),
		"go_version": runtime.Version(),
	}).Debug("Build info")

	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
