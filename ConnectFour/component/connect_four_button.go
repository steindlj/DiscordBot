package component

import (
	"image/color"
	"math/rand"
	"strings"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
	"github.com/steindlj/dc-plugins/ConnectFour/message"
)

type ConnectFourButton struct{}

var _ common.Component = ConnectFourButton{}

// CustomID implements bacotell_common.Component.
func (ConnectFourButton) CustomID() (string, error) {
	return "btn", nil
}

// Handle implements bacotell_common.Component.
func (ConnectFourButton) Handle(proxy common.HandleProxy) error {
	proxy.Defer(false)
	message.Proxy = proxy
	member, err := proxy.Member()
	if err != nil {
		message.ErrorEdit(err)
	}
	userId := member.User.ID
	if !strings.EqualFold(userId, game.Player1.ID) && !strings.EqualFold(userId, game.Player2.ID) {
		message.ErrorEdit(err)
	}
	red := uint8(rand.Intn(256))
	green := uint8(rand.Intn(256))
	blue := uint8(rand.Intn(256))
	alpha := uint8(rand.Intn(256))
	image.Background = color.RGBA{red, blue, green, alpha}
	image.ColorBackground()
	return message.NewMessage()
}
