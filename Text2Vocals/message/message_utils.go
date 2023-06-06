package message

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	htgotts "github.com/hegedustibor/htgo-tts"
)

var Proxy common.InteractionProxy // current proxy
var Prefix = "text2vocals"

// CreateFile generates a text-to-speech file from the given text in the specified language (speaker language, not translation).
// It returns the path of the generated mp3-file.
func CreateFile(text string, lang string) string {
	dir := os.TempDir()
	speech := htgotts.Speech{Folder: dir, Language: lang}
	name := generateHash(text + lang)
	speech.CreateSpeechFile(text, name)
	return dir + "\\" + name + ".mp3"
}

// generateHash generates hashstring from given string and returns it.
func generateHash(name string) string {
	byte := md5.Sum([]byte(name))
	return hex.EncodeToString(byte[:])
}

// ErrorEdit changes the content of the discord message from the current proxy to an error message.
// This function makes the command unusable.
func ErrorEdit(error error) {
	Proxy.Edit("", common.Response{
		Content: error.Error(),
	})
}
