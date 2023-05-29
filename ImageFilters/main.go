package main

import (
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/EliasStar/BacoTell/pkg/bacotell_plugin"
	"github.com/hashicorp/go-hclog"
)

var logger hclog.Logger
var closeChan <-chan struct{}

var commands = []bacotell.Command{
	LomoPurpleCommand{},
}

var components = []bacotell.Component{}

func main() {
	bacotell_plugin.SetApplicationCommands(commands)
	bacotell_plugin.SetMessageComponents(components)

	logger, closeChan, _ = bacotell_plugin.Debug("image_filters", "MTA4OTk4MzgzMDgwMDIxNjIxNA.G97SCC.1bHce01KOL4w1ybgzcOVDQYGgtR3fxW_PmHs4E")
	<-closeChan
}
