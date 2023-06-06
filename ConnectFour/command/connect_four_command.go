package command

import (
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	util "github.com/EliasStar/BacoTell/pkg/bacotell_util"
	"github.com/bwmarrin/discordgo"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
	"github.com/steindlj/dc-plugins/ConnectFour/message"
)

type ConnectFourCommand struct{}

var _ common.Command = ConnectFourCommand{}

// Defines structure of command.
func (ConnectFourCommand) Data() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        message.Prefix+"-connect_four",
		Description: "Connect Four",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishUS: "connect_four",
			discordgo.German: "4_gewinnt",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "opponent",
				Description: "Opponent",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "your_color",
				Description: "Chip color",
				MinValue:    util.Ptr(0.0),
				MaxValue:    0xFFFFFF,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "opponent_color",
				Description: "Chip color",
				MinValue:    util.Ptr(0.0),
				MaxValue:    0xFFFFFF,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "grid",
				Description: "Grid value",
				MinValue:    util.Ptr(0.0),
				MaxValue:    0xFFFFFF,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "cell",
				Description: "Cell value",
				MinValue:    util.Ptr(0.0),
				MaxValue:    0xFFFFFF,
			},
		},
	}, nil
}

// Execution of command.
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
	image.ColorP1, err = proxy.IntegerOption("your_color")
	if err != nil {
		image.ColorP1 = 0xFF0000
	} 
	image.ColorP2, err = proxy.IntegerOption("opponent_color")
	if err != nil {
		image.ColorP2 = 0xFFFF00
	} 
	image.Grid, err = proxy.IntegerOption("grid")
	if err != nil {
		image.Grid = 0x0000FF
	}
	image.Cell, err = proxy.IntegerOption("cell")
	if err != nil {
		image.Cell = 0xFFFFFF
	}
	image.GenerateImg()
	game.CurrPlayer = player1.User
	return message.NewMessage()
}

// Has to be implented but is not used by this command.
func (ConnectFourCommand) Autocomplete(common.AutocompleteProxy) error {
	panic("unimplemented")
}
