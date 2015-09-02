package main

import (
	"errors"
	"fmt"
	"image/png"
	"os"

	"github.com/tortis/nn"
	"github.com/tortis/nn/activation"
)

const IMGDIM = 8

// Row major ordering
type letter [IMGDIM * IMGDIM]float64

func (this *letter) tofs() []float64 {
	return this[:IMGDIM*IMGDIM]
}

func main() {
	n := nn.NewANN(64, 4, activation.Logistic)
	n.AddHidden(30, true, activation.Logistic)
	n.Wire()

	// Load the training letters
	a, _ := loadLetter("letters/af.png")
	ra, _ := loadLetter("letters/a.png")
	b, _ := loadLetter("letters/b.png")
	c, _ := loadLetter("letters/c.png")
	d, _ := loadLetter("letters/d.png")
	fmt.Printf("Letter A: %v\n", a)
	fmt.Printf("Letter B: %v\n", b)
	fmt.Printf("Letter C: %v\n", c)
	fmt.Printf("Letter D: %v\n", d)

	// Load drop 3 samples
	ad3, _ := loadLetter("letters/ad3.png")
	bd3, _ := loadLetter("letters/bd3.png")
	cd3, _ := loadLetter("letters/cd3.png")
	dd3, _ := loadLetter("letters/dd3.png")
	fmt.Printf("Letter A drop 3 px: %v\n", ad3)
	fmt.Printf("Letter B drop 3 px: %v\n", bd3)
	fmt.Printf("Letter C drop 3 px: %v\n", cd3)
	fmt.Printf("Letter D drop 3 px: %v\n", dd3)

	// Load translate 1 samples
	at1, _ := loadLetter("letters/at1.png")
	fmt.Printf("Letter A translate 1px: %v\n", at1)

	// Create the training data and train
	exin := [][]float64{a.tofs(), b.tofs(), c.tofs(), d.tofs()}
	exout := [][]float64{{1.0, 0.0, 0.0, 0.0}, {0.0, 1.0, 0.0, 0.0}, {0.0, 0.0, 1.0, 0.0}, {0.0, 0.0, 0.0, 1.0}}
	lerr := n.Backprop(exin, exout, 100000, 0.025, 1, 0.5)
	fmt.Printf("The network error is: %f\n", lerr)

	// Validation
	ap := n.Predict(ra.tofs())
	bp := n.Predict(b.tofs())
	cp := n.Predict(c.tofs())
	dp := n.Predict(d.tofs())
	fmt.Printf("Prediction for pattern a: %v\n", ap)
	fmt.Printf("Prediction for pattern b: %v\n", bp)
	fmt.Printf("Prediction for pattern c: %v\n", cp)
	fmt.Printf("Prediction for pattern d: %v\n", dp)

	ad3p := n.Predict(ad3.tofs())
	bd3p := n.Predict(bd3.tofs())
	cd3p := n.Predict(cd3.tofs())
	dd3p := n.Predict(dd3.tofs())
	fmt.Printf("Prediction for pattern ad3: %v\n", ad3p)
	fmt.Printf("Prediction for pattern bd3: %v\n", bd3p)
	fmt.Printf("Prediction for pattern cd3: %v\n", cd3p)
	fmt.Printf("Prediction for pattern dd3: %v\n", dd3p)

	at1p := n.Predict(at1.tofs())
	fmt.Printf("Prediction for pattern at1: %v\n", at1p)
}

func loadLetter(fn string) (*letter, error) {
	// Open the file
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Attempt to decode as PNG
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	// Ensure the size is correct
	if img.Bounds().Max.X != IMGDIM || img.Bounds().Max.Y != IMGDIM {
		return nil, errors.New(fmt.Sprintf("The input image must be %dx%d pixels.", IMGDIM, IMGDIM))
	}

	// Build letter
	var l letter
	for y := 0; y < IMGDIM; y++ {
		for x := 0; x < IMGDIM; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			avg := float64(r+g+b) / 3.0
			scaled := avg / float64(0xFFFF)
			l[y*IMGDIM+x] = scaled
		}
	}
	return &l, nil
}
