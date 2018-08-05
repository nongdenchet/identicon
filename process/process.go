package process

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
)

const dirPath = "output"

func GenerateImage(hash []byte, size int) (string, error) {
	r, g, b := pickColor(hash)
	drawColor := color.RGBA{r, g, b, 255}
	bgColor := color.RGBA{255, 255, 255, 255}

	imageRect := image.Rect(0, 0, size, size)
	result := image.NewRGBA(imageRect)
	draw.Draw(result, imageRect, &image.Uniform{bgColor}, image.ZP, draw.Src)

	grid := buildGrid(hash)
	pixelMap := buildPixelMap(grid, size/5)
	for _, rect := range pixelMap {
		draw.Draw(result, rect, &image.Uniform{drawColor}, image.ZP, draw.Src)
	}

	filePath := dirPath + "/" + string(hash) + ".png"
	file, err := prepareFile(filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	err = png.Encode(file, result)
	if err != nil {
		return "", err
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

func prepareFile(filePath string) (*os.File, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0600)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
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
