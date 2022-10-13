package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"log"
	"os"
)

type Output interface {
	Draw() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type DrawOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (d *DrawOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

func (d *DrawOutput) SetWriter(w io.Writer) {
	d.Writer = w
}

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

type TextSquare2 struct {
	DrawOutput
}

func (t *TextSquare2) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}

type ImageSquare2 struct {
	DrawOutput
}

func (i *ImageSquare2) Draw() error {
	width := 800
	height := 600

	bgColor := image.Uniform{color.RGBA{0, 0, 0, 1}}
	origin := image.Point{0, 0}
	quality := &jpeg.Options{Quality: 75}

	bgRectangle := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	draw.Draw(bgRectangle, bgRectangle.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)

	draw.Draw(bgRectangle, squareImg.Bounds(), &squareColor, origin, draw.Src)

	if i.Writer == nil {
		return fmt.Errorf("No writer stored on ImageSquare")
	}
	if err := jpeg.Encode(i.Writer, bgRectangle, quality); err != nil {
		return fmt.Errorf("Error writing image to disk")
	}

	if i.LogWriter != nil {
		i.LogWriter.Write([]byte("Image written in provided writer\n"))
	}

	return nil
}

func Factory(s string) (Output, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare2{
			DrawOutput: DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare2{
			DrawOutput: DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}

var output1 = flag.String("output", "text", "The output to use between "+
	"'console' and 'image' file")

func main() {
	flag.Parse()

	activeStrategy, err := Factory(*output1)
	if err != nil {
		log.Fatal(err)
	}

	switch *output1 {
	case TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case IMAGE_STRATEGY:
		w, err := os.Create("./image.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()

		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
