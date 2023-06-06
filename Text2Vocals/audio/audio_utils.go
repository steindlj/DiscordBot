package audio

import (
	"math/rand"
	"os"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/wav"
	gowav "github.com/go-audio/wav"
	"github.com/steindlj/dc-plugins/Text2Vocals/message"
)

// Mix decodes file specified by the filepath into pcm values
// and modifies them based on the desired effect. 
// It returns the file path of the new file with the modified sound.
func Mix(filepath string, effect int64) string {
	mp3File, err := os.Open(filepath)
	if err != nil {
		message.ErrorEdit(err)
	}
	defer mp3File.Close()
	wavFile, err := os.Open(converToWAV(mp3File))
	if err != nil {
		message.ErrorEdit(err)
	}
	defer wavFile.Close()
	// Decodes the WAV file into PCM buffer
	buffer, err := gowav.NewDecoder(wavFile).FullPCMBuffer()
	if err != nil {
		message.ErrorEdit(err)
	}
	switch effect {
	case 0: // distortion: multiplying the pcm value by 5 --> will often exceed limit --> clipping occurs
		for i := range buffer.Data {
			buffer.Data[i] *= 5
		}
	case 1: // vintage/old recording: changing the pcm by random values between -200 and +199 to create background noise
		for i := range buffer.Data {
			buffer.Data[i] += rand.Intn(400) - 200
		}
	case 2: // slowed: decreasing sample rate --> less samples per second --> sounds slower
		buffer.Format.SampleRate /= 2
	case 3: // sped up: increasing sampe rate --> more samples per second --> sounds faster
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

// Decodes the specified MP3-file and returns the file path of the converted WAV-file.
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