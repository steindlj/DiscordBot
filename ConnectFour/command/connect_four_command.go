package command

import (
	"image/color"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	util "github.com/EliasStar/BacoTell/pkg/bacotell_util"
	"github.com/bwmarrin/discordgo"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
	"github.com/steindlj/dc-plugins/ConnectFour/message"
)

type ConnectFourCommand struct{}

var _ common.Command = ConnectFourCommand{}


func (ConnectFourCommand) Data() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "connectfour",
		Description: "Connect Four",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "opponent",
				Description: "Opponent",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "chip_color",
				Description: "Red or Yellow",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "red",
						Value: 0,
					},
					{
						Name:  "yellow",
						Value: 1,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "red",
				Description: "Red value",
				Required:    true,
				MinValue:    util.Ptr(0.0),
				MaxValue:    255,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "green",
				Description: "Green value",
				Required:    true,
				MinValue:    util.Ptr(0.0),
				MaxValue:    255,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "blue",
				Description: "Blue value",
				Required:    true,
				MinValue:    util.Ptr(0.0),
				MaxValue:    255,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "alpha",
				Description: "Alpha value",
				Required:    true,
				MinValue:    util.Ptr(0.0),
				MaxValue:    255,
			},
		},
	}, nil
}

func (ConnectFourCommand) Execute(proxy common.ExecuteProxy) error {
	proxy.Defer(false)
	game.Grid = [6][7]int{}
	game.RoundCount = 1
	message.Proxy = proxy
	player1, err := proxy.Member()
	if err != nil {
		message.ErrorEdit(err)
	}
	game.Player1 = player1.User
	player2, err := proxy.UserOption("opponent")
	if err != nil {
		message.ErrorEdit(err)
	}
	game.Player2 = player2
	chipColor, err := proxy.IntegerOption("chip_color")
	if err != nil {
		message.ErrorEdit(err)
	}
	red, err := proxy.IntegerOption("red")
	if err != nil {
		message.ErrorEdit(err)
	}
	green, err := proxy.IntegerOption("green")
	if err != nil {
		message.ErrorEdit(err)
	}
	blue, err := proxy.IntegerOption("blue")
	if err != nil {
		message.ErrorEdit(err)
	}
	alpha, err := proxy.IntegerOption("alpha")
	if err != nil {
		message.ErrorEdit(err)
	}
	image.Background = color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alpha)}
	image.ColorBackground()
	if chipColor == 0 {
		image.ColorP1 = color.RGBA{255, 0, 0, uint8(alpha)}
		image.ColorP2 = color.RGBA{255, 255, 0, uint8(alpha)}
	} else {
		image.ColorP1 = color.RGBA{255, 255, 0, uint8(alpha)}
		image.ColorP2 = color.RGBA{255, 0, 0, uint8(alpha)}
	}
	game.CurrPlayer = player1.User
	return message.NewMessage()
}

// Autocomplete implements bacotell_common.Command.
func (ConnectFourCommand) Autocomplete(common.AutocompleteProxy) error {
	panic("unimplemented")
}
