package service

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/nongdenchet/identicon/encode"
	"github.com/nongdenchet/identicon/helpers"
	"github.com/nongdenchet/identicon/process"
)

type IdenticonServiceImpl struct{}

func (IdenticonServiceImpl) Generate(_ context.Context, text string, size int) (string, error) {
	if len(text) == 0 {
		return "", errors.New("text should not empty")
	}

	if size <= 0 {
		return "", errors.New("size should not empty and must be larger than 0")
	}

	hash, err := encode.GetMD5Hash(text)
	if err != nil {
		return "", err
	}

	filePath := helpers.GetIdenticonFilePath(hash, size)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		filePath, err = process.GenerateImage(hash, size)
		if err != nil {
			return "", err
		}
	}

	url, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	return url, nil
}
