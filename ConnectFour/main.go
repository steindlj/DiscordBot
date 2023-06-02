package main

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	plugin "github.com/EliasStar/BacoTell/pkg/bacotell_plugin"
	"github.com/hashicorp/go-hclog"
)

var logger hclog.Logger
var closeChan <-chan struct{}

var commands = []common.Command{
	ConnectFourCommand{},
}

var components = []common.Component{}

func main() {
	plugin.SetApplicationCommands(commands...)
	plugin.SetMessageComponents(components...)

	logger, closeChan, _ = plugin.Debug("connect_four", "MTA4OTk3MjA5MTE0MDQ0NDIwMA.GD1PSP.3i4VfMmnHuPjenyFgUUxkpDoJqw_zC_pW1sMsQ")
	<-closeChan
}
