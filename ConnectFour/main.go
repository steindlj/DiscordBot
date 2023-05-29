package main

import (
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/EliasStar/BacoTell/pkg/bacotell_plugin"
	"github.com/hashicorp/go-hclog"
)

var logger hclog.Logger
var closeChan <-chan struct{}

var commands = []bacotell.Command{
	ConnectFourCommand{},
}

var components = []bacotell.Component{}

func main() {
	bacotell_plugin.SetApplicationCommands(commands)
	bacotell_plugin.SetMessageComponents(components)

	logger, closeChan, _ = bacotell_plugin.Debug("connect_four", "MTA4OTk3MjA5MTE0MDQ0NDIwMA.GD1PSP.3i4VfMmnHuPjenyFgUUxkpDoJqw_zC_pW1sMsQ")
	<-closeChan
}
