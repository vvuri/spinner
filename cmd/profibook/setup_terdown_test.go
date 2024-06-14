package profibook

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"
)

func load(filePath string) (grid [][]color.Color) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		grid = append(grid, y)
	}
	return
}

// save the image to file
func save(filePath string, grid [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()
	png.Encode(file, img.SubImage(img.Rect))
}

func flip(grid [][]color.Color) {
	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col)/2; y++ {
			z := len(col) - y - 1
			col[y], col[z] = col[z], col[y]
		}
	}
}

// ---------------------------------------------

func setup(filename string) (teardown func(tempfile string), grid [][]color.Color) {
	grid = load(filename)
	teardown = func(tempfile string) {
		os.Remove(tempfile)
	}
	return
}

func TestFlip(t *testing.T) {
	teardown, grid := setup("monalisa.png")
	defer teardown("flipped.png")

	flip(grid)
	save("flipped.png", grid)
	g := load("flipped.png")
	if len(g) != 321 || len(g[0]) != 480 {
		t.Error("Grid is wrong size", "width:", len(g), "length:", len(g[0]))
	}
}
