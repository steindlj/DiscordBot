package command

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/PerformLine/go-stockutil/colorutil"
	"github.com/bwmarrin/discordgo"
)

type LomoPurpleCommand struct{}

var _ common.Command = LomoPurpleCommand{}
var currProxy common.ExecuteProxy
var Prefix = "image_filter"

// Defines structure of command.
func (LomoPurpleCommand) Data() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        Prefix+"-lomo_filter",
		Description: "Lomo Purple Filter",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishUS: "lomo_filter",
			discordgo.German: "lomo_filter",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionAttachment,
				Name:        "attachment",
				Description: "Attached image",
				Required:    true,
			},
		},
	}, nil
}

// Execution of command.
func (LomoPurpleCommand) Execute(proxy common.ExecuteProxy) error {
	proxy.Defer(false)
	currProxy = proxy
	img, err := proxy.AttachmentOption("attachment")
	if err != nil {
		return errorEdit(err)
	}

	url := img.URL
	tempDir := "temp"
	path, err := downloadImage(url, tempDir)
	if err != nil {
		return errorEdit(err)
	}
	grid, err := load(path)
	if err != nil {
		return errorEdit(err)
	}
	newPath := save(tempDir, img.Filename, filter(grid))

	sendImg, err := os.Open(newPath)
	if err != nil {
		return errorEdit(err)
	}
	defer sendImg.Close()

	proxy.Followup(common.Response{
		Files: []*discordgo.File{
			{
				Name:   img.Filename,
				Reader: sendImg,
			},
		},
	}, false)
	deleteDir(tempDir)
	return nil
}

// This function load the image from the directory and create an array 
// from the image with the color values.
func load(filePath string) ([][]color.Color, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	if img.ColorModel() != color.YCbCrModel {
		return nil, errors.New("not a jpg")
	}

	bounds := img.Bounds()
	xlen, ylen := bounds.Max.X, bounds.Max.Y

	imgArray := make([][]color.Color, xlen)
	for x := 0; x < xlen; x++ {
		imgArray[x] = make([]color.Color, ylen)
		for y := 0; y < ylen; y++ {
			imgArray[x][y] = img.At(x, y)
		}
	}
	return imgArray, nil
}

// Has to be implented but is not used by this command.
func (LomoPurpleCommand) Autocomplete(common.AutocompleteProxy) error {
	panic("unimplemented")
}

func save(directory string, fileName string, grid [][]color.Color) string {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}

	filePath := filepath.Join(directory, "IR_"+fileName)
	file, err := os.Create(filePath)
	if err != nil {
		errorEdit(err)
		return ""
	}
	defer file.Close()
	jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	return filePath
}

// This function will go through each pixel of the saved image and convered the colors
// and save it back into a new array.
func filter(grid [][]color.Color) (irImage [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	irImage = make([][]color.Color, xlen)
	for i := 0; i < len(irImage); i++ {
		irImage[i] = make([]color.Color, ylen)
	}
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			pix := grid[x][y].(color.YCbCr)
			R, G, B := color.YCbCrToRGB(pix.Y, pix.Cb, pix.Cr)

			hue, sat, light := colorutil.RgbToHsl(float64(R), float64(G), float64(B))

			//Red -> Orange
			if hue >= 0 && hue <= 15 {
				hue += 30
				//Yellow -> Magenta
			} else if hue >= 50 && hue <= 75 {
				hue += 250
				//Green -> Magenta
			} else if hue > 75 && hue <= 150 {
				hue += 225
				//Blue -> Green
			} else if hue >= 165 && hue <= 255 {
				hue -= 75
			}

			newR, newG, newB := colorutil.HslToRgb(hue, sat, light)
			irImage[x][y] = color.NRGBA{newR, newG, newB, 255}
		}
	}
	return
}

// This function downloads the image from user to a temporary directory
// to process the image later.
func downloadImage(url string, directory string) (string, error) {
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := filepath.Base(url)
	filePath := filepath.Join(directory, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New(strconv.Itoa(response.StatusCode))
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// This function will delete the temporary directory after the image was send
// back to the user. 
func deleteDir(directory string) error {
	err := os.RemoveAll(directory)
	if err != nil {
		return err
	}
	return nil
}

// ErrorEdit changes the content of the discord message from the current proxy to an error message.
// This function makes the command unusable.
func errorEdit(err error) error {
	return currProxy.Edit("", common.Response{
		Content: err.Error(),
	})
}