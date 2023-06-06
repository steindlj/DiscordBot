package message

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	htgotts "github.com/hegedustibor/htgo-tts"
)

var Proxy common.InteractionProxy

// 
func CreateFile(text string, lang string) string {
	dir := os.TempDir()
	speech := htgotts.Speech{Folder: dir, Language: lang}
	name := generateHash(text + lang)
	speech.CreateSpeechFile(text, name)
	return dir + "\\" + name + ".mp3"
}

// Generates hashstring from given string.
func generateHash(name string) string {
	byte := md5.Sum([]byte(name))
	return hex.EncodeToString(byte[:])
}

// Changes content of discord message from current proxy to error message.
func ErrorEdit(error error) {
	Proxy.Edit("", common.Response{
		Content: error.Error(),
	})
}
