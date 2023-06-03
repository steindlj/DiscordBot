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

var components = []common.Component{
	ConnectFourComponent{},
}

func main() {
	plugin.SetApplicationCommands(commands...)
	plugin.SetMessageComponents(components...)

	logger, closeChan, _ = plugin.Debug("connect_four", "MTA4OTk3MjA5MTE0MDQ0NDIwMA.GP2KIy.leSaPM-gs15w2o44798L41qj68-drv2zDvigSk")
	<-closeChan
}
