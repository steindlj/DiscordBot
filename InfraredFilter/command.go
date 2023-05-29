package main

import (
	"net/http"

	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/PerformLine/go-stockutil/colorutil"
	"github.com/bwmarrin/discordgo"

	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

type Infrared struct{}

var (
	permission int64 = discordgo.PermissionSendMessages
	_ provider.Command = Infrared{}
)

// CommandData implements provider.Command
func (Infrared) CommandData() (discordgo.ApplicationCommand, error){
	return discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "lomo_filter",
		Description: "Lomo Purple Filter",
		DefaultMemberPermissions: &permission,
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
func (Infrared) Execute(proxy provider.ExecuteProxy) error {
	img, err := proxy.AttachmentOption("attachment")
	if err != nil {
		log.Println("Cannot find attachment:", err)
	}
	url := img.URL

	proxy.Respond(provider.Response{
		Content: url,
	}, false, false, false)
	return nil
}

/* filepath noch korrigieren mit Input */
func load (filePath string) (grid [][]color.Color) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
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

func save (filePath string, grid [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0,0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x,y, grid[x][y])
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()
	jpeg.Encode(file, img, nil)
}

func filter (grid [][]color.Color) (irImage [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	irImage = make([][]color.Color, xlen)
	for i := 0; i < len(irImage); i++ {
		irImage[i] = make([]color.Color, ylen)
	}
	for x := 0; x < xlen; x++ {
		for y :=0; y < ylen; y++ {
			pix := grid[x][y].(color.YCbCr)
			R, G, B := color.YCbCrToRGB(pix.Y, pix.Cb, pix.Cr)

			hue, sat, light := colorutil.RgbToHsl(float64(R), float64(G), float64(B))

			if hue >= 45 && hue <= 75 {
				hue+= 270
			} else if hue >= 75 && hue <= 135 {
				hue+= 180
			} else if hue >= 225 && hue <= 255 {
				hue-= 120
			}

			newR, newG, newB := colorutil.HslToRgb(hue, sat, light)
			irImage[x][y] = color.NRGBA{newR, newG, newB, 255}
		}
	}
	return
}

func downloadImage (url string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Println("Cannot create file", err)
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Println("Cannot get request", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		
	}
}