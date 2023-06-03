package audio

import (
	"math/rand"
	"os"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	gowav "github.com/go-audio/wav"
	"github.com/steindlj/dc-plugins/Text2Vocals/message"
)

func Mix(filename string, effect int64) string {
	mp3File, err := os.Open(filename)
	if err != nil {
		message.ErrorEdit(err)
	}
	defer mp3File.Close()
	wavFile, err := os.Open(converToWAV(mp3File))
	if err != nil {
		message.ErrorEdit(err)
	}
	defer wavFile.Close()
	buffer, err := gowav.NewDecoder(wavFile).FullPCMBuffer()
	if err != nil {
		message.ErrorEdit(err)
	}
	switch effect {
	case 0: // default
		return wavFile.Name()
	case 1: // distortion
		for i := range buffer.Data {
			buffer.Data[i] *= 5
		}
	case 2: // vintage
		for i := range buffer.Data {
			buffer.Data[i] = int(float64(buffer.Data[i]) * 0.8)
			buffer.Data[i] += rand.Intn(400) - 200
		}
	case 3: // slowed
		buffer.Format.SampleRate /= 2
	case 4: // sped up
		buffer.Format.SampleRate *= 2
	}
	newFile, err := os.CreateTemp(os.TempDir(), "*.wav")
	if err != nil {
		message.ErrorEdit(err)
	}
	encoder := gowav.NewEncoder(newFile, buffer.Format.SampleRate, buffer.SourceBitDepth, buffer.Format.NumChannels, 1)
	encoder.Write(buffer)
	encoder.Close()
	return newFile.Name()
}

func converToWAV(file *os.File) string {
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		message.ErrorEdit(err)
	}
	wavFile, err := os.CreateTemp(os.TempDir(), "*.wav")
	if err != nil {
		message.ErrorEdit(err)
	}
	wav.Encode(wavFile, streamer, format)
	return wavFile.Name()
}