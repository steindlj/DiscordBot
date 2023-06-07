package component

import (
	"errors"
	"strconv"
	"strings"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
	"github.com/steindlj/dc-plugins/ConnectFour/message"
)

type ConnectFourSelectMenu struct{}

var _ common.Component = ConnectFourSelectMenu{}

// Returns the customID of this component so it can be assigned to the correct component.
func (ConnectFourSelectMenu) CustomID() (string, error) {
	return message.Prefix + "-colsm", nil
}

// Handles the input when this component is used.
func (ConnectFourSelectMenu) Handle(proxy common.HandleProxy) error {
	proxy.Defer(false)
	message.Proxy = proxy
	member, err := proxy.Member()
	if err != nil {
		message.ErrorEdit(err)
	}
	userId := member.User.ID
	// Resends message if component is used by non-players.
	// This is necessary so nobody can intefere.
	if !strings.EqualFold(userId, game.CurrPlayer.ID) {
		return message.ErrorEditPlayer(errors.New("unauthorized user tried to interact"))
	}
	colString, err := proxy.SelectedValues()
	if err != nil {
		message.ErrorEdit(err)
	}
	col, err := strconv.Atoi(colString[0])
	if err != nil {
		message.ErrorEdit(err)
	}
	image.ColorCell(game.PlaceChip(col))
	if game.CheckWin() {
		return message.WinMessage()
	}
	game.SetNextPlayer()
	return message.NewMessage()
}
