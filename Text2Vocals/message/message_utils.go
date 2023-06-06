package message

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	htgotts "github.com/hegedustibor/htgo-tts"
)

var Proxy common.InteractionProxy // current proxy

// Creates text-to-speech file from given text in given lang (speaker language, not translation).
// Returns path of generated mp3-file.
func CreateFile(text string, lang string) string {
	dir := os.TempDir()
	speech := htgotts.Speech{Folder: dir, Language: lang}
	name := generateHash(text + lang)
	speech.CreateSpeechFile(text, name)
	return dir + "\\" + name + ".mp3"
}

// Generates hashstring from given string and returns it.
func generateHash(name string) string {
	byte := md5.Sum([]byte(name))
	return hex.EncodeToString(byte[:])
}

// Changes content of discord message from the current proxy to error message.
// Command will become unusable. 
func ErrorEdit(error error) {
	Proxy.Edit("", common.Response{
		Content: error.Error(),
	})
}
