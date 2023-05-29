package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"

	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"

	//"github.com/go-audio/audio"
	gowav "github.com/go-audio/wav"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

type Text2Vocals struct{}

var _ provider.Command = Text2Vocals{}
var permission int64 = discordgo.PermissionSendMessages

func (Text2Vocals) CommandData() (discordgo.ApplicationCommand, error) {
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

func (Text2Vocals) Execute(proxy provider.ExecuteProxy) error {
	text, _ := proxy.StringOption("text")
	lang, _ := proxy.StringOption("lang")
	//style, _ := proxy.IntegerOption("effect")
	audioFile, _ := os.Open(Mix(CreateFile(text, lang), 4))
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
	mp3File, _ := os.Open(filename)
	defer mp3File.Close()
	wavFile, _ := os.Open(converToWAV(mp3File))
	defer wavFile.Close()
	buffer, _ := gowav.NewDecoder(wavFile).FullPCMBuffer()
	switch style {
	case 0:
		return wavFile.Name()
	case 1:
		for i := range buffer.Data {
			buffer.Data[i] *= 5
		}
	case 2:
		for i := range buffer.Data {
			buffer.Data[i] = int(float64(buffer.Data[i])*0.9)
			buffer.Data[i] /= 2
			buffer.Data[i] += rand.Intn(100) - 50
		}
	case 3:
		buffer.Format.SampleRate /= 2 
	case 4: 
		buffer.Format.SampleRate *= 2
	}
	newFile, _ := os.CreateTemp(os.TempDir(), "*.wav")
	encoder := gowav.NewEncoder(newFile, buffer.Format.SampleRate, buffer.SourceBitDepth, buffer.Format.NumChannels, 1)
	encoder.Write(buffer)
	encoder.Close()
	return newFile.Name()
}

func converToWAV(file *os.File) string {
	streamer, format, _ := mp3.Decode(file)
	wavFile, _ := os.CreateTemp(os.TempDir(), "*.wav")
	wav.Encode(wavFile, streamer, format)
	return wavFile.Name()
}

func generateHash(name string) string {
	byte := md5.Sum([]byte(name))
	return hex.EncodeToString(byte[:])
}

// func (Text2Vocals) Autocomplete(provider.InteractionProxy) error {
// 	logger.Info("execute command")
// 	return nil
// }
