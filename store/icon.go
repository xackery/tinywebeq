package store

import (
	"fmt"
	"image"
	"os"

	"github.com/xypwn/filediver/dds"
)

func fetchDDS(path string) (image.Image, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer r.Close()

	img, err := dds.Decode(r, false)
	if err != nil {
		return nil, fmt.Errorf("dds.Decode: %w", err)
	}

	return img, nil
}
