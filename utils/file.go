package utils

import (
	"image"
	_ "image/png"
	"io"
	"net/http"
	"os"
)

func DownloadRemoteImage(url string) (*image.Image, error) {
	var result, donwloadErr = http.Get(url)
	if donwloadErr != nil {
		return nil, donwloadErr
	}
	defer result.Body.Close()

	var file, fileErr = os.Create("tech.png")
	if fileErr != nil {
		return nil, fileErr
	}
	defer file.Close()

	var _, copyErr = io.Copy(file, result.Body)
	if copyErr != nil {
		return nil, copyErr
	}

	file.Seek(0, 0)

	var decodedImage, _, decodeErr = image.Decode(file)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return &decodedImage, nil
}
