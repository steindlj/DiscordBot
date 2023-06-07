package main

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	plugin "github.com/EliasStar/BacoTell/pkg/bacotell_plugin"
	"github.com/hashicorp/go-hclog"
	"github.com/steindlj/dc-plugins/Text2Vocals/command"
	"github.com/steindlj/dc-plugins/Text2Vocals/message"
)

var logger hclog.Logger
var closeChan <-chan struct{}

var commands = []common.Command{
	command.TTSCommand{},
}

var components = []common.Component{}

func main() {
	plugin.SetApplicationCommands(commands...)
	plugin.SetMessageComponents(components...)

	logger, closeChan, _ = plugin.Run(message.Prefix)
	<-closeChan
}
