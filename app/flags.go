package app

import (
	"gopkg.in/urfave/cli.v1"
)

var (

	// pprof

	PprofFlag = cli.BoolFlag{
		Name:  "pprof",
		Usage: "Enable the pprof HTTP server",
	}
	PprofPortFlag = cli.Uint64Flag{
		Name:  "pprof.port",
		Usage: "pprof HTTP server listening port",
		Value: 6060,
	}

	PprofAddrFlag = cli.StringFlag{
		Name:  "pprof.addr",
		Usage: "pprof HTTP server listening interface",
		Value: "127.0.0.1",
	}

	AllFlags = []cli.Flag{

		// pprof
		PprofFlag,
		PprofPortFlag,
		PprofAddrFlag,
	}
)
