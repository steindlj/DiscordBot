package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

type TestCommand struct{}

var _ provider.Command = TestCommand{}
var permission int64 = discordgo.PermissionSendMessages

func (TestCommand) CommandData() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "tts",
		Description: "Text-to-Speech",
		DefaultMemberPermissions: &permission,
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
						Name: "sped up",
						Value: 1,
					},
					{
						Name: "slowed",
						Value: 2,
					},
				},
			},
		},
	}, nil
}

func (TestCommand) Execute(proxy provider.ExecuteProxy) error {
	text, _ := proxy.StringOption("text")
	lang, _ := proxy.StringOption("lang")
	style, _ := proxy.IntegerOption("effect")
	audioFile, _ := os.Open(Mix(CreateFile(text, lang), style))
	proxy.Respond(provider.Response{
		Content: text + " in " + lang,
		Files: []*discordgo.File{
			{
			Name: "audio.wav",
			Reader: audioFile,
			},
		},
	}, false, false, false)
	return nil
}

func CreateFile(text string, lang string) string {
	dir := os.TempDir()
	speech := htgotts.Speech{Folder: dir, Language: lang}
	name := generateHash(text + lang)
	speech.CreateSpeechFile(text, name)
	return dir + "\\" + name + ".mp3"
}

func Mix(filename string, style int64) string {
	file, _ := os.Open(filename)
	streamer, format, _ := mp3.Decode(file)
	wavFile, _ := os.CreateTemp(os.TempDir(), "*.wav")
	switch style {
	case 0:
		wav.Encode(wavFile, streamer, format)
	case 1:
		speed := beep.Resample(4, format.SampleRate, format.SampleRate*3/4, streamer)
		wav.Encode(wavFile, speed, format)
	case 2:
		slowed := beep.Resample(4, format.SampleRate, format.SampleRate*3/2, streamer)
		wav.Encode(wavFile, slowed, format)
	}
	wavFile.Close()
	return wavFile.Name()
}

func generateHash(name string) string {
	byte := md5.Sum([]byte(name))
	return hex.EncodeToString(byte[:])
}

// func (TestCommand) Autocomplete(provider.InteractionProxy) error {
// 	logger.Info("execute command")
// 	return nil
// }
