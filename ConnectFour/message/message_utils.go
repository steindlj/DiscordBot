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
	return Proxy.Edit("", common.Response{
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
	})
}

func WinMessage() error {
	return Proxy.Edit("", common.Response{
		Files: []*discordgo.File{
			{
				Name:   "image.png",
				Reader: newFile(),
			},
		},
	}) 
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

func ErrorEdit(error error) {
	Proxy.Edit("", common.Response{
		Content: error.Error(),
	})
}

func ErrorRespond(proxy common.InteractionProxy, error error, ephemeral bool) {
	proxy.Respond(common.Response{
		Content: error.Error(),
	}, ephemeral)
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
				break;
			}
		}
	}
	return discordgo.SelectMenu{
		CustomID: "colsm",
		MenuType: discordgo.StringSelectMenu,
		Options: options,
	}
}

func emptyCols() []int {
	var cols []int
	for i := 0; i < 7; i++ {
		for j := 5; j >= 0; j-- {
			if game.Grid[j][i] == 0 {
				cols = append(cols, i)
				break;
			}
		}
	}
	return cols
}