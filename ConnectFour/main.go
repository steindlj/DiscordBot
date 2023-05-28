package main

import (
	"os"

	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-hclog"
)

var logger = hclog.New(&hclog.LoggerOptions{
	Name:   "connect_four",
	Output: os.Stdout,
	Level:  hclog.Debug,
})

var commands = []provider.Command{
	TestCommand{},
}

var components = []provider.Component{
	TestComponent{},
}

func main() {
	bacotell.SetInteractionProvider(provider.NewInteractionProvider("connect_four", commands, components))
	bacotell.DebugPlugin(logger, "MTA4OTk4NTcwODc1Nzg5MzE4MQ.Gm0ijG.KOn1Xus4UjaM9W0jxt2DXaV68S319WY0qWUJ80")
}
