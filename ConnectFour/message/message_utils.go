package message

import (
	"os"
	"strconv"
	"strings"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
)

var Proxy common.InteractionProxy

func NewMessage() error {
	Proxy.Delete("")
	_, err := Proxy.Followup(Response(game.Player1.Mention()+" vs. "+game.Player2.Mention()), false)
	return err
}

func WinMessage() error {
	Proxy.Delete("")
	var winner *discordgo.User
	if strings.EqualFold(game.CurrPlayer.ID, game.Player1.ID) {
		winner = game.Player2
	} else {
		winner = game.Player1
	}
	_, err := Proxy.Followup(common.Response{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: winner.Username + " won!",
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
	Proxy.Followup(Response(game.Player1.Mention()+" vs. "+game.Player2.Mention() + "; Error: " + error.Error()), false)
	return nil
}

func ErrorEdit(error error) {
	Proxy.Edit("", common.Response{
		Content: error.Error(),
	})
}

func Response(content string) common.Response {
	return common.Response{
		Content: content,
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: game.CurrPlayer.Username + "'s turn!",
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
