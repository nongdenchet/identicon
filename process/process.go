package process

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"github.com/nongdenchet/identicon/helpers"
)

func GenerateImage(hash []byte, size int) (string, error) {
	// colors
	r, g, b := pickColor(hash)
	drawColor := color.RGBA{r, g, b, 255}
	bgColor := color.RGBA{255, 255, 255, 255}

	// init image
	imageRect := image.Rect(0, 0, size, size)
	result := image.NewRGBA(imageRect)
	draw.Draw(result, imageRect, &image.Uniform{bgColor}, image.ZP, draw.Src)

	// generate
	grid := buildGrid(hash)
	pixelMap := buildPixelMap(grid, size/5)
	for _, rect := range pixelMap {
		draw.Draw(result, rect, &image.Uniform{drawColor}, image.ZP, draw.Src)
	}

	// prepare file
	filePath := helpers.GetIdenticonFilePath(hash, size)
	file, err := helpers.PrepareFile(filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	// store image
	err = png.Encode(file, result)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func pickColor(hash []byte) (r, g, b byte) {
	if len(hash) >= 3 {
		return hash[0], hash[1], hash[2]
	}

	return 0, 0, 0
}

func buildPixelMap(grid []byte, size int) []image.Rectangle {
	var result []image.Rectangle

	for index, item := range grid {
		if item%2 == 1 {
			x := (index % 5) * size
			y := (index / 5) * size
			result = append(result, image.Rect(x, y, x+size, y+size))
		}
	}

	return result
}

func buildGrid(data []byte) []byte {
	var result []byte

	for i := 0; i < len(data); i += 3 {
		end := i + 3
		if end > len(data) {
			break
		}

		row := append([]byte(nil), data[i:end]...)
		for j := end - 2; j >= i; j-- {
			row = append(row, data[j])
		}

		result = append(result, row...)
	}

	return result
}
