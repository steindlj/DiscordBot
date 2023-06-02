package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"path/filepath"

	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/PerformLine/go-stockutil/colorutil"
	"github.com/bwmarrin/discordgo"
)

type LomoPurpleCommand struct{}

var _ common.Command = LomoPurpleCommand{}

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
func (LomoPurpleCommand) Execute(proxy common.ExecuteProxy) error {
	proxy.Defer(true)

	img, err := proxy.AttachmentOption("attachment")
	if err != nil {
		return fmt.Errorf("failed to retrieve attachment: %w", err)
	}

	url := img.URL
	tempDir := "temp"
	path, err := downloadImage(url, tempDir)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	grid,_ := load(path)
	newPath := save(tempDir, img.Filename, filter(grid))

	sendImg, err := os.Open(newPath)
	if err != nil {
		return fmt.Errorf("failed to open new image: %w", err)
	}
	defer sendImg.Close()

	proxy.Followup(common.Response{
		Content: url,
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

func load(filePath string) ([][]color.Color, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w",err)
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
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	fileName := filepath.Base(url)
	filePath := filepath.Join(directory, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to perform HTTP GET request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download failed with status code: %d", response.StatusCode)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
<<<<<<< HEAD
		return "", fmt.Errorf("failed to write file: ", err)
=======
		return "", fmt.Errorf("failed to write file: %w",err)
>>>>>>> 08e6e19 (add to error)
	}

	return filePath, nil
}

func deleteDir(directory string) error {
	err := os.RemoveAll(directory)
	if err != nil {
<<<<<<< HEAD
		return fmt.Errorf("Failed to delete directory: ", err)
=======
		return fmt.Errorf("failed to delete directory: %w",err)
>>>>>>> 08e6e19 (add to error)
	}
	return nil
}
