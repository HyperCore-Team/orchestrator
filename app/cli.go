package app

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"orchestrator/metadata"

	"gopkg.in/urfave/cli.v1"
)

var (
	app         = cli.NewApp()
	nodeManager *Manager
)

func Run() {
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Stop() {
	err := nodeManager.Stop()
	if err != nil {
		panic(err)
	}
}

func init() {
	app.Name = filepath.Base(os.Args[0])
	app.HideVersion = false
	app.Version = metadata.Version
	app.Compiled = time.Now()
	app.Usage = "orchestrator Node"
	app.Commands = []cli.Command{
		versionCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = AllFlags
	app.Before = beforeAction
	app.Action = action
	app.After = afterAction
}
func beforeAction(ctx *cli.Context) error {

	max := runtime.NumCPU()
	fmt.Printf("Starting orchestrator.\n")
	fmt.Printf("current time is %v\n", time.Now().Format("2009-01-03 18:15:05"))
	fmt.Printf("version: %v\n", metadata.Version)
	fmt.Printf("git-commit-hash: %v\n", metadata.GitCommit)
	fmt.Printf("orchestrator will use at most %v cpu-cores\n", max)
	runtime.GOMAXPROCS(max)

	// pprof server
	if ctx.GlobalIsSet(PprofFlag.Name) {
		listenHost := ctx.String(PprofAddrFlag.Name)

		port := ctx.Int(PprofPortFlag.Name)

		address := fmt.Sprintf("%s:%d", listenHost, port)

		go func() {
			if err := http.ListenAndServe(address, nil); err != nil {
				nodeManager.logger.Error(err.Error())
			}
		}()
	}

	return nil
}
func action(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}
	var err error
	nodeManager, err = NewNodeManager(ctx)
	if err != nil {
		return err
	}

	return nodeManager.Start()

}
func afterAction(*cli.Context) error {
	return nil
}
