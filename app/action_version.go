package app

import (
	"fmt"
	"os"
	"runtime"

	"orchestrator/metadata"

	"gopkg.in/urfave/cli.v1"
)

var (
	versionCommand = cli.Command{
		Action:    versionAction,
		Name:      "version",
		Usage:     "Print version numbers",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
	}
)

func versionAction(*cli.Context) error {
	fmt.Printf(`orchestrator
Version:%v
Architecture:%v
Go Version:%v
Operating System:%v
GOPATH:%v
GOROOT:%v
Commit hash:%v
`, metadata.Version, runtime.GOARCH, runtime.Version(), runtime.GOOS, os.Getenv("GOPATH"), runtime.GOROOT(), metadata.GitCommit)
	return nil
}
