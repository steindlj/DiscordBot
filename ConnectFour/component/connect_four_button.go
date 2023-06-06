package component

import (
	"errors"
	"math/rand"
	"strings"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
	"github.com/steindlj/dc-plugins/ConnectFour/message"
)

type ConnectFourButton struct{}

var _ common.Component = ConnectFourButton{}

// Returns the customID of this component so it can be assigned to the correct component.
func (ConnectFourButton) CustomID() (string, error) {
	return "btn", nil
}

// Handles the input when this component is used.
func (ConnectFourButton) Handle(proxy common.HandleProxy) error {
	proxy.Defer(false)
	message.Proxy = proxy
	member, err := proxy.Member()
	if err != nil {
		message.ErrorEdit(err)
	}
	userId := member.User.ID
	// Resends message if component is used by non-players.
	// This is necessary so nobody can intefere.
	if !strings.EqualFold(userId, game.Player1.ID) && !strings.EqualFold(userId, game.Player2.ID) {
		return message.ErrorEditPlayer(errors.New("unauthorized user tried to interact"))
	}
	red := rand.Intn(256)
	green := rand.Intn(256)
	blue := rand.Intn(256)
	image.Grid = int64((red << 24) | (green << 16) | (blue << 8) | 255)
	image.GenerateImg()
	return message.NewMessage()
}
