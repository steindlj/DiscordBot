package message

import (
	"os"
	"strconv"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
)

var Proxy common.InteractionProxy

func NewMessage() error {
	Proxy.Delete("")
	_, err := Proxy.Followup(Response(basicTitle()), false)
	return err
}

func WinMessage() error {
	Proxy.Delete("")
	_, err := Proxy.Followup(common.Response{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: game.CurrPlayer.Username + " won!",
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://image.png",
				},
			},
		},
		Files: []*discordgo.File{
			{
				Name:   "image.png",
				Reader: newFile(),
			},
		},
	}, false)
	return err
}

func newFile() *os.File {
	file, err := os.CreateTemp(os.TempDir(), "*.png")
	if err != nil {
		ErrorEdit(err)
	}
	image.GenerateImg(file)
	sendFile, err := os.Open(file.Name())
	if err != nil {
		ErrorEdit(err)
	}
	return sendFile
}

func ErrorEditPlayer(error error) error {
	Proxy.Delete("")
	_, err := Proxy.Followup(Response(basicTitle() + "; Error: " + error.Error()), false)
	return err
}

// Changes content of discord message from current proxy to error message.
func ErrorEdit(error error) {
	Proxy.Edit("", common.Response{
		Content: error.Error(),
	})
}

// Returns the base title with the current players and their color.
func basicTitle() string {
	var p1Color string
	var p2Color string
	if image.ColorP1.G == 0 {
		p1Color = "(red)"
		p2Color = "(yellow)"
	} else {
		p1Color = "(yellow)"
		p2Color = "(red)"
	}
	return game.Player1.Mention()+p1Color+" vs. "+game.Player2.Mention()+p2Color
}
func Response(content string) common.Response {
	var turn string
	if []rune(game.CurrPlayer.Username)[len(game.CurrPlayer.Username)-1] == 's' {
		turn = " turn!"
	} else {
		turn = "'s turn!"
	}
	return common.Response{
		Content: content,
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: game.CurrPlayer.Username + turn,
				Description: "Round: " + strconv.Itoa(game.RoundCount),
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://image.png",
				},
			},
		},
		Files: []*discordgo.File{
			{
				Name:   "image.png",
				Reader: newFile(),
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					generateSelectMenu(),
				},
			},
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						CustomID: "btn",
						Label: "Random Color",
						Style: discordgo.SuccessButton,
					},
				},
			},
		},
	}
}

func generateSelectMenu() discordgo.SelectMenu {
	var options []discordgo.SelectMenuOption
	for i := range emptyCols() {
		for j := 5; j >= 0; j-- {
			if game.Grid[j][i] == 0 {
				options = append(options,
					discordgo.SelectMenuOption{
						Label: "Column: " + strconv.Itoa(i+1),
						Value: strconv.Itoa(i),
					})
				break
			}
		}
	}
	return discordgo.SelectMenu{
		CustomID: "colsm",
		MenuType: discordgo.StringSelectMenu,
		Options:  options,
	}
}

// Returns slice containing indices of the empty columns in the grid.
func emptyCols() []int {
	var cols []int
	for i := 0; i < 7; i++ {
		for j := 5; j >= 0; j-- {
			if game.Grid[j][i] == 0 {
				cols = append(cols, i)
				break
			}
		}
	}
	return cols
}
