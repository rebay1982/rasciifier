package main

import (
	"fmt"
	"image/png"
	"flag"
	"os"

	"github.com/rebay1982/redscii"
)

const (
	appName = "rASCIIfier"
)

var (
	filename string
	widthMax int
)

func parseCmdlineInput() {
	flag.StringVar(&filename, "f", "", "Filename of PNG image to rASCIIfy.")
	flag.IntVar(&widthMax, "w", 300, "Maximum character width of output ASCII art.")
	flag.Parse()
}

func printUsage() {
	fmt.Printf("%s\n\nUsage\n\n", appName)
	flag.PrintDefaults()

}

func main() {
	parseCmdlineInput()
	if filename == "" {
		printUsage()
		os.Exit(1)
	}

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Printf("Cannot open %s file. File doesn't exist or read permission denied.", filename)
		os.Exit(1)
	}

	pngImage, err := png.Decode(file)
	if err != nil {
		fmt.Println("Cannot decode PNG file -- file isn't a PNG file perhaps?")
		os.Exit(1)
	}

	bounds := pngImage.Bounds()

	smallImage := redscii.DownscaleImage(pngImage, float64(widthMax >> 1) / float64(bounds.Dx()))
	greyImage := redscii.GreyScaleImage(smallImage)

	fmt.Println(appName)
	fmt.Printf("Generating from image %s (size %v), with maximum width of %d characters.\n", filename, bounds, widthMax)
	fmt.Println()
	fmt.Println()

	redscii.ASCIIfy(greyImage)
}
