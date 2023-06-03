package main

import (
	"strconv"
	"strings"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
	"github.com/steindlj/dc-plugins/ConnectFour/message"
)

type ConnectFourComponent struct{}

var _ common.Component = ConnectFourComponent{}

// CustomID implements bacotell_common.Component.
func (ConnectFourComponent) CustomID() (string, error) {
	return "colsm", nil
}

// Handle implements bacotell_common.Component.
func (ConnectFourComponent) Handle(proxy common.HandleProxy) error {
	proxy.Defer(true)
	message.Proxy = proxy
	member, err := proxy.Member()
	if err != nil {
		message.ErrorEdit(err)
	}
	userId := member.User.ID
	if !strings.EqualFold(userId, Player1Id) && !strings.EqualFold(userId, Player2Id) {
		message.ErrorEdit(err)
	}
	colString, err := proxy.SelectedValues()
	if err != nil {
		message.ErrorEdit(err)
	}
	col, err := strconv.Atoi(colString[0])
	if err != nil {
		message.ErrorEdit(err)
	}
	image.ColorField(game.SetChip(userId, col))
	if game.CheckWin() {
		return message.WinMessage()
	}
	return message.NewMessage()
}

