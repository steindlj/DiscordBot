package main

import (
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/EliasStar/BacoTell/pkg/bacotell_plugin"
	"github.com/hashicorp/go-hclog"
)

var logger hclog.Logger
var closeChan <-chan struct{}

var commands = []bacotell.Command{
	TTSCommand{},
}

var components = []bacotell.Component{}

func main() {
	bacotell_plugin.SetApplicationCommands(commands)
	bacotell_plugin.SetMessageComponents(components)

	logger, closeChan, _ = bacotell_plugin.Debug("text2vocals", "MTA4OTk4NTcwODc1Nzg5MzE4MQ.Gm0ijG.KOn1Xus4UjaM9W0jxt2DXaV68S319WY0qWUJ80")
	<-closeChan
}
