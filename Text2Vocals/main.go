package main

import (
	"os"

	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-hclog"
)

var logger = hclog.New(&hclog.LoggerOptions{
	Name:   "text2vocals",
	Output: os.Stdout,
	Level:  hclog.Debug,
})

var commands = []provider.Command{}

var components = []provider.Component{}

func main() {
	bacotell.SetInteractionProvider(provider.NewInteractionProvider("text2vocals", commands, components))
	bacotell.DebugPlugin(logger, os.Getenv("BACOTELL_BOT_TOKEN"))
}
