package main

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/PerformLine/go-stockutil/colorutil"
	"github.com/bwmarrin/discordgo"

	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type LomoPurpleCommand struct{}

var _ bacotell.Command = LomoPurpleCommand{}

// CommandData implements provider.Command
func (LomoPurpleCommand) CommandData() (discordgo.ApplicationCommand, error) {
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "lomo_filter",
		Description: "Lomo Purple Filter",
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

// Execute implements provider.Command
func (LomoPurpleCommand) Execute(proxy bacotell.ExecuteProxy) error {
	img, err := proxy.AttachmentOption("attachment")
	if err != nil {
		logger.Info("Cannot find attachment:", "err", err)
	}
	url := img.URL
	path, err := downloadImage(url)
	if err != nil {
		logger.Info("Something went wrong:", err)
	}
	grid := load(path)
	newPath := save("temp", img.Filename, filter(grid))
	sendImg, err := os.Open(newPath)
	if err != nil {
		logger.Info("Something went wrong,", err)
	}

	proxy.Respond(bacotell.Response{
		Content: url,
		Files: []*discordgo.File{
			{
				Name:   "img.jpg",
				Reader: sendImg,
			},
		},
	}, false, false, false)

	// deleteDir("temp")
	return nil
}

/* filepath noch korrigieren mit Input */
func load(filePath string) (grid [][]color.Color) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		logger.Info("Cannot read file:", "err", err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		logger.Info("Cannot decode file:", "err", err)
	}

	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		grid = append(grid, y)
	}
	return
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
		logger.Info("Cannot create file", "err", err)
	}
	defer file.Close()
	jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	return filePath
}

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

			if hue >= 45 && hue <= 75 {
				hue += 270
			} else if hue >= 75 && hue <= 135 {
				hue += 180
			} else if hue >= 225 && hue <= 255 {
				hue -= 120
			}

			newR, newG, newB := colorutil.HslToRgb(hue, sat, light)
			irImage[x][y] = color.NRGBA{newR, newG, newB, 255}
		}
	}
	return
}

func downloadImage(url string, directory string) (string, error) {
	err:= os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		logger.Info("Create directory failed", err)
	}
	
	fileName := filepath.Base(url)
	filePath := filepath.Join(directory, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		logger.Info("Cannot create file", "err", err)
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		logger.Info("Cannot get request", "err", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download failed with status code: %d", response.StatusCode)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		logger.Info("Something went wrong", "err", err)
	}

	return filePath, nil
}

func deleteDir(directory string) error {
	err := os.RemoveAll(directory)
	if err != nil {
		logger.Info("Cannot find the directory", err)
	}
	return nil
}
