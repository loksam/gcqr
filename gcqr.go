package main

import (
	"image/png"
	"os"
	"fmt"
	"image"
	"flag"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/fatih/color"
)

func main() {

	flag.Parse()
    var url = flag.Arg(0)

	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 50, 50)

	file, _ := os.Create("qrcode.png")
	defer file.Close()
	png.Encode(file, qrCode)

	fName := "qrcode.png"
    f, err := os.Open(fName)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
	defer f.Close()

    img, _, err := image.Decode(f)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
	}

	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if float64((r+g+b))/3 > 0.5 {
				color.New(color.BgWhite).Print("  ")
			} else {
				color.New(color.BgBlack).Print("  ")
			}
		}
		fmt.Printf("\n")
	}

}