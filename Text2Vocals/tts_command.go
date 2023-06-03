package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	gowav "github.com/go-audio/wav"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

type TTSCommand struct{}

var _ common.Command = TTSCommand{}

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

func (TTSCommand) Execute(proxy common.ExecuteProxy) error {
	proxy.Defer(true)
	text, err := proxy.StringOption("text")
	if err != nil {
		logger.Info("StringOption error for text", "err", err)
	}
	lang, err := proxy.StringOption("lang")
	if err != nil {
		logger.Info("StringOption error for lang", "err", err)
	}
	effect, err := proxy.IntegerOption("effect")
	if err != nil {
		logger.Info("IntegerOption error for effect", "err", err)
	}
	fileToSend, err := os.Open(Mix(CreateFile(text, lang), effect))
	if err != nil {
		logger.Info("Error opening file", "err", err)
	}
	defer fileToSend.Close()
	proxy.Followup(common.Response{
		Content: text + " in " + lang,
		Files: []*discordgo.File{
			{
				Name:   "audio.wav",
				Reader: fileToSend,
			},
		},
	}, false)
	return nil
}

// Autocomplete implements bacotell_common.Command.
func (TTSCommand) Autocomplete(common.AutocompleteProxy) error {
	panic("unimplemented")
}

func CreateFile(text string, lang string) string {
	dir := os.TempDir()
	speech := htgotts.Speech{Folder: dir, Language: lang}
	name := generateHash(text + lang)
	speech.CreateSpeechFile(text, name)
	return dir + "\\" + name + ".mp3"
}

func Mix(filename string, effect int64) string {
	mp3File, err := os.Open(filename)
	if err != nil {
		logger.Info("Error opening mp3", "err", err)
	}
	defer mp3File.Close()
	wavFile, err := os.Open(converToWAV(mp3File))
	if err != nil {
		logger.Info("Error opening wav", "err", err)
	}
	defer wavFile.Close()
	buffer, err := gowav.NewDecoder(wavFile).FullPCMBuffer()
	if err != nil {
		logger.Info("Error getting PCMBuffer", "err", err)
	}
	switch effect {
	case 0: // default
		return wavFile.Name()
	case 1: // distortion
		for i := range buffer.Data {
			buffer.Data[i] *= 5
		}
	case 2: // "vintage"
		for i := range buffer.Data {
			buffer.Data[i] = int(float64(buffer.Data[i]) * 0.8)
			buffer.Data[i] += rand.Intn(100) - 50
		}
	case 3: // slowed
		buffer.Format.SampleRate /= 2
	case 4: // sped up
		buffer.Format.SampleRate *= 2
	}
	newFile, err := os.CreateTemp(os.TempDir(), "*.wav")
	if err != nil {
		logger.Info("Error creating temp wav", "err", err)
	}
	encoder := gowav.NewEncoder(newFile, buffer.Format.SampleRate, buffer.SourceBitDepth, buffer.Format.NumChannels, 1)
	encoder.Write(buffer)
	encoder.Close()
	return newFile.Name()
}

func converToWAV(file *os.File) string {
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		logger.Info("Error decoding mp3", "err", err)
	}
	wavFile, err := os.CreateTemp(os.TempDir(), "*.wav")
	if err != nil {
		logger.Info("Error creating temp wav", "err", err)
	}
	wav.Encode(wavFile, streamer, format)
	return wavFile.Name()
}

func generateHash(name string) string {
	byte := md5.Sum([]byte(name))
	return hex.EncodeToString(byte[:])
}
