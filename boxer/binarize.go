package boxer

import (
	"image"
	"log"
)
import "image/color"

var (
	black = color.Gray{0x00}
	white = color.Gray{0xFF}
)

const (
	MIN_DR        = 24
	LOW_DR_THRESH = 60
	AVG_OFFSET    = -40
)

func Binarize(img image.Image, blockSize int) *image.Gray {
	// Create a new Gray that is the same size as the given image.
	g := image.NewGray(img.Bounds())

	// Loop over each block
	for y := 0; y < img.Bounds().Max.Y; y += blockSize {
		log.Printf("On row %d\n", y)
		for x := 0; x < img.Bounds().Max.X; x += blockSize {
			t := computeBlockThreshold(x, y, blockSize, img)
			setBlock(x, y, blockSize, t, img, g)
		}
	}

	return g
}

func computeBlockThreshold(x, y, blockSize int, img image.Image) int {
	max := 0
	min := 255
	avg := 0
	pixels := 0

	// Loop through all the pixels in the block
	for cy := y; cy < y+blockSize; cy++ {
		if cy >= img.Bounds().Max.Y {
			break
		}
		for cx := x; cx < x+blockSize; cx++ {
			if cx >= img.Bounds().Max.X {
				break
			}
			l := int(lumAt(cx, cy, img))
			if l > max {
				max = l
			}
			if l < min {
				min = l
			}
			avg += l
			pixels += 1
		}
	}

	avg /= pixels
	log.Printf("Avg: %d, Pixels: %d, max: %d, min: %d\n", avg, pixels, max, min)

	// Compute the threshold
	if (max - min) > MIN_DR {
		log.Printf("Average for HDR block is: %d\n", avg)
		return avg + AVG_OFFSET
	} else {
		if avg > LOW_DR_THRESH {
			return 0
		} else {
			return 255
		}
	}
}

func setBlock(x, y, blockSize, t int, img image.Image, g *image.Gray) {
	// Loop through all the pixels in the block
	for cy := y; cy < y+blockSize; cy++ {
		if cy >= img.Bounds().Max.Y {
			break
		}
		for cx := x; cx < x+blockSize; cx++ {
			if cx >= img.Bounds().Max.X {
				continue
			}
			l := int(lumAt(cx, cy, img))
			if l > t {
				g.SetGray(cx, cy, white)
			} else {
				g.SetGray(cx, cy, black)
			}
		}
	}
}

func lumAt(x, y int, img image.Image) uint8 {
	R, G, B, _ := img.At(x, y).RGBA()
	Rf := float64(R & 0xFF)
	Gf := float64(G & 0xFF)
	Bf := float64(B & 0xFF)
	//return uint8((Rf + Rf + Bf + Gf + Gf + Gf) / 6.0)
	return uint8(0.2126*Rf + 0.7152*Gf + 0.0722*Bf)
}

func toGray(img image.Image) *image.Gray {
	g := image.NewGray(img.Bounds())
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			g.SetGray(x, y, color.Gray{lumAt(x, y, img)})
		}
	}
	return g
}
