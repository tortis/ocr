package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
	"github.com/tortis/ocr/boxer"
)

const (
	INPUT_FILE     = "input.png"
	BIN_BLOCK_SIZE = 32
	OUT_DIR        = "letters"
)

func main() {
	// Load an image.
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Binarize the image.
	g := boxer.Binarize(img, BIN_BLOCK_SIZE)

	// Box the image.
	blobs := boxer.Blobify(g)

	// Write each box to a file
	for i, b := range blobs {
		if b.Bridges != nil {
			// Split the blob and write multiple files
			log.Println("Skipping wide blob.")
		} else {
			// Write one file for the whole blob
			outf, err := os.Create(OUT_DIR + fmt.Sprintf("/blob-%d.png", i))
			if err != nil {
				log.Println("Failed to open output letter file.")
			} else {
				defer outf.Close()
				m := resize.Resize(20, 20, b.Img, resize.NearestNeighbor)
				err = png.Encode(outf, m)
				if err != nil {
					log.Println("Failed to write image:", outf.Name(), ":", err)
				}
			}
		}
	}
}
