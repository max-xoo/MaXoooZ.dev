package utils

import (
	"bytes"
	"fmt"
	"github.com/tfriedel6/canvas"
	Backend "github.com/tfriedel6/canvas/backend/softwarebackend"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

func Mt_rand(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func GetImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	image, _, err := image.DecodeConfig(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}

func GetIMG(category string, specific string) ([]byte, *Error) {
	arrFiles := []string{}

	files, err := ioutil.ReadDir(path.Join("assets", "imgs", category, specific))

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		arrFiles = append(arrFiles, f.Name())
	}
	fileName := arrFiles[Mt_rand(0, len(arrFiles))]
	i := strings.Index(fileName, ".png")

	if i > -1 {
		width, height := GetImageDimension(path.Join("assets", "imgs", category, specific, fileName))

		backend := Backend.New(width, height)
		backend.MSAA = 1

		ctx := canvas.New(backend)
		img, err := ctx.LoadImage(path.Join("assets", "imgs", category, specific, fileName))

		if err != nil {
			return nil, NewError(503, "Service Unavailable")
		}
		w, h := float64(ctx.Width()), float64(ctx.Height())

		ctx.DrawImage(img, 0, 0, w, h)
		ctx.Fill()

		ctx.Stroke()
		defer img.Delete()

		f := bytes.NewBuffer([]byte{})
		err = png.Encode(f, backend.Image)

		if err != nil {
			return nil, NewError(503, "Service Unavailable")
		}
		b := f.Bytes()
		return b, nil
	} else {
		width, height := GetImageDimension(path.Join("assets", "imgs", category, specific, fileName))

		backend := Backend.New(width, height)
		backend.MSAA = 1

		ctx := canvas.New(backend)
		img, err := ctx.LoadImage(path.Join("assets", "imgs", category, specific, fileName))

		if err != nil {
			return nil, NewError(503, "Service Unavailable")
		}
		w, h := float64(ctx.Width()), float64(ctx.Height())

		ctx.DrawImage(img, 0, 0, w, h)
		ctx.Fill()

		ctx.Stroke()
		defer img.Delete()

		f := bytes.NewBuffer([]byte{})
		err = jpeg.Encode(f, backend.Image, &jpeg.Options{jpeg.DefaultQuality})

		if err != nil {
			return nil, NewError(503, "Service Unavailable")
		}
		b := f.Bytes()
		return b, nil
	}
}
