package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
)

type TestCommand struct{}

var _ provider.Command = TestCommand{}
var permission int64 = discordgo.PermissionSendMessages

func (TestCommand) CommandData() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "connectfour",
		Description: "Connect Four",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type: discordgo.ApplicationCommandOptionUser,
				Name: "oponent",
				Description: "Oponent",
				Required: true,
			}, 
			/* {
				Type: ?,
				Name: "chip color",
				Description "Yellow or Red"
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name: "Red",
						Value: ?,
					},
					{
						Name: "Yellow",
						Value: ?,
					}
				}
			},
			{
				Type: ?,
				Name: "background"
				Description: "Background color"
				Choices: []*discordgo.ApplicationCommandOptionChoice{

				}
			}, */
		},
	}, nil
}

func (TestCommand) Execute(proxy provider.ExecuteProxy) error {
	file, _ := os.CreateTemp("*.png")
	proxy.Respond(provider.Response{ Content: "test"}, false, false, false)
	return nil
}

func generateImg(color color.RGBA, grid [][]string, file *os.File) image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, , ))
	
	png.Encode(file, img)
}


// func (TestCommand) Autocomplete(provider.InteractionProxy) error {
// 	logger.Info("execute command")
// 	return nil
// }
