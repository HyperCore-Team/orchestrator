package app

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"runtime"
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
	app.Before = beforeAction
	app.Action = action
	app.After = afterAction
}
func beforeAction(ctx *cli.Context) error {
	if len(ctx.Args()) == 0 {
		max := runtime.NumCPU()
		fmt.Printf("Starting orchestrator.\n")
		runtime.GOMAXPROCS(max)
	}
	return nil
}
func action(ctx *cli.Context) error {
	//Make sure No subCommands were entered,Only the flags
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
