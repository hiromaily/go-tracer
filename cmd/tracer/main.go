package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hiromaily/go-tracer/pkg/config"
	"github.com/hiromaily/go-tracer/pkg/server"
	"github.com/hiromaily/go-tracer/pkg/tracer"
)

var (
	tomlPath = flag.String("t", "", "Toml file path")
)

var usage = `Usage: %s [options...]
Options:
  -t      Toml file path for config
`

func parseFlag() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	//command-line
	flag.Parse()
}

func main() {
	parseFlag()

	//config
	conf, err := config.New(*tomlPath)
	if err != nil {
		panic(err)
	}
	trs := tracer.NewTracer(conf.Tracer)
	//opentracing.SetGlobalTracer(trs)

	srv := server.NewServer(conf.Server.Port, trs)
	srv.Listen()
}
