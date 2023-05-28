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
var grid [6][7]int 

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
	file, _ := os.CreateTemp(os.TempDir(), "*.png")
	logger.Info(file.Name())
	checkWin()
	proxy.Respond(provider.Response{ Content: "test"}, false, false, false)
	return nil
}

func generateImg(color color.RGBA, file *os.File) {
	img := image.NewRGBA(image.Rect(0, 0, 0, 0))
	
	png.Encode(file, img)
}

func checkWin() bool {
	return checkRows() || checkCols() || checkDiagonalsLeft() || checkDiagonalsRight()
}

func checkRows() bool {
	for i := 0; i < 6; i++{
		for j := 0; j < 7-3; j++ {
			if grid[i][j] == grid[i][j+1] && grid[i][j+1] == grid[i][j+2] && grid[i][j+2] == grid[i][j+3] && grid[i][j+3] != 0 { return true }
		}
	}
	return false
}

func checkCols() bool {
	for i := 0; i < 7; i++{
		for j := 0; j < 6-3; j++ {
			if grid[j][i] == grid[j+1][i] && grid[j+1][i] == grid[j+2][i] && grid[j+2][i] == grid[j+3][i] && grid[j+3][i] != 0 { return true }
		}
	}
	return false
}

func checkDiagonalsLeft() bool {
	for i := 3; i > 0; i-- {
		if fromUpperLeft(0, i) {
			return true
		}
	} 
	for i := 0; i < 3; i++ {
		if fromUpperLeft(i, 0) {
			return true
		}
	}
	return false
}

func checkDiagonalsRight() bool {
	for i := 3; i < 6; i++ {
		if fromUpperRight(0, i) {
			return true
		}
	} 
	for i := 0; i < 3; i++ {
		if fromUpperRight(i, 6) {
			return true
		}
	}
	return false
}

func fromUpperLeft(i,j int) bool {
	for i+3 <= 5 && j+3 <= 6 {
		if grid[i][j] == grid[i+1][j+1] && grid[i+1][j+1] == grid[i+2][j+2] && grid[i+2][j+2] == grid[i+3][j+3] && grid[i+3][j+3] != 0 {
			return true
		}
		j++
		i++
	}
	return false
}

func fromUpperRight(i,j int) bool {
	for i+3 <= 5 && j-3 >= 0 {
		if grid[i][j] == grid[i+1][j-1] && grid[i+1][j-1] == grid[i+2][j-2] && grid[i+2][j-2] == grid[i+3][j-3] && grid[i+3][j-3] != 0 {
			return true
		}
		j--
		i++
	}
	return false
}


// func (TestCommand) Autocomplete(provider.InteractionProxy) error {
// 	logger.Info("execute command")
// 	return nil
// }
