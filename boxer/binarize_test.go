package boxer

import (
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestBin(t *testing.T) {
	img := loadPNG("test4.png")
	g := Binarize(img, 32)
	savePNG("final.png", g)
	//gg := toGray(img)
	//savePNG("gg.png", gg)
}

func loadPNG(fname string) image.Image {
	log.Printf("Loading test image %s\n", fname)
	// Open the file
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Attempt to decode as PNG
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func savePNG(fname string, img image.Image) {
	log.Printf("Saving test image %s\n", fname)
	file, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}
