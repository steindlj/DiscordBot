package command

import (
	"os"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/steindlj/dc-plugins/Text2Vocals/message"
	"github.com/steindlj/dc-plugins/Text2Vocals/audio"
)

type TTSCommand struct{}

var _ common.Command = TTSCommand{}

// Defines structure of command.
func (TTSCommand) Data() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "tts",
		Description: "Text-to-Speech",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "text",
				Description: "Input text",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "lang",
				Description: "Language",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "English",
						Value: voices.English,
					},
					{
						Name:  "German",
						Value: voices.German,
					},
					{
						Name:  "Chinese",
						Value: voices.Chinese,
					},
					{
						Name:  "Latin",
						Value: voices.Latin,
					},
					{
						Name:  "African",
						Value: voices.Afrikaans,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "effect",
				Description: "Mixing effect",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "default",
						Value: 0,
					},
					{
						Name:  "distortion",
						Value: 1,
					},
					{
						Name:  "vintage",
						Value: 2,
					},
					{
						Name:  "slowed",
						Value: 3,
					},
					{
						Name:  "sped_up",
						Value: 4,
					},
				},
			},
		},
	}, nil
}

// Execution of command.
func (TTSCommand) Execute(proxy common.ExecuteProxy) error {
	proxy.Defer(false)
	message.Proxy = proxy 
	text, err := proxy.StringOption("text")
	if err != nil {
		message.ErrorEdit(err)
	}
	lang, err := proxy.StringOption("lang")
	if err != nil {
		message.ErrorEdit(err)
	}
	effect, err := proxy.IntegerOption("effect")
	if err != nil {
		message.ErrorEdit(err)
	}
	fileToSend, err := os.Open(audio.Mix(message.CreateFile(text, lang), effect))
	if err != nil {
		message.ErrorEdit(err)
	}
	defer fileToSend.Close()
	return proxy.Edit("", common.Response{
		Content: "\"" + text + "\" in " + lang,
		Files: []*discordgo.File{
			{
				Name:   "audio.wav",
				Reader: fileToSend,
			},
		},
	})
}

// Has to be implented but is not used by this command.
func (TTSCommand) Autocomplete(common.AutocompleteProxy) error {
	panic("unimplemented")
}

