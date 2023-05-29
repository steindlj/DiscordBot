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

var (
	_ provider.Command = TestCommand{}
	permission int64 = discordgo.PermissionSendMessages
	grid [6][7]int
	space = 12
	width = 48
	background color.RGBA
	colorP1 = color.RGBA{255, 0, 0, 255}
	colorP2 = color.RGBA{255, 255, 0, 255}
	min = 0.0
) 



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
			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "chip_color",
				Description: "Yellow or Red",
				Required: true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name: "red",
						Value: 0,
					},
					{
						Name: "yellow",
						Value: 1,
					},
				},
			},
			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "red",
				Description: "Red value",
				Required: true,
				MinValue: &min,
				MaxValue: 255,
			},
			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "green",
				Description: "Green value",
				Required: true,
				MinValue: &min,
				MaxValue: 255,
			},
			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "blue",
				Description: "Blue value",
				Required: true,
				MinValue: &min,
				MaxValue: 255,
			},
			{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "alpha",
				Description: "Alpha value",
				Required: true,
				MinValue: &min,
				MaxValue: 255,
			},
		},
	}, nil
}

func (TestCommand) Execute(proxy provider.ExecuteProxy) error {
	file, _ := os.CreateTemp(os.TempDir(), "*.png")
	chipColor, _ := proxy.IntegerOption("chip_color")
	red, _ := proxy.IntegerOption("red")
	green, _ := proxy.IntegerOption("green")
	blue, _ := proxy.IntegerOption("blue")
	alpha, _ := proxy.IntegerOption("alpha")
	background = color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alpha)}
	if chipColor != 0 {
		colorP1 = color.RGBA{255, 255, 0, 255}
		colorP2 = color.RGBA{255, 0, 0, 255}
	}
	checkWin()
	generateImg(background, file)
	sendFile, _ := os.Open(file.Name())
	proxy.Respond(provider.Response{
		Files: []*discordgo.File{
			{
			Name: "image.png",
			Reader: sendFile,
			},
		},
	}, false, false, false)
	return nil
}

func setChip(player, col int) {
	for i := 5; i >= 0; i++ {
		if grid[i][col] == 0 {
			grid[i][col] = player
			return
		}
	}
}

func generateImg(c color.RGBA, file *os.File) {
	img := image.NewRGBA(image.Rect(0, 0, 7*width+8*space, 6*width+7*space))
	for i := 0; i < 7*width+8*space; i++ {
		for j := 0; j < 6*width+7*space; j++ {
			img.Set(i,j,c)
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			colorField(img, i, j)
		}
	} 
	png.Encode(file, img)
}

func colorField(img *image.RGBA, i, j int) {
	x := (j+1)*space + j*width
	y := (i+1)*space + i*width
	color := color.RGBA{background.R, background.G, background.B, 230}
	if grid[i][j] != 0 {
		if grid[i][j] == 1 {
			color = colorP1
		} else {
			color = colorP2
		}
	}
	for i = 0; i < width; i++ {
		for j = 0; j < width; j++ {
			img.SetRGBA(x+i, y+j, color)
		}
	}
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
	for i := 0; i <= 3; i++ {
		if i == 0 {
			if fromUpperLeft(i, 0) {
				return true
			}
		} else if i == 3 {
			if fromUpperLeft(0, i) {
				return true
			}
		} else {
			if fromUpperLeft(i, 0) || fromUpperLeft(0, i) {
				return true
			}
		}
	}
	return false
}

func checkDiagonalsRight() bool {
	for i := 0; i < 6; i++ {
		if i < 3 {
			if fromUpperRight(i, 6) {
				return true
			}
		} else {
			if fromUpperRight(0, i) {
				return true
			}
		}
	}
	return false
}

func fromUpperLeft(i,j int) bool {
	for i+3 <= 5 && j+3 <= 6 {
		if grid[i][j] == grid[i+1][j+1] && grid[i+1][j+1] == grid[i+2][j+2] && grid[i+2][j+2] == grid[i+3][j+3] && grid[i+3][j+3] != 0 {
			return true
		}
		i++
		j++
	}
	return false
}

func fromUpperRight(i,j int) bool {
	for i+3 <= 5 && j-3 >= 0 {
		if grid[i][j] == grid[i+1][j-1] && grid[i+1][j-1] == grid[i+2][j-2] && grid[i+2][j-2] == grid[i+3][j-3] && grid[i+3][j-3] != 0 {
			return true
		}
		i++
		j--
	}
	return false
}


// func (TestCommand) Autocomplete(provider.InteractionProxy) error {
// 	logger.Info("execute command")
// 	return nil
// }
