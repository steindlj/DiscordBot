package main

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	plugin "github.com/EliasStar/BacoTell/pkg/bacotell_plugin"
	"github.com/hashicorp/go-hclog"
)

var logger hclog.Logger
var closeChan <-chan struct{}

var commands = []common.Command{
	TTSCommand{},
}

var components = []common.Component{}

func main() {
	plugin.SetApplicationCommands(commands)
	plugin.SetMessageComponents(components)

	logger, closeChan, _ = plugin.Debug("text2vocals", "MTA4OTk4NTcwODc1Nzg5MzE4MQ.Gm0ijG.KOn1Xus4UjaM9W0jxt2DXaV68S319WY0qWUJ80")
	<-closeChan
}
