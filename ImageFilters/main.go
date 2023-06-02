package main

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	plugin "github.com/EliasStar/BacoTell/pkg/bacotell_plugin"
	"github.com/hashicorp/go-hclog"
)

var logger hclog.Logger
var closeChan <-chan struct{}

var commands = []common.Command{
	LomoPurpleCommand{},
}

var components = []common.Component{}

func main() {
	plugin.SetApplicationCommands(commands...)
	plugin.SetMessageComponents(components...)

	logger, closeChan, _ = plugin.Debug("image_filters", "MTA4OTk4MzgzMDgwMDIxNjIxNA.G97SCC.1bHce01KOL4w1ybgzcOVDQYGgtR3fxW_PmHs4E")
	<-closeChan
}
