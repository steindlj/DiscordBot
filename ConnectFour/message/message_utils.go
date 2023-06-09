package message

import (
	"os"
	"strconv"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/steindlj/dc-plugins/ConnectFour/game"
	"github.com/steindlj/dc-plugins/ConnectFour/image"
)

var Proxy common.InteractionProxy // current proxy 
var Prefix = "connect_four"

// NewMessage sends a message with an updated game status including image, 
// round counter and current player. 
func NewMessage() error {
	Proxy.Delete("")
	_, err := Proxy.Followup(Response(basicTitle()), false)
	return err
}

// WinMessage sends the winner message announcing the winner and showing the final game state.
// Components are excluded since the game is finished. 
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

// newFile creates a new temporary png-file which will be encoded with
// the current state of the game.
func newFile() *os.File {
	file, err := os.CreateTemp(os.TempDir(), "*.png")
	if err != nil {
		ErrorEdit(err)
	}
	image.EncodeImage(file)
	sendFile, err := os.Open(file.Name())
	if err != nil {
		ErrorEdit(err)
	}
	return sendFile
}

// ErrorEditPlayer will only be called if a non-player user uses a command.
// It resends the message and points out the error in the content field of the response. 
func ErrorEditPlayer(error error) error {
	Proxy.Delete("")
	_, err := Proxy.Followup(Response(basicTitle() + "; Error: " + error.Error()), false)
	return err
}

// ErrorEdit changes the content of the discord message to an error message.
// The command becomes unusable. 
func ErrorEdit(err error) {
	Proxy.Edit("", common.Response{
		Content: err.Error(),
	})
}

// basicTitle returns the base title with the current players.
func basicTitle() string {
	return game.Player1.Mention()+" vs. "+game.Player2.Mention()
}

// Response returns the base response used for each round with all components
// and the image file showing the current state of the game.
// The content of the response is specified by the content parameter.
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
						CustomID: Prefix+"-btn",
						Label: "Random Color",
						Style: discordgo.SuccessButton,
					},
				},
			},
		},
	}
}

// generateSelectMenu returns a select menu containing all empty columns as an option to choose from.
func generateSelectMenu() discordgo.SelectMenu {
	var options []discordgo.SelectMenuOption
	for _,i := range emptyCols() {
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
		CustomID: Prefix+"-colsm",
		MenuType: discordgo.StringSelectMenu,
		Options:  options,
	}
}

// emptyCols returns a slice containing all indices of empty columns in the grid.
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
