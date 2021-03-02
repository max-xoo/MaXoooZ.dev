package utils

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"github.com/tfriedel6/canvas"
	Backend "github.com/tfriedel6/canvas/backend/softwarebackend"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
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
		NewError(500, "Internal Server Error")
		return 0, 0
	}
	image, _, err := image.DecodeConfig(file)

	if err != nil {
		NewError(500, "Internal Server Error")
		return 0, 0
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

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		magenta := color.FgMagenta.Render
		fmt.Printf(time.Now().Format("01/02/2006 15:04:05") + " %s " + r.RequestURI + " %s" + r.RemoteAddr + "%s\n", magenta(":"), magenta("["), magenta("]"))

		next.ServeHTTP(w, r)
	})
}
