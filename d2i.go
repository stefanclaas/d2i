package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

const blockWidth = 8
const maxImageWidth = 320

func main() {
	// Define the colors for the numbers 0-9.
	colors := []color.RGBA{
		{230, 25, 75, 255},    // 0: Red
		{245, 130, 49, 255},   // 1: Orange
		{255, 225, 25, 255},   // 2: Yellow
		{191, 239, 69, 255},   // 3: Lime
		{60, 180, 75, 255},    // 4: Green
		{66, 212, 244, 255},   // 5: Cyan
		{67, 99, 216, 255},    // 6: Blue
		{145, 30, 180, 255},   // 7: Purple
		{240, 50, 230, 255},   // 8: Magenta
		{169, 160, 169, 255},  // 9: Grey
	}

	// Define the flags for encoding and decoding.
	encode := flag.Bool("e", false, "To encode")
	decode := flag.Bool("d", false, "To decode")
	flag.Parse()

	// The file name is the first unchecked argument.
	filename := flag.Arg(0)

	if *encode {
		// Read the number string from the standard input.
		reader := bufio.NewReader(os.Stdin)
		numbers, _ := reader.ReadString('\n')
		numbers = strings.TrimSpace(numbers) // Remove the newline character.

		// Calculate the number of rows needed.
		numRows := (len(numbers)*blockWidth + maxImageWidth - 1) / maxImageWidth

		// Create a new image.
		img := image.NewRGBA(image.Rect(0, 0, maxImageWidth, numRows*blockWidth))

		// Fill the image with the corresponding colors.
		for i := 0; i < len(numbers); i++ {
			// Convert the number to a color.
			n, _ := strconv.Atoi(string(numbers[i]))
			for x := (i % (maxImageWidth / blockWidth)) * blockWidth; x < (i%(maxImageWidth/blockWidth)+1)*blockWidth; x++ {
				for y := (i / (maxImageWidth / blockWidth)) * blockWidth; y < ((i/(maxImageWidth/blockWidth))+1)*blockWidth; y++ {
					img.Set(x, y, colors[n])
				}
			}
		}

		// Save the image as a .png file.
		f, _ := os.Create(filename)
		defer f.Close()
		png.Encode(f, img)
	} else if *decode {
		// Open the .png file.
		in, _ := os.Open(filename)
		defer in.Close()
		imgIn, _, _ := image.Decode(in)

		// Decode the colors back into numbers and print in a single line.
		for y := 0; y < imgIn.Bounds().Dy(); y += blockWidth {
			for x := 0; x < imgIn.Bounds().Dx(); x += blockWidth {
				c := color.RGBAModel.Convert(imgIn.At(x, y)).(color.RGBA)
				for j, col := range colors {
					if c.R == col.R && c.G == col.G && c.B == col.B {
						fmt.Print(j)
						break
					}
				}
			}
		}

		// Add a newline at the end to separate the prompt.
		fmt.Println()
	} else {
		fmt.Println("Please enter either -e for encoding or -d for decoding.")
	}
}

